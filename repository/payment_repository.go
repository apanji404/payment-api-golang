package repository

import (
	"fmt"
	"mnc/model"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	BaseRepository[model.Payment]
}

type paymentRepository struct {
	db *gorm.DB
}

func (p *paymentRepository) Create(payload model.Payment) error {
	result := p.db.Create(&payload)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Payment created successfully")
	return nil
}

func (p *paymentRepository) List() ([]model.Payment, error) {
	var payments []model.Payment
	result := p.db.Find(&payments)
	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Println("Payment retrieve all successfully")
	return payments, nil
}

func (p *paymentRepository) Get(id string) (model.Payment, error) {
	var payments model.Payment
	err := p.db.Where("id = $1", id).First(&payments).Error
	return payments, err
}

func (p *paymentRepository) Update(payload model.Payment) error {
	result := p.db.Save(&payload)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Payment successfully Updated")
	return nil
}

func (p *paymentRepository) Delete(id string) error {
	result := p.db.Delete(&model.Payment{}, id)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Payment successfully Deleted")
	return nil
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}
