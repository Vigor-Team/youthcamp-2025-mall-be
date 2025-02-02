package service

import (
	"context"
	"errors"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/mallagent"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/mallagent/conversation"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"github.com/cloudwego/eino/schema"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"io"
)

type StreamMessageService struct {
	ctx context.Context
}

// NewStreamMessageService new StreamMessageService
func NewStreamMessageService(ctx context.Context) *StreamMessageService {
	return &StreamMessageService{ctx: ctx}
}

func (s *StreamMessageService) Run(req *llm.ChatRequest) (sr *schema.StreamReader[*schema.Message], err error) {
	msg := req.Message

	runnable, err := mallagent.BuildMallAgent(s.ctx, &mallagent.BuildConfig{MallAgent: &mallagent.MallAgentBuildConfig{}})
	if err != nil {
		klog.CtxErrorf(s.ctx, "build mall agent error: %v", err)
		return
	}
	userMessage := &mallagent.UserMessage{
		ChatId: uuid.New().String(),
		Query:  msg,
		UserId: req.UserId,
	}

	memory := conversation.GetDefaultMemory(s.ctx)
	conversation := memory.GetConversation(userMessage.UserId, true)

	streamResult, err := runnable.Stream(s.ctx, userMessage)
	if err != nil {
		klog.CtxErrorf(s.ctx, "failed to stream: %v", err)
		return
	}

	srs := streamResult.Copy(2)

	go func() {
		fullMsgs := make([]*schema.Message, 0)
		defer func() {
			srs[1].Close()

			_ = conversation.Append(schema.UserMessage(msg))

			fullMsg, err := schema.ConcatMessages(fullMsgs)
			if err != nil {
				klog.CtxErrorf(s.ctx, "error concatenating messages: %v", err)
			}
			_ = conversation.Append(fullMsg)
		}()
	loop:
		for {
			select {
			case <-s.ctx.Done():
				klog.CtxInfof(s.ctx, "context done: %v", s.ctx.Err())
				break loop
			default:
				chunk, err := srs[1].Recv()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break loop
					}
				}
				fullMsgs = append(fullMsgs, chunk)
			}
		}
	}()
	return srs[0], nil
}
