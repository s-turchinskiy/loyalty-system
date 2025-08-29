package models

//go:generate easyjson accrualservice.go

type AccrualStatus int

const (
	REGISTERED AccrualStatus = iota
	INVALID
	PROCESSING
	PROCESSED
)

func (a AccrualStatus) AsString() string {

	switch a {
	case REGISTERED:
		return "REGISTERED"
	case INVALID:
		return "INVALID"
	case PROCESSING:
		return "PROCESSING"
	case PROCESSED:
		return "PROCESSED"
	default:
		return ""

	}
}

//easyjson:json
type AccrualData struct {
	Order   string        `json:"order"`
	Status  AccrualStatus `json:"status"`
	Accrual float64       `json:"accrual,omitempty"`
}

type OrdersForAccrualCalculation struct {
	CurrentStatus string `db:"status"`
	OrderID       string `db:"order_id"`
	UserID        uint   `db:"user_id"`
}
