package service

import (
	"context"

	"github.com/assidik12/go-restfull-api/model/web"
)

type TransactionService interface {
	FindAll(ctx context.Context, User_id int) []web.TransactionResponse
	Create(ctx context.Context, request web.TransactionRequest, User_id int) web.TransactionResponse
	Delete(ctx context.Context, TransactionId int)
}
