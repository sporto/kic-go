package models

import (
	"time"
)

type Account struct {
	Name             string
	LastInterestPaid time.Time
	Id               string `json:"id,omitempty"` // (will appear in json as "id", and not be sent if empty)
	CurrentBalance   float64
}
