package service

import (
	"context"

	"github.com/assidik12/go-restfull-api/model/web"
)

type CategoryService interface {
	FindAll(ctx context.Context) []web.CategoryResponse
	FindById(ctx context.Context, CategoryId int) web.CategoryResponse
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, CategoryId int)
}
