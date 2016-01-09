package main

import (
	"net/http"
)

func Hello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("<h1>Hello, world!</h1>"))
}

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}
