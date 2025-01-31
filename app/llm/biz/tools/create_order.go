package tools

import (
	"context"
	"encoding/json"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/consts"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/infra/rpc"
	rpccart "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/cart"
	rpcorder "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

type CreateOrderTool struct {
}

func (cot *CreateOrderTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "create_order",
		Desc: "Create an order for a selected product",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"product_id": {
				Desc:     "The ID of the product to order",
				Type:     schema.String,
				Required: true,
			},
			"quantity": {
				Desc:     "The quantity of the product to order",
				Type:     schema.Integer,
				Required: true,
			},
			"user_id": {
				Desc:     "The ID of the user who is ordering",
				Type:     schema.Integer,
				Required: true,
			},
		}),
	}, nil
}

func (cot *CreateOrderTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	var args struct {
		UserId    uint32 `json:"user_id"`
		ProductID uint32 `json:"product_id"`
		Quantity  int32  `json:"quantity"`
	}
	if err := json.Unmarshal([]byte(argumentsInJSON), &args); err != nil {
		return "", consts.ErrJsonUnmarshal
	}

	// Just Test
	var oi []*rpcorder.OrderItem
	oi = append(oi, &rpcorder.OrderItem{
		Item: &rpccart.CartItem{
			ProductId: args.ProductID,
			Quantity:  args.Quantity,
		},
		Cost: 1 * float32(args.Quantity),
	})
	order, err := rpc.OrderClient.PlaceOrder(ctx, &rpcorder.PlaceOrderReq{
		UserId:       args.UserId,
		UserCurrency: "USD",
		Address: &rpcorder.Address{
			StreetAddress: "123 Main St",
			City:          "San Francisco",
			State:         "CA",
			ZipCode:       94105,
			Country:       "USA",
		},
		OrderItems: oi,
	})

	if err != nil {
		return "", consts.ErrCreateOrder
	}

	result := map[string]string{"order_id": order.Order.OrderId}
	resultJSON, err := json.Marshal(result)
	if err != nil {
		return "", consts.ErrJsonMarshal
	}
	return string(resultJSON), nil
}
