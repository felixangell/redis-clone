package exec

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/bat-labs/krake/cache"
	"github.com/bat-labs/krake/pkg/api"
	"log"
	"net"
)

type KafkaNodeOrchestrator struct {
	id      string
	backend cache.Cache
}

func generateRunID() (string, error) {
	b := make([]byte, 20)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func NewKafkaNodeOrchestrator(backend cache.Cache) *KafkaNodeOrchestrator {
	id, err := generateRunID()
	if err != nil {
		panic(err)
	}

	return &KafkaNodeOrchestrator{id: id, backend: backend}
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

func (o *KafkaNodeOrchestrator) ClusterState() api.Value {
	return api.KeyValueMap{
		"run_id":     o.id,
		"redis_mode": "standalone",
	}
}
