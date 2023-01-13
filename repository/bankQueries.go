package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wigwamwam/CRUD_app/models"
	customErrors "github.com/wigwamwam/CRUD_app/repository/errors"
)

type DB struct {
	pool *pgxpool.Pool
}

func NewDb(pool *pgxpool.Pool) *DB {
	return &DB{pool: pool}
}

func (db *DB) SelectAllBanks() ([]models.Bank, error) {
	query := "SELECT id, name, iban, created_at, updated_at, deleted_at FROM banks"
	rows, err := db.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var banks []models.Bank
	for rows.Next() {
		bankList := models.Bank{}
		err = rows.Scan(&bankList.ID, &bankList.Name, &bankList.IBAN, &bankList.CreatedAt, &bankList.UpdatedAt, &bankList.DeletedAt)
		if err != nil {
			return nil, customErrors.NewNotFoundError()
		}
		banks = append(banks, bankList)
	}

	return banks, nil
}

func (db *DB) InsertBank(bank models.Bank) (models.Bank, error) {
	query := "INSERT INTO banks (name, iban, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5) returning id"
	err := db.pool.QueryRow(context.Background(), query, bank.Name, bank.IBAN, bank.CreatedAt, bank.UpdatedAt, bank.DeletedAt).Scan(&bank.ID)
	if err != nil {
		return models.Bank{}, customErrors.NewCreatingBankError()
	}
	return bank, nil
}

func (db *DB) UpdateBank(id int, bank models.Bank) (models.Bank, error) {
	query := "UPDATE banks SET name = $1, iban = $2, updated_at = $3 WHERE id = $4 RETURNING id, name, iban, created_at, updated_at, deleted_at"
	err := db.pool.QueryRow(context.Background(), query, bank.Name, bank.IBAN, time.Now(), id).Scan(&bank.ID, &bank.Name, &bank.IBAN, &bank.CreatedAt, &bank.UpdatedAt, &bank.DeletedAt)
	if err != nil {
		return models.Bank{}, customErrors.NewIdNotFoundError(id)
	}
	return bank, nil
}

func (db *DB) SelectBankByID(id int) (models.Bank, error) {
	query := "SELECT id, name, iban, created_at, updated_at, deleted_at FROM banks WHERE id = $1"
	row := db.pool.QueryRow(context.Background(), query, id)
	bank := models.Bank{}
	err := row.Scan(&bank.ID, &bank.Name, &bank.IBAN, &bank.CreatedAt, &bank.UpdatedAt, &bank.DeletedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return models.Bank{}, customErrors.NewIdNotFoundError(id)
		} else {
			return models.Bank{}, customErrors.NewScanningIdError(id)
		}
	}

	return bank, nil
}

// error when inputing a

func (db *DB) DeleteBankByID(id int) error {
	query := "DELETE FROM banks WHERE id = $1"
	row, err := db.pool.Exec(context.Background(), query, id)
	if err != nil {
		return customErrors.NewDeletingBankError()
	}

	count := row.RowsAffected()
	if count == 0 {
		return customErrors.NewIdNotFoundError(id)
	}
	return nil

}
