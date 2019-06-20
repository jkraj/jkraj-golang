package main

import "fmt"

func main() {
	var c chan int
	fmt.Println(c)

	chnl := make(chan int)
	fmt.Printf("Type of `chnl` is %T\n", chnl)
	fmt.Printf("Value of `chnl` is %v\n", chnl)
}
