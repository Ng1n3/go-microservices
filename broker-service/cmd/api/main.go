package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "3000"

type Config struct {
}

func main() {
	app := Config{}

	log.Printf("Starting Broker service on port %s\n", webPort)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
