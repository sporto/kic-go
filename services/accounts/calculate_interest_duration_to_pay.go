package accounts

import (
	"github.com/sporto/kic/models"
	"time"
)

func CalculateInterestDurationToPay(account models.Account) (days time.Duration) {
	now := time.Now()
	days = account.LastInterestPaid.Sub(now)
	return
}
