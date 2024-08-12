package controller

import (
	"net/http"
	"strconv"

	"github.com/assidik12/go-restfull-api/helper"
	"github.com/assidik12/go-restfull-api/internal/product/service"
	"github.com/assidik12/go-restfull-api/model/web"
	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(ProductService service.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{
		ProductService: ProductService,
	}
}

func (controller *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	productCreateRequest := web.CreateProductRequest{}

	helper.ReadRequestBody(request, &productCreateRequest)

	categoryResponse := controller.ProductService.Create(request.Context(), productCreateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    categoryResponse,
	}

	helper.WriteResponseBody(writer, webResponse)

}

func (controller *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	UpdateProductRequest := web.UpdateProductRequest{}

	helper.ReadRequestBody(request, &UpdateProductRequest)

	productID := params.ByName("productId")
	id, err := strconv.Atoi(productID)
	helper.PanicError(err)
	UpdateProductRequest.Id = id

	categoryResponse := controller.ProductService.Update(request.Context(), UpdateProductRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    categoryResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}
func (controller *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	productID := params.ByName("productId")
	id, err := strconv.Atoi(productID)

	helper.PanicError(err)

	controller.ProductService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "OK",
	}

	helper.WriteResponseBody(writer, webResponse)
}
func (controller *ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	ProductResponse := controller.ProductService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    ProductResponse,
	}

	helper.WriteResponseBody(writer, webResponse)
}
func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productID := params.ByName("productId")
	id, err := strconv.Atoi(productID)
	helper.PanicError(err)

	CategoryResponse := controller.ProductService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    CategoryResponse,
	}
	helper.WriteResponseBody(writer, webResponse)
}
