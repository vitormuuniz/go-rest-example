package routes

import (
	"go-rest-example/api/controllers"
	"net/http"
)

type ProductRoutes interface {
	Routes() []*Route
}

type productRoutesImpl struct {
	productsController controllers.ProductsController
}

func NewProductRoutes(productsController controllers.ProductsController) *productRoutesImpl {
	return &productRoutesImpl{productsController}
}

func (r *productRoutesImpl) Routes() []*Route {
	return []*Route{
		{
			Path:    "/products",
			Method:  http.MethodPost,
			Handler: r.productsController.Save,
		},
		{
			Path:    "/products/{product_id}",
			Method:  http.MethodGet,
			Handler: r.productsController.FindById,
		},
		{
			Path:    "/products",
			Method:  http.MethodGet,
			Handler: r.productsController.FindAll,
		},
		{
			Path:    "/products/{product_id}",
			Method:  http.MethodPut,
			Handler: r.productsController.Update,
		},
		{
			Path:    "/products/{product_id}",
			Method:  http.MethodDelete,
			Handler: r.productsController.Delete,
		},
		{
			Path:    "/search/products",
			Method:  http.MethodGet,
			Handler: r.productsController.Search,
		},
	}
}
