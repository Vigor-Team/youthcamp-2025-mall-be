package service

import (
	"context"
	"strconv"

	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/llm"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	gatewayutils "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/utils"
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
	llmResp, err := rpc.LlmClient.SendMessage(h.Context, &rpcllm.ChatRequest{
		Message:        req.GetMessage(),
		UserId:         strconv.Itoa(int(gatewayutils.GetUserIdFromCtx(h.RequestContext))),
		ConversationId: req.ConversationId,
	})
	if err != nil {
		return nil, err
	}
	resp = &llm.ChatResponse{
		Response: llmResp.Response,
	}
	return
}
