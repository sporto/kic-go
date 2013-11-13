package accounts

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	"github.com/sporto/kic/api/models"
)

type UpdateAllBalancesServ struct {
	updateBalanceServ UpdateBalanceServ
}

func (serv *UpdateAllBalancesServ) Run(dbSession *r.Session) (err error) {

	// Fetch all the items from the database
	rows, err := r.Table("accounts").OrderBy(r.Asc("CreatedAt")).Run(dbSession)

	for rows.Next() {
		var account models.Account
		err := rows.Scan(&account)
		if err != nil {
			fmt.Println("err", err)
			break
		}
		err = serv.updateBalanceServ.Run(dbSession, &account)
		if err != nil {
			fmt.Println("err", err)
			break
		}
	}
	// if err = rows.Err(); err != nil {
	// 	fmt.Println("err:", err)
	// }
	return
}
