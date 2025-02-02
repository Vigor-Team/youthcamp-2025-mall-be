package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/mallagent"
	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"github.com/cloudwego/kitex/pkg/klog"
)

type SendMessageService struct {
	ctx context.Context
} // NewSendMessageService new SendMessageService
func NewSendMessageService(ctx context.Context) *SendMessageService {
	return &SendMessageService{ctx: ctx}
}

// Run create note info
func (s *SendMessageService) Run(req *llm.ChatRequest) (resp *llm.ChatResponse, err error) {
	// todo: chat history
	msg := req.Message
	userId := req.UserId
	convId := req.ConversationId
	if msg == "" || userId == "" || convId == "" {
		return nil, consts.ErrReqParamNotFound
	}

	runnable, err := mallagent.BuildMallAgent(s.ctx, &mallagent.BuildConfig{MallAgent: &mallagent.MallAgentBuildConfig{}})
	if err != nil {
		klog.CtxErrorf(s.ctx, "build mall agent error: %v", err)
		return
	}
	userMessage := &mallagent.UserMessage{
		ConversationId: convId,
		Query:          msg,
		UserId:         userId,
	}
	rs, err := runnable.Invoke(s.ctx, userMessage)
	if err != nil {
		klog.CtxErrorf(s.ctx, "invoke mall agent error: %v", err)
		return
	}
	resp = &llm.ChatResponse{
		Response: rs.Content,
	}
	return
}
