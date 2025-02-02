package main

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/service"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"github.com/cloudwego/kitex/pkg/klog"
	"io"
)

// LlmServiceImpl implements the last service interface defined in the IDL.
type LlmServiceImpl struct{}

// SendMessage implements the LlmServiceImpl interface.
func (s *LlmServiceImpl) SendMessage(ctx context.Context, req *llm.ChatRequest) (resp *llm.ChatResponse, err error) {
	resp, err = service.NewSendMessageService(ctx).Run(req)

	return resp, err
}

func (s *LlmServiceImpl) StreamMessage(req *llm.ChatRequest, stream llm.LlmService_StreamMessageServer) (err error) {
	ctx := context.Background()
	sr, err := service.NewStreamMessageService(ctx).Run(req)
	for {
		select {
		case <-stream.Context().Done():
			return nil
		default:
			msg, err := sr.Recv()
			if err != nil {
				klog.CtxErrorf(stream.Context(), "stream read error: %v", err)
				if err == io.EOF {
					return nil
				}
			}
			resp := &llm.ChatResponse{
				Response: msg.Content,
			}

			if err := stream.Send(resp); err != nil {
				klog.CtxErrorf(stream.Context(), "stream send error: %v", err)
				return err
			}
		}
	}
}

// GetHistory implements the LlmServiceImpl interface.
func (s *LlmServiceImpl) GetHistory(ctx context.Context, req *llm.GetHistoryRequest) (resp *llm.GetHistoryResponse, err error) {
	resp, err = service.NewGetHistoryService(ctx).Run(req)

	return resp, err
}

// DeleteHistory implements the LlmServiceImpl interface.
func (s *LlmServiceImpl) DeleteHistory(ctx context.Context, req *llm.DeleteHistoryRequest) (resp *llm.DeleteHistoryResponse, err error) {
	resp, err = service.NewDeleteHistoryService(ctx).Run(req)

	return resp, err
}

// GetConversationId implements the LlmServiceImpl interface.
func (s *LlmServiceImpl) GetConversationId(ctx context.Context, req *llm.GetConversationIdRequest) (resp *llm.GetConversationIdResponse, err error) {
	resp, err = service.NewGetConversationIdService(ctx).Run(req)

	return resp, err
}
