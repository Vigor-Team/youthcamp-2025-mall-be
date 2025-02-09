package application

import (
	"context"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"testing"
)

func TestGetProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetProductService(ctx)
	// init req and assert value

	req := &product.GetProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
