package models

type Account struct {
	LastInterestPaid string
	Id               string `json:"id,omitempty"` // (will appear in json as "id", and not be sent if empty)
}
