package domain

type Product struct {
	ID          int
	Name        string
	Price       int
	Stock       int
	Description string
	Img         string
	CategoryId  int
}
