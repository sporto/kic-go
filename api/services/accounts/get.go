package accounts

import (
	r "github.com/dancannon/gorethink"
	"github.com/sporto/kic/api/models"
	"errors"
)

type GetServ struct {
}

func (serv *GetServ) Run(dbSession *r.Session, id string) (account models.Account, err error) {

	row, err := r.Table("accounts").Get(id).RunRow(dbSession)
	if err != nil {
		return
	}

	if row.IsNil() {
		err = errors.New("Not found")
		return
	}

	err = row.Scan(&account)

	return
}
