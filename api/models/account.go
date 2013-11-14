package models

import (
	"time"
)

type Account struct {
	Name             string     `json:"name",gorethink:"name"`
	CreatedAt        time.Time  `json:"createdAt",gorethink:"createdAt"`
	UpdatedAt        time.Time  `json:"createdAt",gorethink:"createdAt"`
	LastInterestPaid time.Time  `json:"lastInterestPaid",gorethink:"lastInterestPaid"`
	Id               string     `json:"id",gorethink:"id,omitempty"` // (will appear in json as "id", and not be sent if empty)
	CurrentBalance   float64    `json:"currentBalance",gorethink:"currentBalance"`
}
