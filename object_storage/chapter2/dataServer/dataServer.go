package main

import (
	"golab/object_storage/chapter2/dataServer/heartbeat"
	"golab/object_storage/chapter2/dataServer/locate"
	"golab/object_storage/chapter2/dataServer/objects"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
