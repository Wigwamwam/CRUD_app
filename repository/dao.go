package repository

import "github.com/wigwamwam/CRUD_app/models"

type DAO interface {
	SelectAllBanks() ([]models.Bank, error)
	InsertBank(bank models.Bank) (models.Bank, error)
	UpdateBank(id int, bank models.Bank) (models.Bank, error)
	SelectBankByID(id int) (models.Bank, error)
	DeleteBankByID(id int) error
}
