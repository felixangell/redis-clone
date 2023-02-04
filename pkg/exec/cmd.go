package exec

import (
	"github.com/bat-labs/krake/pkg/api"
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
		//"cluster_slots_assigned":          "16384",
		//"cluster_slots_ok":                "16384",
		"cluster_slots_pfail": "0",
		"cluster_slots_fail":  "0",
		"cluster_known_nodes": "1",
		"cluster_size":        "1",
		//"cluster_current_epoch":           "7",
		//"cluster_my_epoch":                "6",
		//"cluster_stats_messages_sent":     "2174",
		//"cluster_stats_messages_received": "2173",
	}
	return configMap
}

func NewHelloCommand() HelloCommand {
	return HelloCommand{}
}
