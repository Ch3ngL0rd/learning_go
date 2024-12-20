package maps

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrUndefinedWord = DictionaryErr("Word not found in dictionary")
	ErrWordExists    = DictionaryErr("Word already exists in dictionary")
	ErrWordNotFound  = DictionaryErr("Cannot update word without definition")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (definition string, err error) {
	value, ok := d[word]
	if !ok {
		return "", ErrUndefinedWord
	}
	return value, nil
}

func (d *Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrUndefinedWord:
		(*d)[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d *Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrUndefinedWord:
		return ErrWordNotFound
	case nil:
		(*d)[word] = definition
		return nil
	default:
		return err
	}
}

func (d *Dictionary) Delete(word string) {
	delete((*d), word)
}