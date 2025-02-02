package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	rpcllm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/llm"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetHistoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetHistoryService(Context context.Context, RequestContext *app.RequestContext) *GetHistoryService {
	return &GetHistoryService{RequestContext: RequestContext, Context: Context}
}

func (h *GetHistoryService) Run(req *llm.GetHistoryRequest) (resp *llm.GetHistoryResponse, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	convId := req.ConversationId
	if convId == "" {
		hlog.CtxErrorf(h.Context, "get history failed, err: conversation id is empty")
		return
	}
	history, err := rpc.LlmClient.GetHistory(h.Context, &rpcllm.GetHistoryRequest{
		ConversationId: convId,
		UserId:         "12345", // todo test, should get from jwt
	})
	if err != nil {
		hlog.CtxErrorf(h.Context, "get history failed, err: %v", err)
		return
	}
	msg := make([]*llm.Message, 0)
	for _, message := range history.History {
		msg = append(msg, &llm.Message{
			Role:    message.Role,
			Content: message.Content,
		})
	}
	resp = &llm.GetHistoryResponse{
		History: msg,
	}
	return
}
