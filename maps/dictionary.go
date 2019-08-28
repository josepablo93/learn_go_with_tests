package maps

// Dictionary to look up
type Dictionary map[string]string

// DictionaryErr error of the dictionaries
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	// ErrNotFound error 404
	ErrNotFound = DictionaryErr("could not find the word you were looking for")

	//ErrWordExists error already exists
	ErrWordExists = DictionaryErr("cannot add word because it already exists")

	//ErrWordDoesNotExists error when word isnt in the dictionary and wants to be updated
	ErrWordDoesNotExists = DictionaryErr("cannot update word because it does not exists")
)

// Search looks inside a map and returns the value
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add function adds a new word with its definition to the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update with a new value the definition of a word
func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
