package manager

import "mnc/usecase"

type UsecaseManager interface {
	BankUsecase() usecase.BankUseCase
	CustomerUsecase() usecase.CustomerUseCase
	MerchantUsecase() usecase.MerchantUseCase
	TransactionUsecase() usecase.TransactionUseCase
	AuthUsecase() usecase.AuthUseCase
}

type usecaseManager struct {
	repo RepoManager
}

// AuthUsecase implements UsecaseManager.
func (u *usecaseManager) AuthUsecase() usecase.AuthUseCase {
	return usecase.NewAuthUseCase(u.CustomerUsecase())
}

// BankUsecase implements UsecaseManager.
func (u *usecaseManager) BankUsecase() usecase.BankUseCase {
	return usecase.NewBankUseCase(u.repo.BankRepo())
}

// CustomerUsecase implements UsecaseManager.
func (u *usecaseManager) CustomerUsecase() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repo.CustomerRepo())
}

// MerchantUsecase implements UsecaseManager.
func (u *usecaseManager) MerchantUsecase() usecase.MerchantUseCase {
	return usecase.NewMerchantUseCase(u.repo.MerchantRepo())
}

// TransactionUsecase implements UsecaseManager.
func (u *usecaseManager) TransactionUsecase() usecase.TransactionUseCase {
	return usecase.NewTransactionUseCase(u.repo.TransactionRepo(), u.CustomerUsecase(), u.MerchantUsecase(), u.BankUsecase())
}

func NewUseCaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{repo: repoManager}
}
