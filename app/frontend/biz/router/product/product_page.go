// Code generated by hertz generator. DO NOT EDIT.

package product

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	product "github.com/zheyuanf/ecommerce-tiktok/app/frontend/biz/handler/product"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.GET("/product", append(_getproductMw(), product.GetProduct)...)
	_product := root.Group("/product", _productMw()...)
	_product.PUT("/:id", append(_updateproductMw(), product.UpdateProduct)...)
	_product.DELETE("/:id", append(_deleteproductMw(), product.DeleteProduct)...)
	root.POST("/product", append(_createproductMw(), product.CreateProduct)...)
	root.GET("/search", append(_searchproductsMw(), product.SearchProducts)...)
}
