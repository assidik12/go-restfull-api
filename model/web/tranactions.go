package web

// Struktur data untuk request
type TransactionRequest struct {
	TotalPrice int `json:"total_price"`
	User_id    int `json:"user_id"`
	IDTrx      int `json:"id_trx"`
	Products   []struct {
		ID  int `json:"id"`
		Qty int `json:"qty"`
	} `json:"products"`
}

// Struktur data untuk response

type TransactionResponse struct {
	ID         int `json:"id"`
	TotalPrice int `json:"total_price"`
	Products   struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
		Qty   int    `json:"qty"`
	} `json:"products"`
}
