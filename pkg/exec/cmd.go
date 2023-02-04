package exec

import (
	"github.com/bat-labs/krake/pkg/api"
	"log"
)

type Command interface {
	Execute(*KafkaNodeOrchestrator) api.Value
}

type HelloCommand struct {
}

func (h HelloCommand) Execute(*KafkaNodeOrchestrator) api.Value {
	// HELLO command returns a MAP of the configuration for this
	// cluster according to the redis spec.

	// TODO grab this from orchestrator
	// research all the values to send incl mandatory
	// and whats available.
	configMap := api.KeyValueMap{
		"REDIS000000000017.0": "",
		"run_id":              "54a7ed1c17b0a7f9dbb1286b4e64a14f88c0dde7",
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
	return configMap
}

func NewHelloCommand([]api.Value) HelloCommand {
	return HelloCommand{}
}

type SetCommand struct {
	args []api.Value
}

func (s SetCommand) Execute(orchestrator *KafkaNodeOrchestrator) api.Value {
	// set key, value
	key := string(s.args[0].(api.BulkString).Data)
	log.Println("Received key", key, "setting to", s.args[1])
	orchestrator.Set(key, s.args[1])
	return api.MakeSimpleString("OK")
}

func NewSetCommand(args []api.Value) SetCommand {
	return SetCommand{
		args: args,
	}
}

type GetCommand struct {
	args []api.Value
}

func (g GetCommand) Execute(orchestrator *KafkaNodeOrchestrator) api.Value {
	key := string(g.args[0].(api.BulkString).Data)
	result := orchestrator.Get(key)
	return result
}

func NewGetCommand(args []api.Value) GetCommand {
	return GetCommand{args: args}
}
