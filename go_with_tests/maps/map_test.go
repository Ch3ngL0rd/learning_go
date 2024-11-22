package maps

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("Test fetching words", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		result, err := dictionary.Search(word)

		assertError(t, err, nil)
		assertStrings(t, result, definition)
	})
	t.Run("Test fetching non-existant word", func(t *testing.T) {
		word := "hello"
		dictionary := Dictionary{}

		_, err := dictionary.Search(word)

		assertError(t, err, ErrUndefinedWord)
	})
	t.Run("Allow adding words to dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		result, err := dictionary.Search("test")

		assertError(t, err, nil)
		assertStrings(t, result, "this is just a test")
	})
}

func assertStrings(t *testing.T, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func assertError(t *testing.T, result, expected error) {
	t.Helper()

	if result != expected {
		t.Fatalf("expected error %q, got %q", expected, result)
	}
}
