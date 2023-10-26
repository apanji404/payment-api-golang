package model

type Bank struct {
	ID        string `json:"id"`
	Bank_Name string `json:"bank_name"`
	Email     string `json:"email"`
	Telp      string `json:"telp"`
	Address   string `json:"address"`
}

// nama tabel yang digunakan sesuai dengan yang ada di skema basis data
func (Bank) TableName() string {
	return "bank"
}
