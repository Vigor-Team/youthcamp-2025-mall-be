package mallagent

import (
	"context"

	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
)

func defaultReactAgentConfig(ctx context.Context) (config *react.AgentConfig, err error) {
	chatModel, err := NewArkChatModel(ctx, nil)
	if err != nil {
		return
	}
	tools, err := GetTools(ctx)

	config = &react.AgentConfig{
		Model:              chatModel,
		MaxStep:            25,
		ToolReturnDirectly: map[string]struct{}{},
		ToolsConfig: compose.ToolsNodeConfig{
			Tools: tools,
		},
	}
	return
}

func NewReactAgent(ctx context.Context, config *react.AgentConfig) (lba *compose.Lambda, err error) {
	if config == nil {
		config, err = defaultReactAgentConfig(ctx)
		if err != nil {
			return nil, err
		}
	}

	ins, err := react.NewAgent(ctx, config)
	if err != nil {
		return nil, err
	}
	lba, err = compose.AnyLambda(ins.Generate, ins.Stream, nil, nil)
	if err != nil {
		return nil, err
	}
	return lba, nil
}
