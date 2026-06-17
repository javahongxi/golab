package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func HelloServer(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	log.Printf("Request started: %s", ctx)

	select {
	case <-ctx.Done():
		log.Printf("Request cancelled: %v", ctx.Err())
		return
	case <-time.After(2 * time.Second):
		io.WriteString(writer, "Hello, world!\n")
	}
}

func HelloServerWithContext(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	if deadline, ok := ctx.Deadline(); ok {
		log.Printf("Request deadline: %v", deadline)
	}

	ctx = context.WithValue(ctx, "requestID", fmt.Sprintf("req-%d", time.Now().UnixNano()))

	fmt.Fprintf(writer, "<h1>Hello world %s!</h1>", request.FormValue("name"))
	fmt.Fprintf(writer, "<p>Request ID: %s</p>", ctx.Value("requestID"))
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/hello2", HelloServerWithContext)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
