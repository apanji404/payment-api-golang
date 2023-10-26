package manager

import "mnc/repository"

type RepoManager interface {
	BankRepo() repository.BankRepository
	CustomerRepo() repository.CustomerRepository
	MerchantRepo() repository.MerchantRepository
	TransactionRepo() repository.TransactionRepository
}

type repoManager struct {
	infra InfraManager
}

// BankRepo implements RepoManager.
func (r *repoManager) BankRepo() repository.BankRepository {
	return repository.NewBankRepository(r.infra.Conn())
}

// CustomerRepo implements RepoManager.
func (r *repoManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infra.Conn())
}

// MerchantRepo implements RepoManager.
func (r *repoManager) MerchantRepo() repository.MerchantRepository {
	return repository.NewMerchantRepository(r.infra.Conn())
}

// TransactionRepo implements RepoManager.
func (r *repoManager) TransactionRepo() repository.TransactionRepository {
	return repository.NewTransactionRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
