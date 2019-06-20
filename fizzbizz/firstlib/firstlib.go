// This is my firstlib
package firstlib

// Compare the 2 number if matches return true else false
func Compare(guess int, theno int) (string, bool) {
	if guess == theno {
		return "You guessed it", true
	} else if guess < theno {
		return "Fizz", false
	} else {
		return "Buzz", false
	}
}
