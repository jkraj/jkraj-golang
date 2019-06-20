package main

import "log"
import "net/http"

type myHandler struct { }

func (m myHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "I am throwing error", 403)
}

func main() {
	log.Println("Starting Server")
	http.ListenAndServe(":8888", myHandler{})
}
