package repository

import (
	"context"
	"database/sql"

	"github.com/assidik12/go-restfull-api/model/domain"
)

type TransactionRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx, user_id int) ([]domain.TransactionDetail, error)
	Create(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction
	Delete(ctx context.Context, tx *sql.Tx, transaction_id int)
}
