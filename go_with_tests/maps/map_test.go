package maps

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("Test fetching words", func (t *testing.T) {
		dictionary := map[string]string{"test": "this is just a test"}

		got := Search(dictionary, "test")
		expected := "this is just a test"

		if got != expected {
			t.Errorf("expected %q, got %q, given %q", expected, got, "test")
		}
	})
}