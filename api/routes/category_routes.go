package routes

import (
	"go-rest-example/api/controllers"
	"net/http"
)

type CategoryRoutes interface {
	Routes() []*Route
}

type categoryRoutesImpl struct {
	categoriesController controllers.CategoriesController
}

func NewCategoryRoutes(categoriesController controllers.CategoriesController) *categoryRoutesImpl {
	return &categoryRoutesImpl{categoriesController}
}

func (r *categoryRoutesImpl) Routes() []*Route {
	return []*Route{
		{
			Path:    "/categories",
			Method:  http.MethodPost,
			Handler: r.categoriesController.Post,
		},
		{
			Path:    "/categories/{category_id}",
			Method:  http.MethodGet,
			Handler: r.categoriesController.FindById,
		},
		{
			Path:    "/categories",
			Method:  http.MethodGet,
			Handler: r.categoriesController.FindAll,
		},
		{
			Path:    "/categories/{category_id}",
			Method:  http.MethodPut,
			Handler: r.categoriesController.Update,
		},
		{
			Path:    "/categories/{category_id}",
			Method:  http.MethodDelete,
			Handler: r.categoriesController.Delete,
		},
	}
}
