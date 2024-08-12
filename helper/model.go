package helper

import (
	"github.com/assidik12/go-restfull-api/model/domain"
	"github.com/assidik12/go-restfull-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

func ToCategoryResponses(catagory []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse

	for _, category := range catagory {
		categoryResponses = append(categoryResponses, web.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}
	return categoryResponses
}
func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Stock:       product.Stock,
		Description: product.Description,
		Img:         product.Img,
	}
}

func ToProductResponses(product []domain.Product) []web.ProductResponse {
	var productResponses []web.ProductResponse

	for _, product := range product {
		productResponses = append(productResponses, web.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Stock:       product.Stock,
			Description: product.Description,
			Img:         product.Img,
		})
	}
	return productResponses
}
