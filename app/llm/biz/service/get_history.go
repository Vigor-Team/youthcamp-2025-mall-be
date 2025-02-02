package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/mallagent/conversation"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GetHistoryService struct {
	ctx context.Context
} // NewGetHistoryService new GetHistoryService
func NewGetHistoryService(ctx context.Context) *GetHistoryService {
	return &GetHistoryService{ctx: ctx}
}

// Run create note info
func (s *GetHistoryService) Run(req *llm.GetHistoryRequest) (resp *llm.GetHistoryResponse, err error) {
	userId := req.UserId
	if userId == "" {
		err = consts.ErrReqParamNotFound
		return
	}
	if req.ConversationId == "" {
		ids, err := conversation.GetDefaultBucket(s.ctx).ListConversations(userId)
		if err != nil {
			klog.CtxErrorf(s.ctx, "get conversation list failed, err: %v", err)
			return nil, consts.ErrGetConversation
		}

		resp = &llm.GetHistoryResponse{
			ConversationIds: ids,
		}
		return resp, nil
	}
	conv, err := conversation.GetDefaultBucket(s.ctx).GetConversation(req.ConversationId, userId)
	if err != nil {
		return nil, err
	}
	history := make([]*llm.Message, 0)
	for _, m := range conv.Messages {
		history = append(history, &llm.Message{
			Role:    string(m.Role),
			Content: m.Content,
		})
	}
	resp = &llm.GetHistoryResponse{
		History: history,
	}
	return
}
