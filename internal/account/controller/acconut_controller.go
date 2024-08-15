package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AccountController interface {
	Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
