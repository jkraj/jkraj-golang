package main

import "fmt"
import "math/rand"
import "github.com/jkraj/jkraj-golang/fizzbizz/firstlib"

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
			m, n := firstlib.Compare(guess, theno)
			fmt.Println(m)
			if n {
				break
			}
		}

	}

}
