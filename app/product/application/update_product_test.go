package application

import (
	"context"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"testing"
)

func TestUpdateProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateProductService(ctx)
	// init req and assert value

	req := &product.UpdateProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
