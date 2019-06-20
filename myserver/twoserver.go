package main

import "fmt"
import "log"
import "net/http"

type myHandler struct { }

func (m myHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello Johnson")
}

func main() {
	log.Println("Starting local Server two supports / & /company")
	http.HandleFunc("/company", func(w http.ResponseWriter, req *http.Request) {
					fmt.Fprintln(w, "Infoblox India, Rajajinagar, Bangalore")
	})
	http.Handle("/", myHandler{})
	http.ListenAndServe(":8888", nil)
}
