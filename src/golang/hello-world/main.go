package main

import "fmt"

func main() {
	fmt.Println(Hello("Chris"))
}

const englishHelloPrefix = "Welcome, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}
