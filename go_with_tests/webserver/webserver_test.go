package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
)

func TestGETPlayers(t *testing.T) {
	server := PlayerServer{&StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd": 10,
		},
	}}
	t.Run("returns Pepper's score", func (t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
	})
	t.Run("return Floyd's score", func (t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%v", name), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}