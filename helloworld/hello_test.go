package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Nto")
	want := "Hello, Nto!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
