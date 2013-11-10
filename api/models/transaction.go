package models

import (
	"time"
)

type Transaction struct {
	Id               string       `gorethink:"id,omitempty"` // (will appear in json as "id", and not be sent if empty)
	CreatedAt        time.Time    `gorethink:"createdAt"`
	UpdatedAt        time.Time    `gorethink:"updatedAt"`
	Kind             string       `gorethink:"kind"`
	AccountId        string       `gorethink:"accountId"`
	Debit            float64      `gorethink:"debit"`
	Credit           float64      `gorethink:"credit"`
}
