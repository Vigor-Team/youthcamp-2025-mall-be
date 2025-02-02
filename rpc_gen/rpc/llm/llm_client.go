package llm

import (
	"context"
	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm/llmservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() llmservice.Client
	Service() string
	SendMessage(ctx context.Context, Req *llm.ChatRequest, callOptions ...callopt.Option) (r *llm.ChatResponse, err error)
	StreamMessage(ctx context.Context, Req *llm.ChatRequest, callOptions ...callopt.Option) (stream llmservice.LlmService_StreamMessageClient, err error)
	GetHistory(ctx context.Context, Req *llm.GetHistoryRequest, callOptions ...callopt.Option) (r *llm.GetHistoryResponse, err error)
	DeleteHistory(ctx context.Context, Req *llm.DeleteHistoryRequest, callOptions ...callopt.Option) (r *llm.DeleteHistoryResponse, err error)
	GetConversationId(ctx context.Context, Req *llm.GetConversationIdRequest, callOptions ...callopt.Option) (r *llm.GetConversationIdResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := llmservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient llmservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() llmservice.Client {
	return c.kitexClient
}

func (c *clientImpl) SendMessage(ctx context.Context, Req *llm.ChatRequest, callOptions ...callopt.Option) (r *llm.ChatResponse, err error) {
	return c.kitexClient.SendMessage(ctx, Req, callOptions...)
}

func (c *clientImpl) StreamMessage(ctx context.Context, Req *llm.ChatRequest, callOptions ...callopt.Option) (stream llmservice.LlmService_StreamMessageClient, err error) {
	return c.kitexClient.StreamMessage(ctx, Req, callOptions...)
}

func (c *clientImpl) GetHistory(ctx context.Context, Req *llm.GetHistoryRequest, callOptions ...callopt.Option) (r *llm.GetHistoryResponse, err error) {
	return c.kitexClient.GetHistory(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteHistory(ctx context.Context, Req *llm.DeleteHistoryRequest, callOptions ...callopt.Option) (r *llm.DeleteHistoryResponse, err error) {
	return c.kitexClient.DeleteHistory(ctx, Req, callOptions...)
}

func (c *clientImpl) GetConversationId(ctx context.Context, Req *llm.GetConversationIdRequest, callOptions ...callopt.Option) (r *llm.GetConversationIdResponse, err error) {
	return c.kitexClient.GetConversationId(ctx, Req, callOptions...)
}
