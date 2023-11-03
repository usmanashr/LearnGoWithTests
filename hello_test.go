package main

import (
	"log"
	"testing"
)

func TestHello(t *testing.T) {

	got := Hello("Chrgo tesis")
	want := "Hello World"

	if got != want {
		log.Fatalln("error occured")
	}
}
