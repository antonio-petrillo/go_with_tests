package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func (t *testing.T) {
		got := Hello(english, "Nto")
		want := "Hello, Nto"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func (t *testing.T) {
		got := Hello(english, "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people in french", func (t *testing.T) {
		got := Hello(french, "Nto")
		want := "Bonjour, Nto"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people in italian", func (t *testing.T) {
		got := Hello(italian, "Nto")
		want := "Ciao, Nto"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
