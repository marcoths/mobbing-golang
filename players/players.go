package players

import (
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) string
	RecordWin(name string)
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.Store.GetPlayerScore(player)
	if score == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	w.Write([]byte(score))
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
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
