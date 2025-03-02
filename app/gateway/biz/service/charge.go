package service

import (
	"context"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/checkout/infra/rpc"
	checkout "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/checkout"
	gatewayutils "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/utils"
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
	if err != nil {
		return
	}
	resp = &checkout.ChargeResp{
		TransactionId: r.TransactionId,
	}
	return
}
