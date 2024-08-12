package controller

import (
	"net/http"
	"strconv"

	"github.com/assidik12/go-restfull-api/helper"
	"github.com/assidik12/go-restfull-api/internal/category/service"
	"github.com/assidik12/go-restfull-api/model/web"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryCreateRequest := web.CategoryCreateRequest{}

	helper.ReadRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    categoryResponse,
	}

	helper.WriteResponseBody(writer, webResponse)

}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	CategoryUpdateRequest := web.CategoryUpdateRequest{}

	helper.ReadRequestBody(request, &CategoryUpdateRequest)

	categoryID := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryID)
	helper.PanicError(err)
	CategoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), CategoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    categoryResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryID := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryID)
	helper.PanicError(err)

	controller.CategoryService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "OK",
	}

	helper.WriteResponseBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	CategoryResponse := controller.CategoryService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    CategoryResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryID := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryID)
	helper.PanicError(err)

	CategoryResponse := controller.CategoryService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    CategoryResponse,
	}
	helper.WriteResponseBody(writer, webResponse)
}
