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
	err := p.db.Preload("Customer").Preload("Merchant").Preload("Bank").
		Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	fmt.Println("Transactions retrieved successfully")
	return transactions, nil
}

func (p *transactionRepository) Get(id string) (model.Transaction, error) {
	var transactions model.Transaction
	err := p.db.Preload("Customer").Preload("Merchant").Preload("Bank").
		Where("id = ?", id).
		First(&transactions).Error
	if err != nil {
		return model.Transaction{}, err
	}

	return transactions, nil
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
