package service

import (
	"context"
	user "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/user"
	"testing"
)

func TestDeleteUser_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteUserService(ctx)
	// init req and assert value

	req := &user.UserDeleteReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
