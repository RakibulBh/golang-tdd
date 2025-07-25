package main

import (
	"log"
	"net/http"

	"github.com/rakibulbh/golang-tdd/httpserver"
)

func main() {
	handler := http.HandlerFunc(httpserver.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
