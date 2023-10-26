package usecase

import (
	"fmt"
	"mnc/model"
	"mnc/repository"
)

type BankUseCase interface {
	RegisterNewBank(payload model.Bank) error
	FindAllBank() ([]model.Bank, error)
	FindByBankId(id string) (model.Bank, error)
	UpdateBank(payload model.Bank) error
	DeleteBank(id string) error
	GetByBankName(name string) (model.Bank, error)
}

type bankUseCase struct {
	repo repository.BankRepository
}

func (b *bankUseCase) RegisterNewBank(payload model.Bank) error {
	if payload.Bank_Name == "" {
		return fmt.Errorf("name are required fields")
	}
	isExistBank, _ := b.repo.GetByName(payload.Bank_Name)
	if isExistBank.Bank_Name == payload.Bank_Name {
		return fmt.Errorf("bank with name %s exists", payload.Bank_Name)
	}
	err := b.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create new bank: %v", err)
	}
	return nil
}

func (b *bankUseCase) FindAllBank() ([]model.Bank, error) {
	return b.repo.List()
}

func (b *bankUseCase) FindByBankId(id string) (model.Bank, error) {
	return b.repo.Get(id)
}

func (b *bankUseCase) UpdateBank(payload model.Bank) error {
	if payload.Bank_Name == "" {
		return fmt.Errorf("bank name is required")
	}

	isExistBank, _ := b.repo.GetByName(payload.Bank_Name)
	if isExistBank.Bank_Name == payload.Bank_Name && isExistBank.ID != payload.ID {
		return fmt.Errorf("bank with name %s exists", payload.Bank_Name)
	}

	err := b.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update bank: %v", err)
	}

	return nil
}

func (b *bankUseCase) DeleteBank(id string) error {
	return b.repo.Delete(id)
}

func (b *bankUseCase) GetByBankName(name string) (model.Bank, error) {
	return b.repo.GetByName(name)
}

func NewBankUseCase(repo repository.BankRepository) BankUseCase {
	return &bankUseCase{repo: repo}
}
