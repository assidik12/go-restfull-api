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

func ToLoginResponse(account domain.AuthToken) web.AuthLoginResponse {
	return web.AuthLoginResponse{
		Token: account.Token,
	}
}

func ToRegisterResponse(account domain.Account) web.AuthRegisterResponse {
	return web.AuthRegisterResponse{
		Message: "register success",
	}
}

func ToUpdateAccountResponse(account domain.Account) web.AuthUpdateResponse {
	return web.AuthUpdateResponse{
		Message: "update success",
	}
}

func ToTransactionResponse(transaction domain.Transaction) web.TransactionResponse {
	return web.TransactionResponse{
		ID:         transaction.Transaction_detail_id,
		TotalPrice: transaction.Total_Price,
		Products: struct {
			Name  string "json:\"name\""
			Price int    "json:\"price\""
			Qty   int    "json:\"qty\""
		}{
			Name:  "contoh",
			Price: transaction.Total_Price,
		},
	}
}

func ToTransactionResponses(transaction []domain.TransactionDetail) []web.TransactionResponse {
	var transactionResponses []web.TransactionResponse
	for _, transaction := range transaction {
		transactionResponses = append(transactionResponses, web.TransactionResponse{
			ID:         transaction.Transaction_id,
			TotalPrice: transaction.Total_Price,
			Products: struct {
				Name  string "json:\"name\""
				Price int    "json:\"price\""
				Qty   int    "json:\"qty\""
			}{
				Name:  transaction.Product_Name,
				Price: transaction.Product_Price,
				Qty:   transaction.Product_Quantyty,
			},
		})
	}
	return transactionResponses

}
