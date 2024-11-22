package maps

import "errors"

type Dictionary map[string]string

func (dictionary Dictionary) Search(word string) (definition string, err error) {
	value, ok := dictionary[word]
	if (!ok) {
		return "", errors.New("Word not found")
	}
	return value, nil
}