package tools

import (
	"context"
	"encoding/json"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/consts"
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
		Desc: "Search for products based on user query",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"query": {
				Desc:     "The search query for products",
				Type:     schema.String,
				Required: true,
			},
			//"sort_by": {
			//	Desc:     "Sort products by price (asc or desc)",
			//	Type:     schema.String,
			//	Required: false,
			//},
		}),
	}, nil
}

func (spt *SearchProductTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	var args struct {
		Query string `json:"query"`
		//SortBy string `json:"sort_by"`
	}
	if err := json.Unmarshal([]byte(argumentsInJSON), &args); err != nil {
		return "", consts.ErrJsonUnmarshal
	}

	products, err := rpc.ProductClient.SearchProducts(ctx, &rpcproduct.SearchProductsReq{Query: args.Query})
	if err != nil {
		return "", consts.ErrSearchProducts
	}

	result, err := json.Marshal(products)
	if err != nil {
		return "", consts.ErrJsonMarshal
	}
	return string(result), nil
}
