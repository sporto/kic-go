package controllers

import (
	r "github.com/christopherhesse/rethinkgo"
	"github.com/sporto/kic/api/models"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"log"
	"net/http"
	"time"
)

type AccountTransactions struct {
	DbSession *r.Session
}

func (c *AccountTransactions) Path() string {
	return "api/accounts/{account_id}/transactions"
}

func (c *AccountTransactions) ReadMany(ctx context.Context) error {

	// ctx.PathValue("account_id")

	var records []models.Transaction

	err := r.Table("transactions").Run(c.DbSession).All(&records)

	if err != nil {
		log.Fatal(err)
		return goweb.Respond.WithStatus(ctx, http.StatusNotFound)
	}
	return goweb.API.RespondWithData(ctx, records)
}

func (c *AccountTransactions) Create(ctx context.Context) error {

	accountId := ctx.PathValue("account_id")

	data, err := ctx.RequestData()

	if err != nil {
		return goweb.API.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
	}

	dataMap := data.(map[string]interface{})

	t := time.Now()
	transaction := models.Transaction{
		AccountId:  accountId, 
		CreatedAt:  t,
		Amount:     dataMap["amount"].(float64),
		Kind:       dataMap["kind"].(string),
	}

	var record r.WriteResponse
	err = r.Table("transactions").Insert(transaction).Run(c.DbSession).One(&record)

	if err != nil {
		log.Fatal(err)
		return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)
	}

	return goweb.API.RespondWithData(ctx, record)
}