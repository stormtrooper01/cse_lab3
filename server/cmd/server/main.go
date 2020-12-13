package main

import (
    "database/sql"
    "flag"
    "github.com/stormtrooper01/cse_lab3/server/db"
    "log"
    "net/http"
    "os"
    "os/signal"
)

var httpPortNumber = flag.Int("p", 8000, "HTTP Port Number")

func NewDbConnection() (*sql.DB, error) {
  c := &db.Connection{
    Host: "localhost",
    User: "postgres",
    Password: "121212",
    DbName: "db",
    DisableSSL: true,
  }
  return c.Open()
}

func main() {
  flag.Parse()

  server, err := ComposeApiServer(HttpPortNumber(*httpPortNumber))
  if err == nil {
    go func() {
      log.Println("Starting server")
      err := server.Start()
      if err == http.ErrServerClosed {
        log.Printf("HTTP server stopped")
      } else {
        log.Fatalf("Cannot start HTTP server: %s", err)
      }
    }()

    sigChannel := make(chan os.Signal, 1)
    signal.Notify(sigChannel, os.Interrupt)
    <-sigChannel
    if err := server.Stop(); err != nil && err != http.ErrServerClosed {
      log.Printf("Error stopping the server: %s", err)
    }
  } else {
    log.Fatalf("Cannot initialize banking server: %s", err)
  }
}

