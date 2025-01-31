package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	rpcllm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm/llmservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/llm"
	"github.com/cloudwego/hertz/pkg/app"
)

type StreamMessageService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewStreamMessageService(Context context.Context, RequestContext *app.RequestContext) *StreamMessageService {
	return &StreamMessageService{RequestContext: RequestContext, Context: Context}
}

func (h *StreamMessageService) Run(req *llm.ChatRequest) (resp llmservice.LlmService_StreamMessageClient, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code
	client, err := rpc.LlmClient.StreamMessage(h.Context, &rpcllm.ChatRequest{Message: req.GetMessage()})
	if err != nil {
		return nil, err
	}
	return client, nil
}
