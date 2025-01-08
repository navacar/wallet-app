package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	DEPOSIT  = "DEPOSIT"
	WITHDRAW = "WITHDRAW"
)

type DB struct {
	db *sqlx.DB
}

func NewDB(db *sqlx.DB) *DB {
	return &DB{
		db: db,
	}
}

func (r *DB) Balance(id int) (float64, error) {
	var balance float64

	query := fmt.Sprintf("SELECT balance FROM %s WHERE walletID = $1", walletsTable)
	err := r.db.Get(&balance, query, id)

	return balance, err
}

func (r *DB) DepOrWithdraw(walletId int, operationType string, amount float64) (float64, error) {
	fmt.Println(walletId, operationType, amount)

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var balString string

	if operationType == DEPOSIT {
		balString = "balance + $1 WHERE walletID = $2"
	} else {
		balString = "balance - $1 WHERE walletID = $2 AND balance >= $1"
	}

	var balanceAfter float64
	query := fmt.Sprintf("UPDATE %s SET balance = %s RETURNING balance", walletsTable, balString)
	row := tx.QueryRow(query, amount, walletId)

	if err := row.Scan(&balanceAfter); err != nil {
		tx.Rollback()
		return 0, err
	}

	return balanceAfter, tx.Commit()
}
