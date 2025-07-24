package maps

import "errors"

type Dictionary map[string]string

const (
	ErrNotFound        = DictionaryErr("could not find the word you were looking for")
	ErrKeyExists       = DictionaryErr("cannot add word because it already exists")
	ErrKeyDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}

	return definition, nil
}

func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrKeyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrKeyDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrKeyDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}
