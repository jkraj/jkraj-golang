package main

import "fmt"
import "log"
import "net/http"

type myHandler struct { }

func (m myHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello Johnson")
}

func main() {
	log.Println("Starting Server")
	http.ListenAndServe(":8888", myHandler{})
}
