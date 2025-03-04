package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/dal/mysql"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/biz/model"
	rpccart "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/cart"
	order "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type QueryOrderService struct {
	ctx context.Context
} // NewQueryOrderService new QueryOrderService
func NewQueryOrderService(ctx context.Context) *QueryOrderService {
	return &QueryOrderService{ctx: ctx}
}

// Run create note info
func (s *QueryOrderService) Run(req *order.QueryOrderReq) (resp *order.QueryOrderResp, err error) {
	get, err := model.GetOrder(mysql.DB, s.ctx, req.UserId, req.OrderId)
	if err != nil {
		klog.CtxErrorf(s.ctx, "model.GetOrder.err:%v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrGetOrder, "model.GetOrder error")
	}

	var ois []*order.OrderItem
	for _, oi := range get.OrderItems {
		ois = append(ois, &order.OrderItem{
			Item: &rpccart.CartItem{
				ProductId: oi.ProductId,
				Quantity:  oi.Quantity,
			},
			Cost: oi.Cost,
		})
	}
	resp = &order.QueryOrderResp{
		Order: &order.Order{
			OrderId:      get.OrderId,
			UserId:       get.UserId,
			UserCurrency: get.UserCurrency,
			Address: &order.Address{
				StreetAddress: get.Consignee.StreetAddress,
				City:          get.Consignee.City,
				State:         get.Consignee.State,
				Country:       get.Consignee.Country,
				ZipCode:       get.Consignee.ZipCode,
			},
			OrderItems: ois,
			Email:      get.Consignee.Email,
			Status:     string(get.OrderState),
			CreatedAt:  int32(get.CreatedAt.Unix()),
		},
	}
	return
}
