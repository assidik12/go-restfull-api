package service

import (
	"context"

	"github.com/assidik12/go-restfull-api/model/web"
)

type AccountService interface {
	Register(ctx context.Context, request web.AuthRegisterRequest) web.AuthRegisterResponse
	Login(ctx context.Context, request web.AuthLoginRequest) web.AuthLoginResponse
	Update(ctx context.Context, request web.AuthUpdateRequest) web.AuthUpdateResponse
}
