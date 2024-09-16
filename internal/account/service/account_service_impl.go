package service

import (
	"context"
	"database/sql"
	"os"
	"time"

	"github.com/assidik12/go-restfull-api/helper"
	"github.com/assidik12/go-restfull-api/helper/exception"
	"github.com/assidik12/go-restfull-api/internal/account/repository"
	"github.com/assidik12/go-restfull-api/model/domain"
	"github.com/assidik12/go-restfull-api/model/web"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AccountServiceImpl struct {
	AccountRepository repository.AccountRepository
	DB                *sql.DB
	validate          *validator.Validate
}

func NewAccountService(accountRepository repository.AccountRepository, DB *sql.DB, validate *validator.Validate) *AccountServiceImpl {
	return &AccountServiceImpl{
		AccountRepository: accountRepository,
		DB:                DB,
		validate:          validate,
	}
}

func (s *AccountServiceImpl) Register(ctx context.Context, request web.AuthRegisterRequest) web.AuthRegisterResponse {
	err := s.validate.Struct(request)

	helper.PanicError(err)

	tx, err := s.DB.Begin()

	helper.PanicError(err)

	defer helper.CommitOrRollback(tx)

	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)

	helper.PanicError(err)

	account := domain.Account{
		Username: request.Username,
		Email:    request.Email,
		Password: string(bcryptPassword),
		Role:     "user",
	}

	account = s.AccountRepository.Save(ctx, tx, account)

	return helper.ToRegisterResponse(account)
}

func (s *AccountServiceImpl) Login(ctx context.Context, request web.AuthLoginRequest) web.AuthLoginResponse {
	err := s.validate.Struct(request)
	helper.PanicError(err)

	tx, err := s.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	account, err := s.AccountRepository.FindByEmail(ctx, tx, request.Email)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(request.Password)); err != nil {
		panic(exception.NewUnauthorizedError("password not match"))
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       account.ID,
		"username": account.Username,
		"email":    account.Email,
		"role":     account.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	})

	var key string = os.Getenv("AUTH_SECRET_KEY")

	token, err := jwtToken.SignedString([]byte(key))

	helper.PanicError(err)

	value := domain.AuthToken{
		Token: token,
	}
	return helper.ToLoginResponse(value)

}

func (s *AccountServiceImpl) Update(ctx context.Context, request web.AuthUpdateRequest) web.AuthUpdateResponse {
	err := s.validate.Struct(request)

	helper.PanicError(err)

	tx, err := s.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	account, err := s.AccountRepository.FindByEmail(ctx, tx, request.Email)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if request.Username != "" {
		account.Username = request.Username
	}
	if request.Email != "" {
		account.Email = request.Email
	}
	if request.Password != "" {
		bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
		helper.PanicError(err)
		account.Password = string(bcryptPassword)
	}

	account, err = s.AccountRepository.Update(ctx, tx, account)
	helper.PanicError(err)

	return helper.ToUpdateAccountResponse(account)

}
