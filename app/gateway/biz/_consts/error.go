package _consts

import "github.com/cloudwego/kitex/pkg/kerrors"

var (
	ErrTokenInvalid  = kerrors.NewBizStatusError(10002, "token is invalid")
	ErrTokenNotFound = kerrors.NewBizStatusError(10001, "token not found")
	ErrTokenFormat   = kerrors.NewBizStatusError(1004, "invalid token format")
	ErrTokenExpired  = kerrors.NewBizStatusError(10003, "token is expired")
)
