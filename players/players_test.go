package players

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Merrick's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Merrick", nil)

		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "15"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns Mikael's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Mikael", nil)

		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}