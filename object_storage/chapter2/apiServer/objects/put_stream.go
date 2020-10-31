package objects

import (
	"fmt"
	"github.com/javahongxi/golab/common/objectstream"
	"github.com/javahongxi/golab/object_storage/chapter2/apiServer/heartbeat"
)

func putStream(object string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}
	return objectstream.NewPutStream(server, object), nil
}
