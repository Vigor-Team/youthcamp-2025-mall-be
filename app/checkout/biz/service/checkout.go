package service

import (
	"context"
	"fmt"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/checkout/biz/consts"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"strconv"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/checkout/infra/mq"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/checkout/infra/rpc"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/cart"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/checkout"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/email"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/order"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/payment"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

/*
	Run

1. get cart
2. calculate cart
3. create order
4. empty cart
5. pay
6. change order result
7. finish
*/
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Idempotent
	// get cart
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		klog.CtxErrorf(s.ctx, "GetCart.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrRPCGetCart, "get cart error")
	}
	if cartResult == nil || cartResult.Cart == nil || len(cartResult.Cart.Items) == 0 {
		return nil, kerrors.NewBizStatusError(consts.ErrRPCGetCart, "empty cart")
	}
	var (
		oi    []*order.OrderItem
		total float32
	)
	for _, cartItem := range cartResult.Cart.Items {
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: cartItem.ProductId})
		if resultErr != nil {
			klog.CtxErrorf(s.ctx, "GetProduct.err:%v", resultErr)
			return nil, kerrors.NewBizStatusError(consts.ErrRPCGetProduct, "get product error")
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		cost := p.Price * float32(cartItem.Quantity)
		total += cost
		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{ProductId: cartItem.ProductId, Quantity: cartItem.Quantity},
			Cost: cost,
		})
	}

	// create order
	orderReq := &order.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: "USD",
		OrderItems:   oi,
		Email:        req.Email,
	}
	if req.Address != nil {
		addr := req.Address
		zipCodeInt, _ := strconv.Atoi(addr.ZipCode)
		orderReq.Address = &order.Address{
			StreetAddress: addr.StreetAddress,
			City:          addr.City,
			Country:       addr.Country,
			State:         addr.State,
			ZipCode:       int32(zipCodeInt),
		}
	}
	orderResult, err := rpc.OrderClient.PlaceOrder(s.ctx, orderReq)
	if err != nil {
		klog.CtxErrorf(s.ctx, "PlaceOrder.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrRPCPlaceOrder, "place order error")
	}
	klog.Info("orderResult", orderResult)

	// empty cart
	emptyResult, err := rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		err = fmt.Errorf("EmptyCart.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrRPCEmptyCart, "empty cart error")
	}
	klog.Info(emptyResult)

	// charge
	var orderId string
	if orderResult != nil || orderResult.Order != nil {
		orderId = orderResult.Order.OrderId
	}
	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
		},
		PaymentMethod: req.PaymentMethod,
	}
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		klog.CtxErrorf(s.ctx, "Charge.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrRPCCharge, "charge error")
	}
	data, _ := proto.Marshal(&email.EmailReq{
		From:        "mall@example.com",
		To:          req.Email,
		ContentType: "text/plain",
		Subject:     "You just created an order in Mall",
		Content:     "You just created an order in Mall",
	})
	msg := &nats.Msg{Subject: "email", Data: data, Header: make(nats.Header)}

	// otel inject
	otel.GetTextMapPropagator().Inject(s.ctx, propagation.HeaderCarrier(msg.Header))

	_ = mq.Nc.PublishMsg(msg)

	klog.Info(paymentResult)
	// change order state
	klog.Info(orderResult)
	_, err = rpc.OrderClient.MarkOrderPaid(s.ctx, &order.MarkOrderPaidReq{UserId: req.UserId, OrderId: orderId})
	if err != nil {
		klog.CtxErrorf(s.ctx, "MarkOrderPaid.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrRPCMarkOrderPaid, "mark order paid error")
	}

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}
	return
}
