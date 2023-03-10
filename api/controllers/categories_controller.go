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

type CategoriesController interface {
	Post(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type categoriesControllerImpl struct {
	categoriesRepository repositories.CategoriesRepository
}

func NewCategoriesController(categoriesRepository repositories.CategoriesRepository) *categoriesControllerImpl {
	return &categoriesControllerImpl{categoriesRepository}
}

func (c *categoriesControllerImpl) Post(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	category := &models.Category{}
	err = json.Unmarshal(bytes, category)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = category.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	category, err = c.categoriesRepository.Save(category)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	buildCreatedResponse(w, buildLocation(r, category.ID))
	utils.WriteAsJson(w, category)
}

func (c *categoriesControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	category_id, err := strconv.ParseUint(params["category_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	category, err := c.categoriesRepository.FindById(category_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusNotFound)
		return
	}

	utils.WriteAsJson(w, category)
}

func (c *categoriesControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	categories, err := c.categoriesRepository.FindAll()
	if err != nil {
		utils.WriteError(w, err, http.StatusNotFound)
		return
	}

	utils.WriteAsJson(w, categories)
}

func (c *categoriesControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	category_id, err := strconv.ParseUint(params["category_id"], 10, 64)
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

	category := &models.Category{}
	err = json.Unmarshal(bytes, category)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	category.ID = category_id

	err = category.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.categoriesRepository.Update(category)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	utils.WriteAsJson(w, map[string]bool{"success": err == nil})
}

func (c *categoriesControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	category_id, err := strconv.ParseUint(params["category_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.categoriesRepository.Delete(category_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusNotFound)
		return
	}

	buildDeleteResponse(w, category_id)
	utils.WriteAsJson(w, "{}")
}
