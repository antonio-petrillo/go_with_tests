package main

import (
	"testing"
)

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want  {
		t.Errorf("got %q want %q given, %q", got.Error(), want.Error(), "teskt")
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T){
		_, err := dictionary.Search("unknown")
		assertError(t, err, WordNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		addErr := dictionary.Add("test", "this is just a test")

		if addErr != nil {
			t.Fatal("should add a new key", addErr)
		}

		want := "this is just a test"
		got, err := dictionary.Search("test")

		if err != nil {
			t.Fatal("should find added word:", err)
		}

		assertString(t, got, want)
	})

	t.Run("add should fail on double insert", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is just a test")
		if err != nil {
			t.Fatal("should add a new key", err)
		}
		err = dictionary.Add("test", "key already exists")
		assertError(t, err, WordAlreadyExist)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update success", func (t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefintion := "new definition"

		err := dictionary.Update(word, newDefintion)

		if err != nil {
			t.Errorf("din't expect an error, got %q", err)
		}

		want := newDefintion
		got, _ := dictionary.Search("test")

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test")
		}
	})

	t.Run("update failure", func (t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		newDefintion := "new definition"

		err := dictionary.Update(word, newDefintion)

		if err == nil {
			t.Errorf("Expected %q error got nothing", WordNotExist)
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete success", func (t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Delete(word)

		if err != nil {
			t.Errorf("din't expect an error, got %q", err)
		}

		_, shouldBeNotFount := dictionary.Search("test")

		if shouldBeNotFount != nil && shouldBeNotFount != WordNotFound {
			t.Errorf("got %q want %q given, %q", shouldBeNotFount, WordNotFound, "test")
		}
	})

	t.Run("update failure", func (t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Delete("Test")

		if err == nil {
			t.Errorf("Expected %q got nil", WordNotExist)
		}
	})
}
