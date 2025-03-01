package service

import (
	"context"
	"strconv"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/rpc"
	gatewayutils "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/utils"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/errno"
	rpcllm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/kerrors"

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
	convId := req.ConversationId
	if convId == "" {
		hlog.CtxErrorf(h.Context, "delete history failed, err: conversation id is empty")
		return nil, kerrors.NewBizStatusError(errno.ErrHTTPRequestParam, "conversation id is empty")
	}
	_, err = rpc.LlmClient.DeleteHistory(h.Context, &rpcllm.DeleteHistoryRequest{
		ConversationId: convId,
		UserId:         strconv.Itoa(int(gatewayutils.GetUserIdFromCtx(h.RequestContext))),
	})
	if err != nil {
		hlog.CtxErrorf(h.Context, "delete history failed, err: %v", err)
		return
	}
	return
}
