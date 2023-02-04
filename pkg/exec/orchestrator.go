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

func (o *KafkaNodeOrchestrator) Submit(conn net.Conn, command Command) {
	response := command.Execute(o)
	_, err := conn.Write(response.Serialize())
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("<", response)
}

func (o *KafkaNodeOrchestrator) Set(key string, value api.Value) {
	o.backend.Set(key, value)
}

func (o *KafkaNodeOrchestrator) Get(key string) api.Value {
	return o.backend.Get(key)
}
