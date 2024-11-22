package maps

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("Test fetching words", func (t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word : definition}

		result, err := dictionary.Search(word)

		assertError(t, err, nil)
		assertStrings(t, result, definition)
	})
	t.Run("Test fetching non-existant word", func (t *testing.T) {
		word := "hello"
		dictionary := Dictionary{}

		_, err := dictionary.Search(word)

		assertError(t, err, ErrUndefinedWord)
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
		t.Fatal("expected error %q, got %q", expected, result)
	}
}