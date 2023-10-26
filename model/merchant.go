package model

type Merchant struct {
	ID            string `json:"id"`
	Merchant_Name string `json:"merchant_name"`
	Email         string `json:"email"`
	Telp          string `json:"telp"`
	Address       string `json:"address"`
}

// nama tabel yang digunakan sesuai dengan yang ada di skema basis data
func (Merchant) TableName() string {
	return "merchant"
}
