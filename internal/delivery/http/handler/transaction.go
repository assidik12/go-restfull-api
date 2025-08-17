package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/assidik12/go-restfull-api/internal/delivery/http/dto"
	"github.com/assidik12/go-restfull-api/internal/pkg/response"
	"github.com/assidik12/go-restfull-api/internal/service"
	"github.com/julienschmidt/httprouter"
)

type TransactionHandler struct {
	service service.TrancationService
}

func NewTransactionHandler(service service.TrancationService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

// Implement methods for TransactionHandler as needed
func (th *TransactionHandler) GetAllTransaction(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id := r.Header.Get("user_id")

	// Logic to get all transactions
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.BadRequest(w, "Invalid user ID")
		return
	}
	transactions, err := th.service.GetAll(r.Context(), idInt)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	// Return response
	response.OK(w, transactions)
}

func (th *TransactionHandler) GetTransactionById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.BadRequest(w, "Invalid transaction ID")
		return
	}
	transaction, err := th.service.FindById(r.Context(), idInt)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}
	response.OK(w, transaction)
}

func (th *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	userId := r.Header.Get("user_id")
	if userId == "" {
		response.BadRequest(w, "Invalid user ID")
		return
	}
	idInt, err := strconv.Atoi(userId)
	if err != nil {
		response.BadRequest(w, "Invalid user ID")
		return
	}
	var req dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	transaction, err := th.service.Save(r.Context(), req, idInt)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}
	response.Created(w, transaction)
}
