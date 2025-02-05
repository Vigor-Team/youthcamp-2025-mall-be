package mallagent

import (
	"context"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
)

const (
	InputToQuery   = "InputToQuery"
	ChatTemplate   = "ChatTemplate"
	ReactAgent     = "ReactAgent"
	InputToHistory = "InputToHistory"
	GraphName      = "MallAgent"
)

type MallAgentBuildConfig struct {
	ChatTemplateKeyOfChatTemplate *ChatTemplateConfig
	ReactAgentKeyOfLambda         *react.AgentConfig
}

type BuildConfig struct {
	MallAgent *MallAgentBuildConfig
}

func BuildMallAgent(ctx context.Context, config *BuildConfig) (r compose.Runnable[*UserMessage, *schema.Message], err error) {
	g := compose.NewGraph[*UserMessage, *schema.Message]()

	err = g.AddLambdaNode(InputToHistory, compose.InvokableLambdaWithOption(NewInputToHistory),
		compose.WithNodeName("UserMessageToVariables"))
	if err != nil {
		return nil, err
	}

	chatTemplateKeyOfChatTemplate, err := NewChatTemplate(ctx, config.MallAgent.ChatTemplateKeyOfChatTemplate)
	if err != nil {
		return nil, err
	}
	err = g.AddChatTemplateNode(ChatTemplate, chatTemplateKeyOfChatTemplate)
	if err != nil {
		return nil, err
	}

	reactAgentKeyOfLambda, err := NewReactAgent(ctx, config.MallAgent.ReactAgentKeyOfLambda)
	if err != nil {
		return nil, err
	}

	err = g.AddLambdaNode(ReactAgent, reactAgentKeyOfLambda, compose.WithNodeName("ReAct Agent"))
	if err != nil {
		return nil, err
	}

	_ = g.AddEdge(compose.START, InputToHistory)
	_ = g.AddEdge(InputToHistory, ChatTemplate)
	_ = g.AddEdge(ChatTemplate, ReactAgent)
	_ = g.AddEdge(ReactAgent, compose.END)

	r, err = g.Compile(ctx, compose.WithGraphName(GraphName), compose.WithNodeTriggerMode(compose.AllPredecessor))
	if err != nil {
		return nil, err
	}
	return
}
