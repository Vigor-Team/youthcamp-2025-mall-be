package service

import (
	"context"
	dalmongo "github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/dal/mongo"
	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"testing"
)

func TestGetHistory_Run(t *testing.T) {
	dalmongo.Init()
	ctx := context.Background()
	s := NewGetHistoryService(ctx)
	// init req and assert value

	req := &llm.GetHistoryRequest{
		UserId:         "12345",
		ConversationId: "1f54733c-67e9-40d7-9d8f-5f6dcf5b62df",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
