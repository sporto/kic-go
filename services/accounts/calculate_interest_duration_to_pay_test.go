package accounts

import (
	"github.com/sporto/kic/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNoDate(t *testing.T) {
	account := models.Account{}
	dur := CalculateInterestDurationToPay(account)
	assert.Equal(t, dur, 1, "they should be equal")
}

func TestDate(t *testing.T) {
	ti := time.Now().AddDate(0, 0, -4)
	account := models.Account{LastInterestPaid: ti}
	dur := CalculateInterestDurationToPay(account)
	assert.Equal(t, dur, 4, "they should be equal")
}
