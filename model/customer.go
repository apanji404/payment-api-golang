package model

type Customer struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Telp     string `json:"telp"`
	Address  string `json:"address"`
}

// nama tabel yang digunakan sesuai dengan yang ada di skema basis data
func (Customer) TableName() string {
	return "customer"
}
