package models

import (
	"time"
)

type Transaction struct {
	Id               string       `json:"id"gorethink:"id,omitempty"` // (will appear in json as "id", and not be sent if empty)
	CreatedAt        time.Time    `json:"createdAt"gorethink:"createdAt"`
	UpdatedAt        time.Time    `json:"updatedAt"gorethink:"updatedAt"`
	Kind             string       `json:"kind"gorethink:"kind"`
	AccountId        string       `json:"accountId"gorethink:"accountId"`
	Debit            float64      `json:"debit"gorethink:"debit"`
	Credit           float64      `json:"credit"gorethink:"credit"`
}
