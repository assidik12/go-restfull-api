package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/assidik12/go-restfull-api/helper"
	"github.com/assidik12/go-restfull-api/model/domain"
)

type ProductRepositoryImpl struct{}

func NewProductRepository() *ProductRepositoryImpl {
	return &ProductRepositoryImpl{}
}

func (c *ProductRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {

	SQL := "INSERT INTO product(name, price, stock,  gambar, description, category_id) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Stock, product.Img, product.Description, product.CategoryId)
	helper.PanicError(err)
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	product.ID = int(id)
	return product

}
func (c *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "UPDATE product SET name, price, stock, description, img = (?, ?, ?, ?, ?) WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Stock, product.Description, product.Img, product.ID)
	helper.PanicError(err)
	return product

}
func (c *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	SQL := "DELETE FROM product WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.ID)
	helper.PanicError(err)
}

func (c *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	SQL := "SELECT * FROM product WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicError(err)

	defer rows.Close()

	product := domain.Product{}

	if rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.Description, &product.Img)
		helper.PanicError(err)
		return product, nil
	} else {
		return product, errors.New("product not found")
	}

}
func (c *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {

	SQL := "SELECT * FROM product lIMIT 50"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicError(err)
	defer rows.Close()

	var products []domain.Product

	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.Description, &product.Img)
		helper.PanicError(err)
		products = append(products, product)
	}

	return products

}
