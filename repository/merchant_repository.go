package repository

import (
	"fmt"
	"mnc/model"

	"gorm.io/gorm"
)

type MerchantRepository interface {
	BaseRepository[model.Merchant]
	GetByName(name string) (model.Merchant, error)
}

type merchantRepository struct {
	db *gorm.DB
}

func (m *merchantRepository) Create(payload model.Merchant) error {
	result := m.db.Create(&payload)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Merchant created successfully")
	return nil
}

func (m *merchantRepository) List() ([]model.Merchant, error) {
	var merchants []model.Merchant
	result := m.db.Find(&merchants)
	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Println("Merchant retrieve all successfully")
	return merchants, nil
}

func (m *merchantRepository) Get(id string) (model.Merchant, error) {
	var merchants model.Merchant
	err := m.db.Where("id = $1", id).First(&merchants).Error
	return merchants, err
}

func (m *merchantRepository) Update(payload model.Merchant) error {
	result := m.db.Save(&payload)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Merchant successfully Updated")
	return nil
}

func (m *merchantRepository) Delete(id string) error {
	result := m.db.Delete(&model.Merchant{}, id)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Merchant successfully Deleted")
	return nil
}

func (m *merchantRepository) GetByName(name string) (model.Merchant, error) {
	var merchant model.Merchant
	result := m.db.Where("name ILIKE ?", "%"+name+"%").First(&merchant)
	if result.Error != nil {
		return model.Merchant{}, result.Error
	}
	return merchant, nil
}

func NewMerchantRepository(db *gorm.DB) MerchantRepository {
	return &merchantRepository{db: db}
}
