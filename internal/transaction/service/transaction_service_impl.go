package service

import (
	"context"
	"database/sql"

	"github.com/assidik12/go-restfull-api/helper"
	"github.com/assidik12/go-restfull-api/helper/exception"
	"github.com/assidik12/go-restfull-api/internal/transaction/repository"
	"github.com/assidik12/go-restfull-api/model/domain"
	"github.com/assidik12/go-restfull-api/model/web"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type TransactionServiceImpl struct {
	TransactionRepostory repository.TransactionRepository
	DB                   *sql.DB
	Validate             *validator.Validate
}

func NewTransactionService(transactionRepostory repository.TransactionRepository, DB *sql.DB, validate *validator.Validate) *TransactionServiceImpl {
	return &TransactionServiceImpl{
		TransactionRepostory: transactionRepostory,
		DB:                   DB,
		Validate:             validate,
	}
}

func (t *TransactionServiceImpl) FindAll(ctx context.Context, User_id int) []web.TransactionResponse {

	tx, err := t.DB.Begin()

	helper.PanicError(err)

	defer helper.CommitOrRollback(tx)

	result, err := t.TransactionRepostory.FindAll(ctx, tx, User_id)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	helper.PanicError(err)

	return helper.ToTransactionResponses(result)
}

func (t *TransactionServiceImpl) Create(ctx context.Context, request web.TransactionRequest, User_id int) web.TransactionResponse {
	err := t.Validate.Struct(request)

	helper.PanicError(err)

	tx, err := t.DB.Begin()

	helper.PanicError(err)

	defer helper.CommitOrRollback(tx)

	var transactions domain.Transaction

	for _, i := range request.Products {
		transactions.Product_id = append(transactions.Product_id, i.ID)
		transactions.Quantyty = append(transactions.Quantyty, i.Qty)
	}

	transactions.Transaction_detail_id = uuid.New().String()
	transactions.User_id = User_id
	transactions.Total_Price = request.TotalPrice

	t.TransactionRepostory.Create(ctx, tx, transactions)

	return helper.ToTransactionResponse(transactions)

}

func (t *TransactionServiceImpl) Delete(ctx context.Context, transaction_id int) {

	tx, err := t.DB.Begin()
	helper.PanicError(err)

	t.TransactionRepostory.Delete(ctx, tx, transaction_id)
}
