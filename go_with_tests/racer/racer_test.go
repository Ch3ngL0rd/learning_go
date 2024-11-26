package racer

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compare speed of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0)
	
		defer slowServer.Close()
		defer fastServer.Close()
		
		slowURL := slowServer.URL
		fastURL := fastServer.URL
	
		want := fastURL
		got, _ := Racer(slowURL, fastURL)
	
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("timeout if a server does not respond within 10 seconds", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}