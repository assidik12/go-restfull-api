package web

type CreateProductRequest struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	Img          string `json:"img"`
	CategoryName string `json:"categoryName"`
}

type UpdateProductRequest struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Img         string `json:"img"`
}
