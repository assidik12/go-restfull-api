package controller

import (
	"net/http"

	"github.com/assidik12/go-restfull-api/helper"
	"github.com/assidik12/go-restfull-api/internal/account/service"
	"github.com/assidik12/go-restfull-api/model/web"
	"github.com/julienschmidt/httprouter"
)

type AccountControllerImpl struct {
	AccountService service.AccountService
}

func NewAccountController(accountService service.AccountService) *AccountControllerImpl {
	return &AccountControllerImpl{
		AccountService: accountService,
	}
}

func (controller *AccountControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	registerRequest := web.AuthRegisterRequest{}

	helper.ReadRequestBody(request, &registerRequest)
	accountResponser := controller.AccountService.Register(request.Context(), registerRequest)
	responseWeb := web.WebResponse{
		Code:    http.StatusCreated,
		Message: "Account has been created",
		Data:    accountResponser,
	}

	helper.WriteResponseBody(writer, responseWeb)
}
func (controller *AccountControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginRequest := web.AuthLoginRequest{}

	helper.ReadRequestBody(request, &loginRequest)

	accountResponser := controller.AccountService.Login(request.Context(), loginRequest)

	responseWeb := web.WebResponse{
		Code:    http.StatusAccepted,
		Message: "Account has been login",
		Data:    accountResponser,
	}

	helper.WriteResponseBody(writer, responseWeb)

}
func (controller *AccountControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	updateRequest := web.AuthUpdateRequest{}
	helper.ReadRequestBody(request, &updateRequest)

	accountResponser := controller.AccountService.Update(request.Context(), updateRequest)

	responseWeb := web.WebResponse{
		Code:    http.StatusAccepted,
		Message: "Account has been updated",
		Data:    accountResponser,
	}

	helper.WriteResponseBody(writer, responseWeb)
}
