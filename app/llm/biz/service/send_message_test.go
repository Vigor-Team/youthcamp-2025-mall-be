package service

import (
	"context"
	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"testing"
)

func TestSendMessage_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSendMessageService(ctx)
	// init req and assert value

	req := &llm.ChatRequest{
		Message: "我觉得自己写的代码太烂了",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
