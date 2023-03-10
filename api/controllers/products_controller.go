package controllers

import (
	"encoding/json"
	"go-rest-example/api/models"
	"go-rest-example/api/repositories"
	"go-rest-example/api/utils"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductsController interface {
	Save(http.ResponseWriter, *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}

type productsControllerImpl struct {
	productsRepository repositories.ProductsRepository
	paginationBuilder  repositories.PaginationBuilderRepository
}

func NewProductsController(productsRepository repositories.ProductsRepository) *productsControllerImpl {
	return &productsControllerImpl{productsRepository, repositories.NewPaginationBuilderRepository(productsRepository)}
}

func (c *productsControllerImpl) Save(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	product := &models.Product{}
	err = json.Unmarshal(bytes, product)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = product.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	product, err = c.productsRepository.Save(product)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	buildCreatedResponse(w, buildLocation(r, product.ID))
	utils.WriteAsJson(w, product)
}

func (c *productsControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	product_id, err := strconv.ParseUint(params["product_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	product, err := c.productsRepository.FindById(product_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusNotFound)
		return
	}

	utils.WriteAsJson(w, product)
}

func (c *productsControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	meta, err := c.paginationBuilder.BuildProductsMetada(r)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	elements, err := c.productsRepository.Paginate(meta)
	if err != nil {
		utils.WriteError(w, err, http.StatusNotFound)
		return
	}

	utils.WriteAsJson(w, elements)
}

func (c *productsControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	product_id, err := strconv.ParseUint(params["product_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	product := &models.Product{}
	err = json.Unmarshal(bytes, product)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	product.ID = product_id

	err = product.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.productsRepository.Update(product)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	utils.WriteAsJson(w, map[string]bool{"success": err == nil})
}

func (c *productsControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	product_id, err := strconv.ParseUint(params["product_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.productsRepository.Delete(product_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusNotFound)
		return
	}

	buildDeleteResponse(w, product_id)
	utils.WriteAsJson(w, "{}")
}

func (c *productsControllerImpl) Search(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	name := params.Get("name")

	products, err := c.productsRepository.Search(name)
	if err != nil {
		utils.WriteError(w, err, http.StatusNotFound)
		return
	}

	utils.WriteAsJson(w, products)
}
