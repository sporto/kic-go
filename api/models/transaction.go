package models

import (
	"time"
)

type Transaction struct {
	CreatedAt        time.Time
	Kind             string
	Id               string `json:"id,omitempty"` // (will appear in json as "id", and not be sent if empty)
	AccountId        string
	Amount           float64
}
