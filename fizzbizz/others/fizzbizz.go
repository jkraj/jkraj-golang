package main

import "fmt"
import "math/rand"

func main() {

	// Generate a random number
	theno := rand.Intn(100) + 1
	fmt.Println(theno)

	var guess int

	// Get the value from user
	fmt.Print("Your guess in int?")
	_, err := fmt.Scanf("%d", &guess)
	if err != nil {
		fmt.Println("Enter proper value")
	} else {

		// compare
		if guess == theno {
			fmt.Println("You guessed it")
		} else {
			if guess < theno {
				fmt.Println("Fizz")
			} else {
				fmt.Println("Buzz")
			}
		}
	}
}
