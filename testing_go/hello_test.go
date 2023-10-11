package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Katerina")
		want := "Hello Katerina"
		assertCorrectMessage(t, got, want)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hello World' when empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"
		assertCorrectMessage(t, got, want)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Katerina", "Spanish")
		want := "Hola Katerina"
		assertCorrectMessage(t, got, want)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
