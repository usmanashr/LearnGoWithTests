package main

import "fmt"

const namingPrefix = "Hello "

func main() {
	fmt.Println(Hello("Chris"))
}

func Hello(name string) string {
	if len(name) == 0 {
		name = "World"
	}
	return namingPrefix + name
}
