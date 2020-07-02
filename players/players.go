package players

import (
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) string
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := p.Store.GetPlayerScore(player)
	if score == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	w.Write([]byte(score))
}

func GetPlayerScore(name string) string {
	if name == "Mikael" {
		return "20"
	}

	if name == "Merrick" {
		return "15"
	}

	return ""
}
