package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/llm"
	rpcllm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"github.com/cloudwego/hertz/pkg/app"
)

type SendMessageService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSendMessageService(Context context.Context, RequestContext *app.RequestContext) *SendMessageService {
	return &SendMessageService{RequestContext: RequestContext, Context: Context}
}

func (h *SendMessageService) Run(req *llm.ChatRequest) (resp *llm.ChatResponse, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code
	llmResp, err := rpc.LlmClient.SendMessage(h.Context, &rpcllm.ChatRequest{
		Message: req.GetMessage(),
		UserId:  "12345", // todo test, should get from jwt
	})
	if err != nil {
		return nil, err
	}
	resp = &llm.ChatResponse{
		Response: llmResp.Response,
	}
	return
}
