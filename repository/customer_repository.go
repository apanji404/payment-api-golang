package repository

import (
	"fmt"
	"mnc/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	BaseRepository[model.Customer]
	GetByUsername(username string) (model.Customer, error)
	GetByUsernamePassword(username, password string) (model.Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}

func (c *customerRepository) Create(payload model.Customer) error {
	result := c.db.Create(&payload)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Customer created successfully")
	return nil
}

func (c *customerRepository) List() ([]model.Customer, error) {
	var customers []model.Customer
	result := c.db.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Println("Customer retrieve all successfully")
	return customers, nil
}

func (c *customerRepository) Get(id string) (model.Customer, error) {
	var customers model.Customer
	err := c.db.Where("id = ?", id).First(&customers).Error
	return customers, err
}

func (c *customerRepository) Update(payload model.Customer) error {
	result := c.db.Save(&payload)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Customer successfully Updated")
	return nil
}

func (c *customerRepository) Delete(id string) error {
	result := c.db.Delete(&model.Customer{}, id)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Customer successfully Deleted")
	return nil
}

func (c *customerRepository) GetByUsername(username string) (model.Customer, error) {
	var customer model.Customer
	err := c.db.Where("username = ?", username).First(&customer).Error
	return customer, err
}

func (c *customerRepository) GetByUsernamePassword(username, password string) (model.Customer, error) {
	customer, err := c.GetByUsername(username)
	if err != nil {
		return model.Customer{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password))
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}
