package main

import (
	"github.com/javahongxi/golab/object_storage/chapter1/objects"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
