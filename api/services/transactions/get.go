package transactions

import (
	"errors"
	r "github.com/dancannon/gorethink"
	"github.com/sporto/kic/api/models"
)

type GetServ struct {
}

func (serv *GetServ) Run(dbSession *r.Session, id string) (transaction models.Transaction, err error) {

	row, err := r.Table("transactions").Get(id).RunRow(dbSession)
	if err != nil {
		return
	}

	if row.IsNil() {
		err = errors.New("Not found")
		return
	}

	err = row.Scan(&transaction)

	return
}
