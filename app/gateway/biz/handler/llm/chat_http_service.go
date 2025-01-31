package llm

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/service"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/utils"
	llm "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/hertz_gen/gateway/llm"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sse"
	"io"
	"net/http"
)

// SendMessage .
// @router /v1/chat/send [POST]
func SendMessage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req llm.ChatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.ErrorResponse(c, consts.StatusOK, err.Error())
		return
	}

	resp := &llm.ChatResponse{}
	resp, err = service.NewSendMessageService(ctx, c).Run(&req)
	if err != nil {
		utils.ErrorResponse(c, consts.StatusOK, err.Error())
		return
	}

	utils.SuccessResponse(c, resp)
}

// StreamMessage .
// @router /v1/chat/stream [POST]
func StreamMessage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req llm.ChatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.ErrorResponse(c, consts.StatusOK, err.Error())
		return
	}
	c.Response.Header.Set("X-Accel-Buffering", "no")

	client, err := service.NewStreamMessageService(ctx, c).Run(&req)
	if err != nil {
		utils.ErrorResponse(c, consts.StatusOK, err.Error())
		return
	}
	err = client.Close()
	if err != nil {
		return
	}

	c.SetStatusCode(http.StatusOK)
	s := sse.NewStream(c)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			resp, err := client.Recv()
			if err != nil {
				if err == io.EOF {
					return
				}
				hlog.CtxErrorf(ctx, "stream recv error: %v", err)
				return
			}
			event := &sse.Event{
				Event: "chat",
				Data:  []byte(resp.GetResponse()),
			}
			err = s.Publish(event)
			if err != nil {
				return
			}
		}
	}
}
