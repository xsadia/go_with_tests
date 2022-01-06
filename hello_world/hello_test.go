package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Felipe", "English")
		want := "Hello, Felipe"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Felipe", "Spanish")
		want := "Hola, Felipe"

		assertCorrectMessage(t, got, want)
	})
}
