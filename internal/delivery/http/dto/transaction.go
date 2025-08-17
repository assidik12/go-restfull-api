package dto

type TransactionRequest struct {
	TotalPrice int `json:"totalPrice" binding:"required"`
	Products   []struct {
		ID  int `json:"id" binding:"required"`
		Qty int `json:"qty" binding:"required"`
	} `json:"products" binding:"required"`
}

type TransactionResponse struct {
	ID         string `json:"id"`
	TotalPrice int    `json:"totalPrice"`
	Products   []struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
		Qty   int    `json:"qty"`
	} `json:"products"`
}
