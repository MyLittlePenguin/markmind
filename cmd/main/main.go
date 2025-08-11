package main

import (
	"log"
	"markmind/internal/application/routes"
	"net/http"
)



func main() {
  log.Fatal(http.ListenAndServe(":8080", routes.Router()))
}
