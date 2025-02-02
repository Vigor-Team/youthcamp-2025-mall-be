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

func GetHistory(ctx context.Context, req *llm.GetHistoryRequest, callOptions ...callopt.Option) (resp *llm.GetHistoryResponse, err error) {
	resp, err = defaultClient.GetHistory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetHistory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteHistory(ctx context.Context, req *llm.DeleteHistoryRequest, callOptions ...callopt.Option) (resp *llm.DeleteHistoryResponse, err error) {
	resp, err = defaultClient.DeleteHistory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteHistory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetConversationId(ctx context.Context, req *llm.GetConversationIdRequest, callOptions ...callopt.Option) (resp *llm.GetConversationIdResponse, err error) {
	resp, err = defaultClient.GetConversationId(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetConversationId call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
