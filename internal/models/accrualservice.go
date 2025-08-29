package models

//go:generate easyjson accrualservice.go

//easyjson:json
type AccrualData struct {
	Order   string  `json:"order"`
	Status  string  `json:"status"`
	Accrual float64 `json:"accrual,omitempty"`
}

type OrdersForAccrualCalculation struct {
	CurrentStatus string `db:"status"`
	OrderID       string `db:"order_id"`
	UserID        uint   `db:"user_id"`
}
