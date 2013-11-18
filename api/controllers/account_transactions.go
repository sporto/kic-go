package controllers

import (
	r "github.com/dancannon/gorethink"
	"github.com/sporto/kic/api/models"
	"github.com/sporto/kic/api/services/transactions"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"log"
	"net/http"
)

type AccountTransactions struct {
	DbSession *r.Session
}

func (c *AccountTransactions) Path() string {
	return "api/accounts/{account_id}/transactions"
}

func (c *AccountTransactions) ReadMany(ctx context.Context) (err error) {

	var transactions []models.Transaction

	rows, err := r.Table("transactions").OrderBy(r.Desc("createdAt")).Run(c.DbSession)

	for rows.Next() {
		var transaction models.Transaction

		err = rows.Scan(&transaction)
		if err != nil {
			return
		}

		transactions = append(transactions, transaction)
	}

	if err != nil {
		log.Fatal(err)
		return goweb.Respond.WithStatus(ctx, http.StatusNotFound)
	}
	return goweb.API.RespondWithData(ctx, transactions)
}

func (c *AccountTransactions) Create(ctx context.Context) (err error) {
	accountId := ctx.PathValue("account_id")
	data, err := ctx.RequestData()

	if err != nil {
		return goweb.API.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
	}

	dataMap := data.(map[string]interface{})

	transaction := models.Transaction{
		AccountId: accountId,
		Debit:     dataMap["debit"].(float64),
		Credit:    dataMap["credit"].(float64),
		Kind:      dataMap["kind"].(string),
	}

	createServ := &transactions.CreateServ{}
	transactionOut, err := createServ.Run(c.DbSession, transaction)
	if err != nil {
		log.Print(err)
		return goweb.API.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
	}

	return goweb.API.RespondWithData(ctx, transactionOut)
}
