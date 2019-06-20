package main

import "fmt"

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

func main() {
	fmt.Println("main started")
	c := make(chan string)
	go greet(c)

	c <- "Johnson"
	fmt.Println("main stopped")
}


