package exec

import (
	"github.com/bat-labs/krake/pkg/api"
	"log"
)

type Command interface {
	Execute(*KafkaNodeOrchestrator, ...api.Value) api.Value
}

type HelloCommand struct{}

func (h HelloCommand) Execute(o *KafkaNodeOrchestrator, args ...api.Value) api.Value {
	// HELLO command returns a MAP of the configuration for this
	// cluster according to the redis spec.
	return o.ClusterState()
}

type ExistsCommand struct{}

func (e ExistsCommand) Execute(o *KafkaNodeOrchestrator, args ...api.Value) api.Value {
	key := string(args[0].(api.BulkString).Data)
	if o.CacheBackend().Exists(key) {
		return api.EncodeInteger(1)
	}

	return api.EncodeInteger(0)
}

type SetCommand struct{}

func (s SetCommand) Execute(orchestrator *KafkaNodeOrchestrator, args ...api.Value) api.Value {
	// set key, value
	key := string(args[0].(api.BulkString).Data)
	log.Println("Received key", key, "setting to", args[1])
	orchestrator.CacheBackend().Set(key, args[1])
	return api.EncodeSimpleString("OK")
}

type GetCommand struct{}

func (g GetCommand) Execute(orchestrator *KafkaNodeOrchestrator, args ...api.Value) api.Value {
	key := string(args[0].(api.BulkString).Data)
	result := orchestrator.CacheBackend().Get(key)
	return result
}

type DelCommand struct{}

func (d DelCommand) Execute(o *KafkaNodeOrchestrator, args ...api.Value) api.Value {
	key := string(args[0].(api.BulkString).Data)

	err := o.CacheBackend().Del(key)
	if err != nil {
		log.Println(err)
		return api.EncodeInteger(0)
	}

	return api.EncodeInteger(1)
}
