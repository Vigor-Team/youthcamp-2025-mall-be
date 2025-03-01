package service

import (
	"context"
	"strconv"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	gatewayutils "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/utils"
	rpcllm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/llm"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetConversationIdsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetConversationIdsService(Context context.Context, RequestContext *app.RequestContext) *GetConversationIdsService {
	return &GetConversationIdsService{RequestContext: RequestContext, Context: Context}
}

func (h *GetConversationIdsService) Run(req *llm.GetConversationIdsRequest) (resp *llm.GetConversationIdsResponse, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	history, err := rpc.LlmClient.GetHistory(h.Context, &rpcllm.GetHistoryRequest{
		UserId: strconv.Itoa(int(gatewayutils.GetUserIdFromCtx(h.RequestContext))),
	})
	if err != nil {
		hlog.CtxErrorf(h.Context, "get history failed, err: %v", err)
		return
	}
	resp = &llm.GetConversationIdsResponse{
		ConversationIds: history.ConversationIds,
	}
	return
}
