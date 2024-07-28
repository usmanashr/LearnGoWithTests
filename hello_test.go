package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		inputString := "Usman"
		got := Hello(inputString)

		want := "Hello Usman"

		if got != want {
			t.Errorf("error occured %q %q", got, inputString)
		}

	})

}
