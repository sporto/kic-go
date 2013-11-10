package accounts

import (
	r "github.com/dancannon/gorethink"
	"github.com/sporto/kic/api/models"
	"errors"
	"time"
)

type UpdateServ struct {
}

func (serv *UpdateServ) Run(dbSession *r.Session, account models.Account) (err error) {
	//map[string]interface{}

	account.UpdatedAt = time.Now()

	response, err := r.Table("accounts").Get(id).Update(account).RunRow(dbSession)

	if err != nil {
		return
	}

	return
}
