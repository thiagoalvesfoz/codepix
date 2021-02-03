package repository

import (
	"fmt"
	"github.com/thiagoalvesfoz/codepix/domain/model"
	"gorm.io/gorm"
)

//type TransactionRepositoryInterface interface {
//	Register(transaction *Transaction) error
//	Save(transaction *Transaction) error
//	Find(id string) (*Transaction, error)
//}

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (repository *TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err := repository.Db.Create(transaction).Error
	return err
}

func (repository *TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := repository.Db.Save(transaction).Error
	return err
}

func (repository *TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	repository.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction found")
	}

	return &transaction, nil
}

