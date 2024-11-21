package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to friends", func(t *testing.T) {
		got := Hello("Bob", "English")
		want := "Hello, Bob"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' in French", func(t *testing.T) {
		got := Hello("Zac", "French")
		want := "Bonjour, Zac"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' in Dog", func(t *testing.T) {
		got := Hello("Zac", "Dog")
		want := "Ruff, Zac"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
