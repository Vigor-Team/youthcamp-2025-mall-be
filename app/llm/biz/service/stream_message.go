package service

import (
	"context"
	"fmt"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/chat"
	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"github.com/cloudwego/kitex/pkg/klog"
	"io"
)

type StreamMessageService struct {
	ctx context.Context
}

// NewStreamMessageService new StreamMessageService
func NewStreamMessageService(ctx context.Context) *StreamMessageService {
	return &StreamMessageService{ctx: ctx}
}

func (s *StreamMessageService) Run(req *llm.ChatRequest, stream llm.LlmService_StreamMessageServer) (err error) {
	streamReader := chat.StreamLlmResponse(stream.Context(), req.Message)

	for {
		select {
		case <-stream.Context().Done():
			return nil
		default:
			msg, err := streamReader.Recv()
			if err != nil {
				klog.CtxErrorf(stream.Context(), "stream read error: %v", err)
				if err == io.EOF {
					return nil
				}
			}
			resp := &llm.ChatResponse{
				Response: msg.Content,
			}

			fmt.Println("resp: ", resp)
			if err := stream.Send(resp); err != nil {
				klog.CtxErrorf(stream.Context(), "stream send error: %v", err)
				return err
			}
		}
	}

}
