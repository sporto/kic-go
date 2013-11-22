package main

import (
	r "github.com/dancannon/gorethink"
	"github.com/sporto/kic/api"
	"github.com/stretchr/goweb"
	"github.com/joho/godotenv"
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
	dbSession, err := r.Connect(map[string]interface{}{
		"address":  os.Getenv("DB_HOST"),
		"database": os.Getenv("DB_NAME"),
	})
	if err != nil {
		log.Fatal(err)
		log.Println("Most likely RethinkDB is not running")
		return
	}

	_, err = r.DbCreate("kic").Run(dbSession)
	if err != nil {
		log.Println(err)
	}

	_, err = r.Db("kic").TableCreate("accounts").Run(dbSession)
	if err != nil {
		log.Println(err)
	}

	_, err = r.Db("kic").TableCreate("transactions").Run(dbSession)
	if err != nil {
		log.Println(err)
	}

	sessionArray = append(sessionArray, dbSession)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	initDb()
	api.MapRoutes(sessionArray)

	log.Print("Starting Goweb powered server...")

	// make a http server using the goweb.DefaultHttpHandler()
	s := &http.Server{
		Addr:           Address,
		Handler:        goweb.DefaultHttpHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	listener, listenErr := net.Listen("tcp", Address)
	log.Printf("  visit: %s", Address)

	if listenErr != nil {
		log.Fatalf("Could not listen: %s", listenErr)
	}

	log.Println("Routes:")
	log.Printf("%s", goweb.DefaultHttpHandler())

	// listen for exit signal i.e. ctrl + C
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	go func() {

		for _ = range signalChannel {
			// sig is a ^C, handle it
			log.Println("^C")
			// stop the HTTP server
			log.Println("Stopping the server...")
			listener.Close()

			os.Exit(0)
		}
	}()

	// begin the server
	log.Fatalf("Error in Serve: %s", s.Serve(listener))

}
