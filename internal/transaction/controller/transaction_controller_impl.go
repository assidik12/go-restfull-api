package controller

import (
	"net/http"
	"strconv"

	"github.com/assidik12/go-restfull-api/helper"
	"github.com/assidik12/go-restfull-api/internal/transaction/service"
	"github.com/assidik12/go-restfull-api/model/web"
	"github.com/julienschmidt/httprouter"
)

type TransactionControllerImpl struct {
	TransactionService service.TransactionService
}

func NewTransactionController(TransactionService service.TransactionService) *TransactionControllerImpl {
	return &TransactionControllerImpl{
		TransactionService: TransactionService,
	}
}

func (controller *TransactionControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	UserId := params.ByName("userId")
	id, err := strconv.Atoi(UserId)
	helper.PanicError(err)

	transactionResponses := controller.TransactionService.FindAll(request.Context(), id)

	webResponse := web.WebResponse{
		Code:    http.StatusOK,
		Message: "transaction fetched",
		Data:    transactionResponses,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *TransactionControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	transactionCreateRequest := web.TransactionRequest{}

	UserId := params.ByName("userId")
	id, err := strconv.Atoi(UserId)
	helper.PanicError(err)

	helper.ReadRequestBody(request, &transactionCreateRequest)

	transactionResponse := controller.TransactionService.Create(request.Context(), transactionCreateRequest, id)
	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Message: "Transaction Created",
		Data:    transactionResponse,
	}

	helper.WriteResponseBody(writer, webResponse)

}

func (controller *TransactionControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	transactionID := params.ByName("transactionId")

	id, err := strconv.Atoi(transactionID)
	helper.PanicError(err)

	controller.TransactionService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:    http.StatusAccepted,
		Message: "Transaction Deleted",
	}

	helper.WriteResponseBody(writer, webResponse)

}
