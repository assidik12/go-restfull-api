package web

// Struktur data untuk request
type TransactionRequest struct {
	TotalPrice int `json:"total_price"`
	Products   []struct {
		ID  int `json:"id"`
		Qty int `json:"qty"`
	} `json:"products"`
}

// Struktur data untuk response
type TransactionResponse struct {
	ID         string `json:"id"`
	TotalPrice int    `json:"total_price"`
	Products   struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
		Qty   int    `json:"qty"`
	} `json:"products"`
}
