package main

import (
	"fmt"
	"os"
)

func main() {
	if hostname, err := os.Hostname(); err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Fails to get hostname")
	}
}
