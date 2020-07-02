package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marcoths/mobbing/players"
)

type InMemoryPlayerStore struct {

}

func (receiver *InMemoryPlayerStore) GetPlayerScore(name string) string {
	return "123"
}

func main() {
	fmt.Println("running....")
	server := &players.PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":9000", server))
}