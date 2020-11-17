package objects

import (
	"fmt"
	"golab/object_storage/chapter2/apiServer/locate"
	"golab/objectstream"
	"io"
)

func getStream(object string) (io.Reader, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object %s locate failed", object)
	}
	return objectstream.NewGetStream(server, object)
}
