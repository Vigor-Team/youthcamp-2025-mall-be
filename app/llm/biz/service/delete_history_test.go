package service

import (
	"context"
	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"testing"
)

func TestDeleteHistory_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteHistoryService(ctx)
	// init req and assert value

	req := &llm.DeleteHistoryRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
