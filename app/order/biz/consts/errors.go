package consts

import "github.com/cloudwego/kitex/pkg/kerrors"

var (
	ErrInvalidParams = kerrors.NewBizStatusError(60001, "参数校验失败")
	ErrJsonParse     = kerrors.NewBizStatusError(60002, "json解析失败")
	ErrRedis         = kerrors.NewBizStatusError(60003, "redis操作失败")
	ErrMysql         = kerrors.NewBizStatusError(60004, "mysql操作失败")
)

var (
	ErrInventoryShortage  = kerrors.NewBizStatusError(60011, "库存不足")
	ErrDuplicateUser      = kerrors.NewBizStatusError(60012, "请勿重复操作")
	ErrOrderExists        = kerrors.NewBizStatusError(60013, "订单已存在")
	ErrDecrementInventory = kerrors.NewBizStatusError(60014, "减库存失败")
	ErrIncrementInventory = kerrors.NewBizStatusError(60015, "加库存失败")
)

var (
	ErrPublishMessage = kerrors.NewBizStatusError(60021, "发布消息失败")
	ErrConsumeMessage = kerrors.NewBizStatusError(60022, "消费消息失败")
)
