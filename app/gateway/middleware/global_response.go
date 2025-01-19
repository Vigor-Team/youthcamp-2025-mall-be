package middleware

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func GlobalResponseMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		c.Next(ctx)

		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if c.Response.StatusCode() == 0 {
			c.Response.SetStatusCode(http.StatusOK)
		}

		if c.Response.StatusCode() >= http.StatusOK && c.Response.StatusCode() < http.StatusMultipleChoices {
			data, exists := c.Get("data")
			if exists {
				utils.SuccessResponse(c, http.StatusOK, data)
			} else {
				utils.SuccessResponse(c, http.StatusOK, nil)
			}
			return
		}
	}
}
