package maps

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("Test fetching words", func (t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word : definition}

		result := dictionary.Search(word)

		assertStrings(t, result, definition)
	})
}

func assertStrings(t *testing.T, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}