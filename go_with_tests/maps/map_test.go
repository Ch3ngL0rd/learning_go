package maps

import (
	"testing"
	"errors"
)

func TestDictionary(t *testing.T) {
	t.Run("Test fetching words", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("Test fetching non-existant word", func(t *testing.T) {
		word := "hello"
		dictionary := Dictionary{}

		_, err := dictionary.Search(word)

		assertError(t, err, ErrUndefinedWord)
	})
	t.Run("Allow adding words to dictionary", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("Error on adding word twice", func(t *testing.T) {
		word := "zac"
		first_definition := "a chill guy"
		second_definition := "the 'tism strikes again"

		dictionary := Dictionary{word: first_definition}

		err := dictionary.Add(word, second_definition)
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, first_definition)
	})
	t.Run("Updating word definition", func(t *testing.T) {
		dictionary := Dictionary{"Zac", "A chill guy"}
		dictionary.Update("Zac", "A cool guy")
		assertDefinition(t, dictionary, "Zac", "A cool guy")
	})
}

func assertDefinition(t *testing.T, d Dictionary, word, definition string) {
	t.Helper()

	result, err := d.Search(word)

	assertError(t, err, nil)
	assertStrings(t, result, definition)
}

func assertStrings(t *testing.T, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func assertError(t *testing.T, result, expected error) {
	t.Helper()

	if !errors.Is(result, expected) {
		t.Fatalf("expected error %q, got %q", expected, result)
	}
}
