package exec

import (
	"github.com/bat-labs/krake/pkg/api"
)

type HSetComamnd struct{}

func (h HSetComamnd) Execute(orchestrator *KafkaNodeOrchestrator, args ...api.Value) api.Value {
	hash := string(args[0].(api.BulkString).Data)
	field := string(args[1].(api.BulkString).Data)
	value := args[2]
	orchestrator.CacheBackend().HSet(hash, field, value)

	// success! but... what about failure?
	return api.EncodeInteger(1)
}

type HGetCommand struct{}

func (h HGetCommand) Execute(o *KafkaNodeOrchestrator, args ...api.Value) api.Value {
	hash := string(args[0].(api.BulkString).Data)
	field := string(args[1].(api.BulkString).Data)
	o.CacheBackend().HGet(hash, field)

	// success! but... what about failure?
	return api.EncodeInteger(1)
}

type HDelCommand struct{}

func (h HDelCommand) Execute(o *KafkaNodeOrchestrator, args ...api.Value) api.Value {
	hash := string(args[0].(api.BulkString).Data)

	var fields []string
	for _, a := range args[1:] {
		fields = append(fields, string(a.(api.BulkString).Data))
	}

	err := o.CacheBackend().HDel(hash, fields...)
	if err != nil {
		panic(err)
	}

	// success! but... what about failure?
	return api.EncodeInteger(1)
}
