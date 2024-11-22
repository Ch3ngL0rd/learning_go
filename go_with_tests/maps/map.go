package maps

type Dictionary map[string]string

func (dictionary Dictionary) Search(word string) (definition string, err error) {
	return dictionary[word], nil
}