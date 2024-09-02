package repository

import (
	"context"
	"database/sql"

	"github.com/assidik12/go-restfull-api/helper"
	"github.com/assidik12/go-restfull-api/model/domain"
)

type TransactionRepositoryImpl struct{}

func NewTransactionRepository() *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{}
}

func (c *TransactionRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, USerId int) (results []domain.TransactionDetail, err error) {
	SQL := "SELECT t.id AS Transaction_id, t.total_price AS Total_Price, td.price AS Product_Price, td.quantity AS Product_Quantyty, p.name AS Product_Name, a.username AS UserName FROM transaction t INNER JOIN transaction_detail td ON t.transaction_detail = td.transaction_id INNER JOIN product p ON td.product_id = p.id INNER JOIN account a ON t.user_id = a.id WHERE a.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, USerId)
	helper.PanicError(err)

	defer rows.Close()
	for rows.Next() {
		var result domain.TransactionDetail
		err := rows.Scan(&result.Transaction_id, &result.Total_Price, &result.Product_Price, &result.Product_Quantyty, &result.Product_Name, &result.UserName)
		helper.PanicError(err)
		results = append(results, result)
	}

	if results == nil {
		return nil, sql.ErrNoRows
	}

	return results, nil
}

func (c *TransactionRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction {

	SQL := "INSERT INTO transaction (transaction_detail, user_id, total_price) VALUES (?,?,?)"
	_, err := tx.ExecContext(ctx, SQL, transaction.Transaction_detail_id, transaction.User_id, transaction.Total_Price)
	helper.PanicError(err)

	detailProduct := domain.Transaction{
		Transaction_detail_id: transaction.Transaction_detail_id,
		Product_id:            transaction.Product_id,
		Quantyty:              transaction.Quantyty,
	}

	SaveDetailTransaction(tx, detailProduct)

	return detailProduct
}

func (c *TransactionRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, transaction_id int) {
	SQL := "DELETE FROM transaction WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, transaction_id)
	helper.PanicError(err)
}

// internal func ✍️

func SaveDetailTransaction(tx *sql.Tx, transaction domain.Transaction) {

	for i, v := range transaction.Product_id {
		queryTransactionDetail := "INSERT INTO transaction_detail (transaction_id, product_id, quantity) VALUES (?,?,?)"
		_, err := tx.Exec(queryTransactionDetail, transaction.Transaction_detail_id, v, transaction.Quantyty[i])
		helper.PanicError(err)
	}

	updateStokck := domain.Transaction{
		Product_id: transaction.Product_id,
		Quantyty:   transaction.Quantyty,
	}
	updateProductChekout(tx, updateStokck)
}

func updateProductChekout(tx *sql.Tx, transaction domain.Transaction) {
	SQL := "UPDATE product SET stock = stock - ? WHERE id = ?"
	for i, v := range transaction.Product_id {
		_, err := tx.Exec(SQL, transaction.Quantyty[i], v)
		helper.PanicError(err)
	}
}
