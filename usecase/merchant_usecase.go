package usecase

import (
	"fmt"
	"mnc/model"
	"mnc/repository"
)

type MerchantUseCase interface {
	RegisterNewMerchant(payload model.Merchant) error
	FindAllMerchant() ([]model.Merchant, error)
	FindByMerchantId(id string) (model.Merchant, error)
	UpdateMerchant(payload model.Merchant) error
	DeleteMerchant(id string) error
	GetByMerchantName(name string) (model.Merchant, error)
}

type merchantUseCase struct {
	repo repository.MerchantRepository
}

func (m *merchantUseCase) RegisterNewMerchant(payload model.Merchant) error {
	if payload.Merchant_Name == "" {
		return fmt.Errorf("name are required fields")
	}
	isExistMerchant, _ := m.repo.GetByName(payload.Merchant_Name)
	if isExistMerchant.Merchant_Name == payload.Merchant_Name {
		return fmt.Errorf("merchant with name %s exists", payload.Merchant_Name)
	}
	err := m.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create new merchant: %v", err)
	}
	return nil
}

func (m *merchantUseCase) FindAllMerchant() ([]model.Merchant, error) {
	return m.repo.List()
}

func (m *merchantUseCase) FindByMerchantId(id string) (model.Merchant, error) {
	return m.repo.Get(id)
}

func (m *merchantUseCase) UpdateMerchant(payload model.Merchant) error {
	if payload.Merchant_Name == "" {
		return fmt.Errorf("merchant name is required")
	}

	isExistMerchant, _ := m.repo.GetByName(payload.Merchant_Name)
	if isExistMerchant.Merchant_Name == payload.Merchant_Name && isExistMerchant.ID != payload.ID {
		return fmt.Errorf("merchant with name %s exists", payload.Merchant_Name)
	}

	err := m.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update merchant: %v", err)
	}

	return nil
}

func (m *merchantUseCase) DeleteMerchant(id string) error {
	return m.repo.Delete(id)
}

func (m *merchantUseCase) GetByMerchantName(name string) (model.Merchant, error) {
	return m.repo.GetByName(name)
}

func NewMerchantUseCase(repo repository.MerchantRepository) MerchantUseCase {
	return &merchantUseCase{repo: repo}
}
