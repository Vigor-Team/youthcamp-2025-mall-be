package service

import (
	"context"
	"fmt"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/mallagent"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/mallagent/conversation"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/errno"
	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"github.com/cloudwego/eino/schema"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
)

type SendMessageService struct {
	ctx context.Context
} // NewSendMessageService new SendMessageService
func NewSendMessageService(ctx context.Context) *SendMessageService {
	return &SendMessageService{ctx: ctx}
}

// Run create note info
func (s *SendMessageService) Run(req *llm.ChatRequest) (resp *llm.ChatResponse, err error) {
	msg := req.Message
	userId := req.UserId
	convId := req.ConversationId
	fmt.Printf("convId: %v\n", convId)
	if msg == "" || userId == "" {
		return nil, kerrors.NewBizStatusError(errno.ErrGRPCRequestParam, "invalid params")
	}
	if convId == "" {
		convId = uuid.New().String()
	}

	s.ctx = context.WithValue(s.ctx, "userId", userId)
	runnable, err := mallagent.BuildMallAgent(s.ctx, &mallagent.BuildConfig{MallAgent: &mallagent.MallAgentBuildConfig{}})
	if err != nil {
		klog.CtxErrorf(s.ctx, "build mall agent error: %v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrBuildAgent, "build mall agent error")
	}

	bucket := conversation.GetDefaultBucket(s.ctx)
	conv, err := bucket.GetConversation(convId, userId)
	if err != nil {
		klog.CtxErrorf(s.ctx, "get conversation error: %v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrGetHistory, "get conversation error")
	}

	userMessage := &mallagent.UserMessage{
		ConversationId: convId,
		Query:          msg,
		UserId:         userId,
		History:        conv.GetMessages(),
	}
	rs, err := runnable.Invoke(s.ctx, userMessage)
	if err != nil {
		klog.CtxErrorf(s.ctx, "invoke mall agent error: %v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrInvokeAgent, "invoke mall agent error")
	}
	_ = conv.Append(schema.UserMessage(msg))
	_ = conv.Append(schema.AssistantMessage(rs.Content, rs.ToolCalls))
	resp = &llm.ChatResponse{
		Response: rs.Content,
	}
	return
}
