package mallagent

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/tools"
	"github.com/cloudwego/eino/components/tool"
)

func GetTools(ctx context.Context) ([]tool.BaseTool, error) {

	searchProductTool, err := tools.NewSearchProductTool(ctx)
	if err != nil {
		return nil, err
	}
	createOrderTool, err := tools.NewCreateOrderTool(ctx)
	if err != nil {
		return nil, err
	}

	return []tool.BaseTool{
		searchProductTool,
		createOrderTool,
	}, nil
}
