package service

import (
	"context"

	"github.com/assidik12/go-restfull-api/model/web"
)

type ProductService interface {
	FindAll(ctx context.Context) []web.ProductResponse
	FindById(ctx context.Context, ProductId int) web.ProductResponse
	Create(ctx context.Context, request web.CreateProductRequest) web.ProductResponse
	Update(ctx context.Context, request web.UpdateProductRequest) web.ProductResponse
	Delete(ctx context.Context, ProductId int)
}
