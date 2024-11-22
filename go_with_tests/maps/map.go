package maps

import "errors"

type Dictionary map[string]string

var ErrUndefinedWord = errors.New("Word not found in dictionary")

func (dictionary Dictionary) Search(word string) (definition string, err error) {
	value, ok := dictionary[word]
	if !ok {
		return "", ErrUndefinedWord
	}
	return value, nil
}

func (dictionary Dictionary) Add(word, definition string) {
	dictionary[word] = definition
}