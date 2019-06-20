package main

import "fmt"
import "log"
import "net/http"

type myHandler struct { }

func (m myHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello Johnson")
}

func getForm(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, `<form method="POST", action="/postform"><input name="myname"/><input type="submit"/></form>`)
}

func postForm(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Printf("my Name is %s", req.PostFormValue("myname"))
}

func main() {
	log.Println("Starting local REST server")
	http.HandleFunc("/company", func(w http.ResponseWriter, req *http.Request) {
					fmt.Fprintln(w, "Infoblox India, Rajajinagar, Bangalore")
	})
	http.Handle("/", myHandler{})
	http.HandleFunc("/postform", postForm)
	http.HandleFunc("/form", getForm)
	http.ListenAndServe(":8888", nil)
}
