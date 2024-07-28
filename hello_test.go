package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		inputString := "Usman"
		got := Hello(inputString, "")

		want := "Hello Usman"
		assertCorrectMessage(t, got, want)

	})
	t.Run("print hello world when input string is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Eldoie", "Spanish")
		want := "Hola Eldoie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in frech", func(t *testing.T) {
		got := Hello("Chris", "French")
		want := "Bonjour Chris"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("error occured %q %q", got, want)
	}
}
