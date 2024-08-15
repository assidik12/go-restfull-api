package exception

import (
	"net/http"

	"github.com/assidik12/go-restfull-api/helper"
	"github.com/assidik12/go-restfull-api/model/web"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if untauthorizedError(writer, request, err) {
		return
	}

	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func untauthorizedError(writer http.ResponseWriter, _ *http.Request, err interface{}) bool {
	exception, ok := err.(UnauthorizedError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Message: exception.Message,
			Data:    nil,
		}
		helper.WriteResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, _ *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		webResponse := web.WebResponse{
			Code:    http.StatusNotFound,
			Message: exception.Error,
			Data:    nil,
		}
		helper.WriteResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func validationError(writer http.ResponseWriter, _ *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		webResponse := web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: exception.Error(),
			Data:    nil,
		}
		helper.WriteResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, _ *http.Request, err interface{}) {
	writer.Header().Set("Contect-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:    http.StatusInternalServerError,
		Message: "INTERNAL SERVER ERROR",
		Data:    err,
	}
	helper.WriteResponseBody(writer, webResponse)
}
