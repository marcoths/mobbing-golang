package main

import (
	"log"
	"net/http"

	"github.com/marcoths/mobbing/players"
)

func main() {

	handler := http.HandlerFunc(players.PlayerServer)

	log.Fatal(http.ListenAndServe(":9000", handler))
}