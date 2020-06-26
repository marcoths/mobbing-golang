package players

import (
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, req *http.Request) {

	player := strings.TrimPrefix(req.URL.Path, "/players/")
	score := GetPlayerScore(player)
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
