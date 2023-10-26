package repository

import (
	"fmt"
	"mnc/model"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	BaseRepository[model.Transaction]
}

type transactionRepository struct {
	db *gorm.DB
}

func (p *transactionRepository) Create(payload model.Transaction) error {
	result := p.db.Create(&payload)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Transaction created successfully")
	return nil
}

func (p *transactionRepository) List() ([]model.Transaction, error) {
	var transactions []model.Transaction
	result := p.db.Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Println("Transaction retrieve all successfully")
	return transactions, nil
}

func (p *transactionRepository) Get(id string) (model.Transaction, error) {
	var transactions model.Transaction
	err := p.db.Where("id = $1", id).First(&transactions).Error
	return transactions, err
}

func (p *transactionRepository) Update(payload model.Transaction) error {
	result := p.db.Save(&payload)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Transaction successfully Updated")
	return nil
}

func (p *transactionRepository) Delete(id string) error {
	result := p.db.Delete(&model.Transaction{}, id)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Transaction successfully Deleted")
	return nil
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}
