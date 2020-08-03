package main

import (
	"io/ioutil"
	"fmt"
	"path/filepath"
	"strings"
	"os"
)


func main() {

	contents, err := ioutil.ReadFile( filepath.Join("/sys/class/net/","enp0s31f6", "device/uevent"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(strings.TrimSpace(string(contents)))

	dest, err := os.Readlink(filepath.Join("/sys/class/net/","enp0s31f6", "device"))
	if err != nil {
		fmt.Println(err)
	}
	pathParts := strings.Split(dest, "/")
	numParts := len(pathParts)
	fmt.Println(pathParts)
	fmt.Println(pathParts[numParts-1])

}

