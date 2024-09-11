package service

import (
	"database/sql"

	"context"

	"github.com/assidik12/go-restfull-api/helper"
	"github.com/assidik12/go-restfull-api/helper/exception"
	"github.com/assidik12/go-restfull-api/internal/category/repository"
	"github.com/assidik12/go-restfull-api/model/domain"
	"github.com/assidik12/go-restfull-api/model/web"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (s *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := s.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := s.DB.Begin()
	helper.PanicError(err)

	defer helper.CommitOrRollback(tx)

	categories := domain.Category{
		Name: request.Name,
	}

	categories = s.CategoryRepository.Create(ctx, tx, categories)

	return helper.ToCategoryResponse(categories)
}

func (s *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {

	err := s.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := s.DB.Begin()
	helper.PanicError(err)

	defer helper.CommitOrRollback(tx)

	categories, err := s.CategoryRepository.FindById(ctx, tx, request.Id)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	categories.Name = request.Name

	categories = s.CategoryRepository.Update(ctx, tx, categories)

	return helper.ToCategoryResponse(categories)
}

func (s *CategoryServiceImpl) Delete(ctx context.Context, CategoryId int) {
	tx, err := s.DB.Begin()

	helper.PanicError(err)

	defer helper.CommitOrRollback(tx)

	categories, err := s.CategoryRepository.FindById(ctx, tx, CategoryId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	s.CategoryRepository.Delete(ctx, tx, categories)

}

func (s *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {

	tx, err := s.DB.Begin()

	helper.PanicError(err)

	defer helper.CommitOrRollback(tx)

	categories := s.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
func (s *CategoryServiceImpl) FindById(ctx context.Context, CategoryId int) web.CategoryResponse {
	tx, err := s.DB.Begin()

	helper.PanicError(err)

	defer helper.CommitOrRollback(tx)

	category, err := s.CategoryRepository.FindById(ctx, tx, CategoryId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)

}
