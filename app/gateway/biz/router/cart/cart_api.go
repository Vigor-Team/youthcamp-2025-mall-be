// Code generated by hertz generator. DO NOT EDIT.

package cart

import (
	cart "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/biz/handler/cart"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_v1 := _api.Group("/v1", _v1Mw()...)
			_v1.POST("/carts", append(_addcartitemMw(), cart.AddCartItem)...)
			_carts := _v1.Group("/carts", _cartsMw()...)
			_carts.PUT("/{productId}", append(_updatecartitemMw(), cart.UpdateCartItem)...)
			_v1.GET("/carts", append(_getcartMw(), cart.GetCart)...)
		}
	}
}
