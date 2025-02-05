package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	rpcllm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/llm"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteMessageService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteMessageService(Context context.Context, RequestContext *app.RequestContext) *DeleteMessageService {
	return &DeleteMessageService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteMessageService) Run(req *llm.DeleteHistoryRequest) (resp *llm.DeleteHistoryResponse, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	convId := req.ConversationId
	if convId == "" {
		hlog.CtxErrorf(h.Context, "delete history failed, err: conversation id is empty")
		return
	}
	_, err = rpc.LlmClient.DeleteHistory(h.Context, &rpcllm.DeleteHistoryRequest{
		ConversationId: convId,
		UserId:         "12345", // todo test, should get from jwt
	})
	if err != nil {
		hlog.CtxErrorf(h.Context, "delete history failed, err: %v", err)
		return
	}
	return
}
