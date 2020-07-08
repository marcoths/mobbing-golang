package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marcoths/mobbing/players"
)

type InMemoryPlayerStore struct {
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) string {
	return "123"
}

func (i *InMemoryPlayerStore) RecordWin(name string) {

}

func main() {
	fmt.Println("running....")
	server := &players.PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":9000", server))
}
