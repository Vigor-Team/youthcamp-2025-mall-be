package service

import (
	"context"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/errno"
	rpcllm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/kerrors"

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
	convId := req.ConversationId
	if convId == "" {
		hlog.CtxErrorf(h.Context, "get history failed, err: conversation id is empty")
		return nil, kerrors.NewBizStatusError(errno.ErrHTTPRequestParam, "conversation id is empty")
	}
	history, err := rpc.LlmClient.GetHistory(h.Context, &rpcllm.GetHistoryRequest{
		ConversationId: convId,
		UserId:         h.RequestContext.Value("user_id").(string),
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
