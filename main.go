package main

import (
	
	"github.com/sporto/kic/api"
	"github.com/stretchr/goweb"
	// "github.com/joho/godotenv"
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

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	// production don't use an env file
	// 	log.Println("No .env file found")
	// }

	dbSession, err := api.StartDb("./")
	if err != nil {
		log.Fatal(err)
	}

	api.MapRoutes(dbSession)

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
