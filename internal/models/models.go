package models

//go:generate easyjson models.go

type Order struct {
	Number     string  `db:"number" json:"number"`
	Status     string  `db:"status" json:"status"`
	Accrual    float64 `db:"accrual" json:"accrual,omitempty"`
	UploadedAt string  `db:"uploaded_at" json:"uploaded_at"`
}

//easyjson:json
type Orders []Order

//easyjson:json
type Balance struct {
	Current   float64 `db:"current" json:"current"`
	Withdrawn float64 `db:"withdrawn" json:"withdrawn"`
}

type NewWithdraw struct {
	Order string
	Sum   float64
}

type Withdrawal struct {
	Order       string  `db:"order" json:"order"`
	Sum         float64 `db:"sum" json:"sum"`
	ProcessedAt string  `db:"processed_at" json:"processed_at"`
}

//easyjson:json
type Withdrawals []Withdrawal
