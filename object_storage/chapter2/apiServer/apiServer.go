package main

import (
	"golab/object_storage/chapter2/apiServer/heartbeat"
	"golab/object_storage/chapter2/apiServer/locate"
	"golab/object_storage/chapter2/apiServer/objects"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
