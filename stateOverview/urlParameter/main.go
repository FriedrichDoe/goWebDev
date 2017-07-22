package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

// localhost:8080/?animal=kadse
func foo(w http.ResponseWriter, req *http.Request) {
	a := req.FormValue("animal")
	io.WriteString(w, "Do my search: "+a)
}
