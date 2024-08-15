package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/assidik12/go-restfull-api/helper"
	"github.com/assidik12/go-restfull-api/model/domain"
)

type AccountRepositoryImpl struct{}

func NewAccountRepository() *AccountRepositoryImpl {
	return &AccountRepositoryImpl{}
}

func (r *AccountRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, account domain.Account) domain.Account {
	SQL := "INSERT INTO account(username, email, password) VALUES(?, ?, ?)"

	result, err := tx.ExecContext(ctx, SQL, account.Username, account.Email, account.Password)
	helper.PanicError(err)

	id, err := result.LastInsertId()
	helper.PanicError(err)

	account.ID = int(id)
	return account
}

func (r *AccountRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, account domain.Account) (domain.Account, error) {
	SQL := "UPDATE account SET "

	var updateValues []string
	var updateArguments []interface{}

	if account.Username != "" {
		updateValues = append(updateValues, "username = ?")
		updateArguments = append(updateArguments, account.Username)
	}

	if account.Email != "" {
		updateValues = append(updateValues, "email = ?")
		updateArguments = append(updateArguments, account.Email)
	}

	if account.Password != "" {
		updateValues = append(updateValues, "password = ?")
		updateArguments = append(updateArguments, account.Password)
	}

	if len(updateValues) == 0 {
		return domain.Account{}, errors.New("missing update value")
	}

	updateSQL := SQL + strings.Join(updateValues, ", ")
	updateSQL += " WHERE email = ?"
	updateArguments = append(updateArguments, account.Email)

	_, err := tx.ExecContext(ctx, updateSQL, updateArguments...)
	helper.PanicError(err)

	return account, nil

}

func (r *AccountRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.Account, error) {

	SQL := "SELECT id, username, email, password FROM account WHERE email = ?"
	rows, err := tx.QueryContext(ctx, SQL, email)

	helper.PanicError(err)

	defer rows.Close()

	account := domain.Account{}
	if rows.Next() {
		err := rows.Scan(&account.ID, &account.Username, &account.Email, &account.Password)
		helper.PanicError(err)
		return account, nil

	} else {
		return account, errors.New("account not found")
	}
}
