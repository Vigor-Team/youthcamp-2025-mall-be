package tools

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/infra/rpc"
	rpccart "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/cart"
	rpcorder "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
	"github.com/cloudwego/kitex/pkg/klog"
)

type CreateOrderTool struct {
}

func (cot *CreateOrderTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "create_order",
		Desc: "Create an order for a selected product, Please make sure that you have searched the id and price of the user's expected product through product search before calling",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"id": {
				Desc:     "The ID of the product",
				Type:     schema.Integer,
				Required: true,
			},
			"quantity": {
				Desc:     "The quantity of the product in order",
				Type:     schema.Integer,
				Required: true,
			},
			"price": {
				Desc:     "The price of the product in order",
				Type:     schema.Number,
				Required: true,
			},
		}),
	}, nil
}

func (cot *CreateOrderTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	var args struct {
		ProductID uint32  `json:"id"`
		Quantity  int32   `json:"quantity"`
		Price     float32 `json:"price"`
	}
	if err := json.Unmarshal([]byte(argumentsInJSON), &args); err != nil {
		klog.CtxErrorf(ctx, "Unmarshal arguments error: %v", err)
		return "", err
	}

	// Just Test
	var oi []*rpcorder.OrderItem
	oi = append(oi, &rpcorder.OrderItem{
		Item: &rpccart.CartItem{
			ProductId: uint32(args.ProductID),
			Quantity:  args.Quantity,
		},
		Cost: args.Price * float32(args.Quantity),
	})
	userId, err := strconv.ParseUint(ctx.Value("userId").(string), 10, 32)
	if err != nil {
		klog.CtxErrorf(ctx, "Parse user id error: %v", err)
		return "", err
	}
	order, err := rpc.OrderClient.PlaceOrder(ctx, &rpcorder.PlaceOrderReq{
		UserId: uint32(userId),
		// 模拟其他用户信息，实际应该询问用户
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
		klog.CtxErrorf(ctx, "Place order error: %v", err)
		return "", err
	}

	result := map[string]string{"order_id": order.Order.OrderId}
	resultJSON, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultJSON), nil
}

func NewCreateOrderTool(ctx context.Context) (tool.BaseTool, error) {
	return &CreateOrderTool{}, nil
}
