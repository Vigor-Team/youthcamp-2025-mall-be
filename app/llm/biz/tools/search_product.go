package tools

import (
	"context"
	"encoding/json"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/infra/rpc"
	rpcproduct "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

type SearchProductTool struct {
}

func (spt *SearchProductTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "search_product",
		Desc: `
			Search for products based on user query, and return the complete information about the product. 
			You can use it to recommend products to users, search for products, query product details, etc.
			You can also search for the product id needed to create an order based on the user's description when you need to create an order.
`,
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"query": {
				Desc:     "The search query for products, e.g. 'T-shirt'",
				Type:     schema.String,
				Required: true,
			},
		}),
	}, nil
}

func (spt *SearchProductTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	var args struct {
		Query string `json:"query"`
	}
	if err := json.Unmarshal([]byte(argumentsInJSON), &args); err != nil {
		return "", err
	}

	products, err := rpc.ProductClient.SearchProducts(ctx, &rpcproduct.SearchProductsReq{Query: args.Query})
	if err != nil {
		return "", err
	}

	result, err := json.Marshal(products)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func NewSearchProductTool(ctx context.Context) (tool.BaseTool, error) {
	return &SearchProductTool{}, nil
}
