package heartbeat

import (
	"golab/rabbitmq"
	"os"
	"time"
)

func StartHeartbeat() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	for {
		q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(5 * time.Second)
	}
}
