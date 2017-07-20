package main

import (
	"net/http"
	"github.com/satori/go.uuid"
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080	" , nil)
	for {
		fmt.Println("end")
		time.Sleep(1000000000000)
	}
}

func foo (w http.ResponseWriter, req *http.Request) {
	fmt.Println("foo call")
	cookie, err := req.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:	"session",
			Value:	id.String(),
			// Secure: true,
			HttpOnly: 	true,
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
