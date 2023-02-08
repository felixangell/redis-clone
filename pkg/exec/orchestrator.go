package exec

import (
	"github.com/bat-labs/krake/cache"
	"github.com/bat-labs/krake/pkg/api"
	"log"
	"net"
)

type KafkaNodeOrchestrator struct {
	backend cache.Cache
}

func NewKafkaNodeOrchestrator(backend cache.Cache) *KafkaNodeOrchestrator {
	return &KafkaNodeOrchestrator{backend: backend}
}

func (o *KafkaNodeOrchestrator) Submit(conn net.Conn, response api.Value) {
	_, err := conn.Write(response.Serialize())
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("<", response)
}

func (o *KafkaNodeOrchestrator) CacheBackend() cache.Cache {
	return o.backend
}
