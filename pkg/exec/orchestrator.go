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
		//"REDIS000000000017.0": "",
		"run_id": o.id,

		"redis_mode": "standalone",

		"cluster_enabled":     "1",
		"cluster_state":       "ok",
		"cluster_slots_pfail": "0",
		"cluster_slots_fail":  "0",
		"cluster_known_nodes": "1",
		"cluster_size":        "1",
		//"cluster_slots_ok":                "16384",
		//"cluster_slots_assigned":          "16384",
		//"cluster_current_epoch":           "7",
		//"cluster_my_epoch":                "6",
		//"cluster_stats_messages_sent":     "2174",
		//"cluster_stats_messages_received": "2173",
	}
}
