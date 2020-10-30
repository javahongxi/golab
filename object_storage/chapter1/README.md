1. set env LISTEN_ADDRESS=:12345 STORAGE_ROOT=/tmp
1. mkdir /tmp/objects
1. curl -v localhost:12345/objects/test
1. curl -v localhost:12345/objects/test -XPUT -d"this is a test object"
1. curl -v localhost:12345/objects/test