package exec

import (
	"fmt"
	"github.com/bat-labs/krake/pkg/api"
	"reflect"
	"strings"
)

func ParseCommand(v api.Value) (Command, error) {
	bs, ok := v.(api.Array)
	if !ok {
		panic(fmt.Sprintf("Did not expect value as command %s (type %s)", v, reflect.TypeOf(v)))
	}

	cmdNameMsg, ok := bs.Data[0].(api.BulkString)
	if !ok {
		panic(fmt.Sprintf("Did not expect value as command %s (type %s)", v, reflect.TypeOf(v)))
	}

	// NOTE do we care about case here?
	cmd := strings.ToLower(string(cmdNameMsg.Data[:cmdNameMsg.Length]))

	switch cmd {
	case "hello":
		return NewHelloCommand(), nil
	case "set":
		return NewSetCommand(), nil
	}

	panic(fmt.Sprintf("Unhandled command %s", cmd))
}
