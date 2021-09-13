package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string

var (
	errorNotFound = errors.New("Not Found")
	errorCantUpdate = errors.New("Can't Update non existing word")
	errorWordExists = errors.New("Word is already exists")
)

// * Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errorNotFound
} 



// * Add a word to the Dictionary
func (d Dictionary) Add(word, def string) error {
	
	_, error := d.Search(word)
	
	switch error {
	case errorNotFound:
		d[word] = def
	case nil:
		return errorWordExists
	}
	return nil
} 

// * Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errorNotFound:
		return errorCantUpdate
	}
	return nil
}
