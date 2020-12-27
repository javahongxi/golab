package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func HelloServer(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/hello2", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "<h1>Hello world %s!</h1>", request.FormValue("name"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
