package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Ng1n3/go-microservices/authentication-service/data"
)

const webPort = "3000"

type Config struct {
  DB *sql.DB
  Models data.Models
}

func main() {
  log.Println("Starting authentication service")

  // TODO connect to DB

  // Set up config
  app := Config{}

  srv := &http.Server{
    Addr: fmt.Sprintf(":%s", webPort),
    Handler: app.routes(),
  }

  err := srv.ListenAndServe()
  if err != nil {
    log.Panic(err)
  }
}