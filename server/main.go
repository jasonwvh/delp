package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jasonw/delp/server/app"
	"github.com/jasonw/delp/server/blockchain"
)

var privateKey = "c018ebe754eea79fc412fd1ec370d8f92d9b22e92487eafc9d7a13ae964df015"
var contractAddress = "0x2038818faD8D2E54CEA52F8501eD861aD205E01A"
var ethNetwork = "http://192.168.1.14:8545"

func main() {
	l := &log.Logger{}
	sm := mux.NewRouter()

	conn, err := blockchain.NewConnection(contractAddress, privateKey, ethNetwork)
	if err != nil {
		panic(err)
	}

	app := &app.App{
		Router:     sm,
		Connection: conn,
	}

	app.SetupApp()

	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	h := ch(app.Router)

	s := http.Server{
		Addr:         ":9090",
		Handler:      h,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		fmt.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Print("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	sig := <-c
	log.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
