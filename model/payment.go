package model

import "time"

type Payment struct {
	ID             string    `json:"id"`
	Customer_ID    string    `json:"customer_id"`
	Merchant_ID    string    `json:"merchant_id"`
	Bank_ID        string    `json:"bank_id"`
	Customer       Customer  `gorm:"foreignkey:Customer_ID"`
	Merchant       Merchant  `gorm:"foreignkey:Merchant_ID"`
	Bank           Bank      `gorm:"foreignkey:Bank_ID"`
	Payment_Amount float64   `json:"payment_amount"`
	Payment_Time   time.Time `json:"payment_time"`
}

// nama tabel yang digunakan sesuai dengan yang ada di skema basis data
func (Payment) TableName() string {
	return "payment"
}
