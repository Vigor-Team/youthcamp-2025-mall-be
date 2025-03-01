package service

import (
	"context"
	"strconv"

	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/llm"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	gatewayutils "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/utils"
	rpcllm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm/llmservice"
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
	client, err := rpc.LlmClient.StreamMessage(h.Context, &rpcllm.ChatRequest{
		Message:        req.GetMessage(),
		UserId:         strconv.Itoa(int(gatewayutils.GetUserIdFromCtx(h.RequestContext))),
		ConversationId: req.ConversationId,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}
