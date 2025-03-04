package service

import (
	"context"
	"fmt"

	checkout "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/checkout"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	gatewayutils "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/utils"
	rpcorder "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/order"
	rpcpayment "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/hertz/pkg/app"
)

type ChargeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewChargeService(Context context.Context, RequestContext *app.RequestContext) *ChargeService {
	return &ChargeService{RequestContext: RequestContext, Context: Context}
}

func (h *ChargeService) Run(req *checkout.ChargeReq) (resp *checkout.ChargeResp, err error) {
	// todo 分布式事务
	fmt.Printf("ChargeService Run req: %+v\n", req)
	transactionId, err := func() (string, error) {
		_, err = rpc.OrderClient.MarkOrderPaid(h.Context, &rpcorder.MarkOrderPaidReq{
			UserId:  gatewayutils.GetUserIdFromCtx(h.RequestContext),
			OrderId: req.OrderId,
		})
		if err != nil {
			return "", err
		}
		r, err := rpc.PaymentClient.Charge(h.Context, &rpcpayment.ChargeReq{
			OrderId:       req.OrderId,
			Amount:        req.Amount,
			UserId:        gatewayutils.GetUserIdFromCtx(h.RequestContext),
			PaymentMethod: req.PaymentMethod,
			CreditCard: &rpcpayment.CreditCardInfo{
				CreditCardNumber:          req.CreditCardNumber,
				CreditCardExpirationYear:  req.CreditCardExpirationYear,
				CreditCardExpirationMonth: req.CreditCardExpirationMonth,
				CreditCardCvv:             req.CreditCardCvv,
			},
		})
		return r.TransactionId, err
	}()
	if err != nil {
		return
	}
	resp = &checkout.ChargeResp{
		TransactionId: transactionId,
	}
	return
}
