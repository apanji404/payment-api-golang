package usecase

import (
	"fmt"
	"mnc/model"
	"mnc/repository"

	"golang.org/x/crypto/bcrypt"
)

type CustomerUseCase interface {
	RegisterNewCustomer(payload model.Customer) error
	FindAllCustomer() ([]model.Customer, error)
	FindByCustomerId(id string) (model.Customer, error)
	UpdateCustomer(payload model.Customer) error
	DeleteCustomer(id string) error
	FindByCustomerUsername(name string) (model.Customer, error)
	FindByCustomerUsernamePassword(username, password string) (model.Customer, error)
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

func (b *customerUseCase) RegisterNewCustomer(payload model.Customer) error {
	if payload.Username == "" {
		return fmt.Errorf("username are required fields")
	}

	isExistCustomer, _ := b.repo.GetByUsername(payload.Username)
	if isExistCustomer.Username == payload.Username {
		return fmt.Errorf("customer with name %s exists", payload.Username)
	}

	bytes, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	payload.Password = string(bytes)
	err := b.repo.Create(payload)

	if err != nil {
		return fmt.Errorf("failed to create new customer: %v", err)
	}
	return nil
}

func (b *customerUseCase) FindAllCustomer() ([]model.Customer, error) {
	return b.repo.List()
}

func (b *customerUseCase) FindByCustomerId(id string) (model.Customer, error) {
	return b.repo.Get(id)
}

func (b *customerUseCase) UpdateCustomer(payload model.Customer) error {
	if payload.ID == "" {
		return fmt.Errorf("customer id is required")
	}

	if payload.Username == "" {
		return fmt.Errorf("customer name is required")
	}

	isExistCustomer, _ := b.repo.GetByUsername(payload.Username)
	if isExistCustomer.Username == payload.Username && isExistCustomer.ID != payload.ID {
		return fmt.Errorf("customer with name %s exists", payload.Username)
	}

	bytes, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	payload.Password = string(bytes)
	err := b.repo.Update(payload)

	if err != nil {
		return fmt.Errorf("failed to update customer: %v", err)
	}

	return nil
}

func (c *customerUseCase) DeleteCustomer(id string) error {
	return c.repo.Delete(id)
}

func (c *customerUseCase) FindByCustomerUsername(name string) (model.Customer, error) {
	return c.repo.GetByUsername(name)
}

func (c *customerUseCase) FindByCustomerUsernamePassword(username string, password string) (model.Customer, error) {
	return c.repo.GetByUsernamePassword(username, password)
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{repo: repo}
}
