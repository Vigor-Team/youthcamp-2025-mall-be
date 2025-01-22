package service

import (
	"context"
	user "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/user"
	"testing"
)

func TestUserInfo_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUserInfoService(ctx)
	// init req and assert value

	req := &user.UserInfoReq{
		UserId: 4,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
