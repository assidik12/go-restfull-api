package domain

type Transaction struct {
	ID                    int
	Transaction_detail_id string
	User_id               int
	Total_Price           int
	Products              []TransactionDetail
}

type TransactionDetail struct {
	Transaction_detail_id string
	Quantyty              int
	Product_id            int
}
