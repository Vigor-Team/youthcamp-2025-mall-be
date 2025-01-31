package llm

import (
	"context"
	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm/llmservice"

	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func SendMessage(ctx context.Context, req *llm.ChatRequest, callOptions ...callopt.Option) (resp *llm.ChatResponse, err error) {
	resp, err = defaultClient.SendMessage(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SendMessage call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func StreamMessage(ctx context.Context, Req *llm.ChatRequest, callOptions ...callopt.Option) (stream llmservice.LlmService_StreamMessageClient, err error) {
	stream, err = defaultClient.StreamMessage(ctx, Req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "StreamMessage call failed,err =%+v", err)
		return nil, err
	}
	return stream, nil
}
