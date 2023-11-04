package main

import (
	"log"
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		inputString := "Chrgo tesis"
		got := Hello(inputString)

		if got != inputString {
			log.Fatalln("error occured")
		}
	})

}
