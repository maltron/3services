package main

import (
	"net/http"
	"fmt"
)

func main() {
	fmt.Println("Service C running Port 8080")
	http.HandleFunc("/", respond)
	http.ListenAndServe(":8080", nil)
}

func respond(r http.ResponseWriter, w *http.Request) {
	r.Write([]byte("C"))
}