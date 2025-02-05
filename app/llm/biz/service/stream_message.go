package service

import (
	"context"
	"errors"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/consts"
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
	userId := req.UserId
	convId := req.ConversationId
	if msg == "" || userId == "" {
		return nil, consts.ErrReqParamNotFound
	}
	if convId == "" {
		convId = uuid.New().String()
	}

	runnable, err := mallagent.BuildMallAgent(s.ctx, &mallagent.BuildConfig{MallAgent: &mallagent.MallAgentBuildConfig{}})
	if err != nil {
		klog.CtxErrorf(s.ctx, "build mall agent error: %v", err)
		return
	}

	bucket := conversation.GetDefaultBucket(s.ctx)
	conv, err := bucket.GetConversation(convId, userId)
	if err != nil {
		klog.CtxErrorf(s.ctx, "get conversation error: %v", err)
		return
	}

	userMessage := &mallagent.UserMessage{
		ConversationId: convId,
		Query:          msg,
		UserId:         userId,
		History:        conv.GetMessages(),
	}

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

			_ = conv.Append(schema.UserMessage(msg))

			fullMsg, err := schema.ConcatMessages(fullMsgs)
			if err != nil {
				klog.CtxErrorf(s.ctx, "error concatenating messages: %v", err)
			}
			_ = conv.Append(fullMsg)
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
