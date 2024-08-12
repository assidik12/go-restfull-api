package service

import (
	"context"
	"database/sql"

	"github.com/assidik12/go-restfull-api/helper"
	"github.com/assidik12/go-restfull-api/internal/product/repository"
	"github.com/assidik12/go-restfull-api/model/domain"
	"github.com/assidik12/go-restfull-api/model/web"
	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) *ProductServiceImpl {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (s *ProductServiceImpl) Create(ctx context.Context, request web.CreateProductRequest) web.ProductResponse {
	err := s.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := s.DB.Begin()
	helper.PanicError(err)

	defer helper.CommitOrRollback(tx)

	products := domain.Product{
		Name:        request.Name,
		Price:       request.Price,
		Stock:       request.Stock,
		Description: request.Description,
		Img:         request.Img,
		Category:    request.CategoryName,
	}

	products = s.ProductRepository.Create(ctx, tx, products)

	return helper.ToProductResponse(products)
}

func (s *ProductServiceImpl) Update(ctx context.Context, request web.UpdateProductRequest) web.ProductResponse {

	err := s.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := s.DB.Begin()
	helper.PanicError(err)

	defer helper.CommitOrRollback(tx)

	product, err := s.ProductRepository.FindById(ctx, tx, request.Id)

	helper.PanicError(err)

	product.Name = request.Name

	// ini harus di maping

	product = s.ProductRepository.Update(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (s *ProductServiceImpl) Delete(ctx context.Context, ProductId int) {
	tx, err := s.DB.Begin()

	helper.PanicError(err)

	defer helper.CommitOrRollback(tx)

	product, err := s.ProductRepository.FindById(ctx, tx, ProductId)

	helper.PanicError(err)

	s.ProductRepository.Delete(ctx, tx, product)

}

func (s *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {

	tx, err := s.DB.Begin()

	helper.PanicError(err)

	defer helper.CommitOrRollback(tx)

	product := s.ProductRepository.FindAll(ctx, tx)

	return helper.ToProductResponses(product)
}
func (s *ProductServiceImpl) FindById(ctx context.Context, CategoryId int) web.ProductResponse {
	tx, err := s.DB.Begin()

	helper.PanicError(err)

	defer helper.CommitOrRollback(tx)

	category, err := s.ProductRepository.FindById(ctx, tx, CategoryId)

	helper.PanicError(err)
	return helper.ToProductResponse(category)

}
