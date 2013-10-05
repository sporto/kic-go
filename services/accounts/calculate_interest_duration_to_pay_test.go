package accounts

import (
	"github.com/sporto/kic/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNoDate(t *testing.T) {
	account := models.Account{}
	now := time.Now()
	dur := CalculateInterestDurationToPay(account, now)
	assert.Equal(t, dur, 0, "they should be equal")
}

func TestDate(t *testing.T) {
	now := time.Now()
	ti := now.AddDate(0, 0, -4)
	account := models.Account{LastInterestPaid: ti}
	dur := CalculateInterestDurationToPay(account, now)
	expectedDur := now.Sub(ti)
	assert.Equal(t, dur, expectedDur, "they should be equal")
}
