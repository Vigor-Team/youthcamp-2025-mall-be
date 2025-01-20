package consts

import (
	"github.com/cloudwego/kitex/pkg/kerrors"
)

var (
	ErrUserNotFound = kerrors.NewBizStatusError(10001, "user not found")
	ErrPassword     = kerrors.NewBizStatusError(10002, "password error")
	ErrorUserExist  = kerrors.NewBizStatusError(10003, "user exist")
)
