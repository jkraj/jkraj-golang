package main

import "fmt"
import "math/rand"

func main() {

	// Generate a random number
	theno := rand.Intn(100) + 1
	fmt.Println(theno)

	var guess int
	for i := 0; i < 5; i++ {
		// Get the value from user
		fmt.Println("Your guess in int?")
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("Enter proper value")
			fmt.Println(err)
			continue
		} else {
			// compare
			if guess == theno {
				fmt.Println("You guessed it")
				break
			} else if guess < theno {
				fmt.Println("Fizz")
			} else {
				fmt.Println("Buzz")
			}
		}

	}

}
