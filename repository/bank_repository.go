package repository

import (
	"fmt"
	"mnc/model"

	"gorm.io/gorm"
)

type BankRepository interface {
	BaseRepository[model.Bank]
	GetByName(name string) (model.Bank, error)
}

type bankRepository struct {
	db *gorm.DB
}

func (b *bankRepository) Create(payload model.Bank) error {
	result := b.db.Create(&payload)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Bank created successfully")
	return nil
}

func (b *bankRepository) List() ([]model.Bank, error) {
	var banks []model.Bank
	result := b.db.Find(&banks)
	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Println("Bank retrieve all successfully")
	return banks, nil
}

func (b *bankRepository) Get(id string) (model.Bank, error) {
	var banks model.Bank
	err := b.db.Where("id = $1", id).First(&banks).Error
	return banks, err
}

func (b *bankRepository) Update(payload model.Bank) error {
	result := b.db.Save(&payload)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Bank successfully Updated")
	return nil
}

func (b *bankRepository) Delete(id string) error {
	result := b.db.Delete(&model.Bank{}, id)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Bank successfully Deleted")
	return nil
}

func (p *bankRepository) GetByName(name string) (model.Bank, error) {
	var bank model.Bank
	result := p.db.Where("name ILIKE ?", "%"+name+"%").First(&bank)
	if result.Error != nil {
		return model.Bank{}, result.Error
	}
	return bank, nil
}

func NewBankRepository(db *gorm.DB) BankRepository {
	return &bankRepository{db: db}
}
