package consts

import "github.com/cloudwego/kitex/pkg/kerrors"

var (
	ErrInputTextNotFound = kerrors.NewBizStatusError(50000, "input text not found")
	ErrJsonUnmarshal     = kerrors.NewBizStatusError(50001, "json unmarshal error")
	ErrJsonMarshal       = kerrors.NewBizStatusError(50002, "json marshal error")
	ErrSearchProducts    = kerrors.NewBizStatusError(50003, "search products error")
	ErrCreateOrder       = kerrors.NewBizStatusError(50004, "create order error")
)
