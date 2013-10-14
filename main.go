package main

import (
	r "github.com/christopherhesse/rethinkgo"
	"github.com/sporto/kic/api/controllers"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	Address string = ":9000"
)

var sessionArray []*r.Session

func initDb() {
	// session, err := r.Connect(os.Getenv("WERCKER_RETHINKDB_URL"), "gettingstarted")
	session, err := r.Connect("localhost:28015", "kic")
	if err != nil {
		log.Fatal(err)
		log.Println("Most likely RethinkDB is not running")
		return
	}
	//defer session.Close()

	err = r.DbCreate("kic").Run(session).Exec()
	if err != nil {
		log.Println(err)
	}

	err = r.TableCreate("accounts").Run(session).Exec()
	if err != nil {
		log.Println(err)
	}

	sessionArray = append(sessionArray, session)
}

func mapRoutes() {
	accountsController := &controllers.Accounts{DbSession: sessionArray[0]}

	goweb.MapController("/api/accounts", accountsController)

	goweb.MapStatic("/public", "src")
	// goweb.MapStaticFile("/", "src/index.html")
	goweb.MapStaticFile("/favicon.ico", "src/favicon.ico")

	// Catch-all handler for everything that we don't understand
	goweb.Map(func(c context.Context) error {
		// just return a 404 message
		return goweb.API.Respond(c, 404, nil, []string{"File not found"})
	})
}

func main() {
	initDb()
	mapRoutes()

	log.Print("Starting Goweb powered server...")

	// make a http server using the goweb.DefaultHttpHandler()
	s := &http.Server{
		Addr:           Address,
		Handler:        goweb.DefaultHttpHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	listener, listenErr := net.Listen("tcp", Address)

	log.Printf("  visit: %s", Address)

	if listenErr != nil {
		log.Fatalf("Could not listen: %s", listenErr)
	}

	log.Println("Routes:")
	log.Printf("%s", goweb.DefaultHttpHandler())

	go func() {
		for _ = range c {
			// sig is a ^C, handle it
			// stop the HTTP server
			log.Print("Stopping the server...")
			listener.Close()

			os.Exit(0)
		}
	}()

	// begin the server
	log.Fatalf("Error in Serve: %s", s.Serve(listener))

	/*

	   END OF WEB SERVER CODE

	*/
}
