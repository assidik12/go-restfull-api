package web

type ProductResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Img         string `json:"img"`
	Description string `json:"description"`
	CategoryId  int    `json:"categoryId"`
}
