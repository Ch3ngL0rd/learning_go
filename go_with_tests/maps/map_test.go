package maps

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("Test fetching words", func (t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word : definition}

		result, err := dictionary.Search(word)

		if err != nil {
			t.Errorf("Expected nil error, got %q", err)
		}

		assertStrings(t, result, definition)
	})
	t.Run("Test fetching non-existant word", func (t *testing.T) {
		word := "hello"
		dictionary := Dictionary{}

		_, err := dictionary.Search(word)

		if err == nil {
			t.Errorf("Expected word missing error, got %q", err)
		}
	})
}

func assertStrings(t *testing.T, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}