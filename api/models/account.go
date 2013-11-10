package models

import (
	"time"
)

type Account struct {
	Name             string     `gorethink:"name"`
	CreatedAt        time.Time  `gorethink:"createdAt"`
	UpdatedAt        time.Time  `gorethink:"updatedAt"`
	LastInterestPaid time.Time  `gorethink:"lastInterestPaid"`
	Id               string     `gorethink:"id,omitempty"` // (will appear in json as "id", and not be sent if empty)
	CurrentBalance   float64    `gorethink:"currentBalance"`
}
