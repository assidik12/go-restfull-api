package repository

import (
	"context"
	"database/sql"

	"github.com/assidik12/go-restfull-api/model/domain"
)

type AccountRepository interface {
	Save(ctx context.Context, tx *sql.Tx, account domain.Account) domain.Account
	Update(ctx context.Context, tx *sql.Tx, account domain.Account) (domain.Account, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.Account, error)
}
