package application

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/infras/repository"
	product "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"testing"
)

func TestOnlineProduct_Run(t *testing.T) {
	repository.Init()
	ctx := context.Background()
	s := NewOnlineProductService(ctx)
	// init req and assert value

	req := &product.OnlineProductReq{
		Id: 2801799168,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
