package main

import "errors"

type Dictionary map[string]string

var (
	WordNotFound = errors.New("could not find the word you were looking for")
	WordAlreadyExist = errors.New("provided word is already present in the dictionary")
	WordNotExist = errors.New("provided word is not present in the dictionary")
)

func (d Dictionary) Search(word string) (string, error) {
	found, ok := d[word]
	if !ok {
		return "", WordNotFound
	}
	return found, nil
}

func (d Dictionary) Add(word, definition string) error {
	if _, ok := d[word]; ok {
		return WordAlreadyExist
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	found, _ := d.Search(word);
	if found != "" {
		d[word] = definition
		return nil
	}
	return WordNotExist
}

func (d Dictionary) Delete(word string) error {
	found, _ := d.Search(word);
	if found != "" {
		delete(d, word)
		return nil
	}
	return WordNotFound
}
