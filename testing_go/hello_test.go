package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Katerina")
	want := "Hello Katerina"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
