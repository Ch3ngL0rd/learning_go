package maps

import "errors"

type Dictionary map[string]string

var ErrUndefinedWord = errors.New("Word not found in dictionary")

func (d Dictionary) Search(word string) (definition string, err error) {
	value, ok := d[word]
	if !ok {
		return "", ErrUndefinedWord
	}
	return value, nil
}

func (d Dictionary) Add(word, definition string) error {
	d[word] = definition
	return nil
}
