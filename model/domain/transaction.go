package domain

type Transaction struct {
	Transaction_detail_id int
	User_id               int
	Total_Price           int
	Product_id            []int
	Quantyty              []int
}

type TransactionDetail struct {
	UserName         string
	Transaction_id   int
	Total_Price      int
	Product_Quantyty int
	Product_Price    int
	Product_Name     string
}