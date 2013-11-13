package api

import (
	r "github.com/dancannon/gorethink"
	"github.com/sporto/kic/api/controllers"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
)

func MapRoutes(sessionArray []*r.Session) {
	
	accounts := &controllers.Accounts{DbSession: sessionArray[0]}
	goweb.MapController(accounts)
	
	accountTransactions := &controllers.AccountTransactions{DbSession: sessionArray[0]}
	goweb.MapController(accountTransactions)
	
	goweb.MapStatic("/public", "src/public")
	goweb.MapStaticFile("/", "src/index.html")
	goweb.MapStaticFile("/favicon.ico", "src/favicon.ico")

	// Catch-all handler for everything that we don't understand
	goweb.Map(func(c context.Context) error {
		// just return a 404 message
		return goweb.API.Respond(c, 404, nil, []string{"File not found"})
	})
}