package consts

import "github.com/cloudwego/kitex/pkg/kerrors"

var (
	ErrReqParamNotFound   = kerrors.NewBizStatusError(50000, "request param not found")
	ErrJsonUnmarshal      = kerrors.NewBizStatusError(50001, "json unmarshal error")
	ErrJsonMarshal        = kerrors.NewBizStatusError(50002, "json marshal error")
	ErrSearchProducts     = kerrors.NewBizStatusError(50003, "search products error")
	ErrCreateOrder        = kerrors.NewBizStatusError(50004, "create order error")
	ErrGetConversation    = kerrors.NewBizStatusError(50005, "get conversation error")
	ErrDeleteConversation = kerrors.NewBizStatusError(50006, "delete conversation error")
	ErrAppendMessage      = kerrors.NewBizStatusError(50007, "append message error")
)
