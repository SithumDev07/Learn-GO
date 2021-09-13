package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string

var errorNotFound = errors.New("Not Found")

// * Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errorNotFound
} 

var errorWordExists = errors.New("Word is already exists")


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