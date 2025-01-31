package service

import (
	"context"
	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
	"time"
)

type StreamMessageService struct {
	ctx context.Context
}

// NewStreamMessageService new StreamMessageService
func NewStreamMessageService(ctx context.Context) *StreamMessageService {
	return &StreamMessageService{ctx: ctx}
}

func (s *StreamMessageService) Run(req *llm.ChatRequest, stream llm.LlmService_StreamMessageServer) (err error) {
	//streamReader := chat.StreamLlmResponse(stream.Context(), req.Message)

	//for {
	//	select {
	//	case <-stream.Context().Done():
	//		return nil
	//	default:
	//		//msg, err := streamReader.Recv()
	//		//if err != nil {
	//		//	klog.CtxErrorf(stream.Context(), "stream read error: %v", err)
	//		//	if err == io.EOF {
	//		//		return nil
	//		//	}
	//		//}
	//		resp := &llm.ChatResponse{
	//			Response: "1",
	//		}
	//
	//		fmt.Println("resp: ", resp)
	//		if err := stream.Send(resp); err != nil {
	//			klog.CtxErrorf(stream.Context(), "stream send error: %v", err)
	//			return err
	//		}
	//	}
	//}
	counter := 0 // 初始化计数器

	for {
		select {
		case <-stream.Context().Done():
			klog.CtxInfof(stream.Context(), "client disconnected")
			return nil
		default:
			counter++
			if counter > 10 {
				return nil
			}

			// 构造响应数据
			resp := &llm.ChatResponse{
				Response: strconv.Itoa(counter),
			}

			klog.CtxInfof(stream.Context(), "sending chunk: %d", counter)

			if err := stream.Send(resp); err != nil {
				klog.CtxErrorf(stream.Context(), "stream send error: %v", err)
				return err
			}

			time.Sleep(500 * time.Millisecond)
		}
	}
}
