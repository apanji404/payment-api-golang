package usecase

import (
	"fmt"
	"mnc/model"
	"mnc/repository"
	"time"
)

type TransactionUseCase interface {
	CreateTransaction(payload model.Transaction) error
	FindByTransactionId(id string) (model.Transaction, error)
	FindAllTransaction() ([]model.Transaction, error)
}

type transactionUseCase struct {
	transactionRepo repository.TransactionRepository
	customerUC      CustomerUseCase
	merchantUC      MerchantUseCase
	bankUC          BankUseCase
}

func (t *transactionUseCase) CreateTransaction(trx model.Transaction) error {

	customer, err := t.customerUC.FindByCustomerId(trx.Customer_ID)
	if err != nil {
		return err
	}

	merchant, err := t.merchantUC.FindByMerchantId(trx.Merchant_ID)
	if err != nil {
		return err
	}

	bank, err := t.bankUC.FindByBankId(trx.Bank_ID)
	if err != nil {
		return err
	}

	trx.Customer_ID = customer.ID
	trx.Merchant_ID = merchant.ID
	trx.Bank_ID = bank.ID
	trx.Payment_Time = time.Now()

	err = t.transactionRepo.Create(trx)
	if err != nil {
		return fmt.Errorf("failed to register new transaction %v", err)
	}

	return nil

}

func (t *transactionUseCase) FindAllTransaction() ([]model.Transaction, error) {
	return t.transactionRepo.List()
}

func (t *transactionUseCase) FindByTransactionId(id string) (model.Transaction, error) {
	return t.transactionRepo.Get(id)
}

func NewTransactionUseCase(transactionRepo repository.TransactionRepository, customerUC CustomerUseCase, merchantUC MerchantUseCase, bankUC BankUseCase) TransactionUseCase {
	return &transactionUseCase{
		transactionRepo: transactionRepo,
		customerUC:      customerUC,
		merchantUC:      merchantUC,
		bankUC:          bankUC,
	}
}
