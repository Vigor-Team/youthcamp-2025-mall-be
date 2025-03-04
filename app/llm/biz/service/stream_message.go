package service

import (
	"context"
	"errors"
	"io"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/mallagent"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/mallagent/conversation"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/errno"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"github.com/cloudwego/eino/schema"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
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

	streamResult, err := runnable.Stream(s.ctx, userMessage)
	if err != nil {
		klog.CtxErrorf(s.ctx, "failed to stream: %v", err)
		return nil, kerrors.NewBizStatusError(consts.ErrInvokeAgent, "failed to stream")
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
