package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/mallagent/conversation"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/errno"
	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type DeleteHistoryService struct {
	ctx context.Context
} // NewDeleteHistoryService new DeleteHistoryService
func NewDeleteHistoryService(ctx context.Context) *DeleteHistoryService {
	return &DeleteHistoryService{ctx: ctx}
}

// Run create note info
func (s *DeleteHistoryService) Run(req *llm.DeleteHistoryRequest) (resp *llm.DeleteHistoryResponse, err error) {
	userId := req.UserId
	convId := req.ConversationId
	if userId == "" || convId == "" {
		return nil, kerrors.NewBizStatusError(errno.ErrGRPCRequestParam, "invalid params")
	}
	err = conversation.GetDefaultBucket(s.ctx).DeleteConversation(convId, userId)
	if err != nil {
		klog.CtxErrorf(s.ctx, "DeleteConversation failed, err: %v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrDelConversation, "DeleteConversation failed")
	}
	return
}
