package objects

import (
	"fmt"
	"golab/object_storage/chapter2/apiServer/heartbeat"
	"golab/objectstream"
)

func putStream(object string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}
	return objectstream.NewPutStream(server, object), nil
}
