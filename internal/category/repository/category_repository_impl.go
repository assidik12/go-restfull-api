package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/assidik12/go-restfull-api/helper"
	"github.com/assidik12/go-restfull-api/model/domain"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}
func (c *CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO category(name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicError(err)
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	category.ID = int(id)
	return category

}
func (c *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE category SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.ID)
	helper.PanicError(err)
	return category

}
func (c *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.ID)
	helper.PanicError(err)
}

func (c *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryID int) (domain.Category, error) {
	SQL := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryID)
	helper.PanicError(err)

	defer rows.Close()

	category := domain.Category{}

	if rows.Next() {
		err := rows.Scan(&category.ID, &category.Name)
		helper.PanicError(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}

}
func (c *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {

	SQL := "SELECT id, name FROM category"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicError(err)
	defer rows.Close()

	var catagories []domain.Category

	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.ID, &category.Name)
		helper.PanicError(err)
		catagories = append(catagories, category)
	}

	return catagories

}
