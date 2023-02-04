package cmd

import (
	"fmt"
	"github.com/bat-labs/krake/pkg/api"
	"reflect"
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

	cmd := string(cmdNameMsg.Data[:cmdNameMsg.Length])
	switch cmd {
	case "HELLO":
		fallthrough
	case "hello":
		return NewHelloCommand(), nil
	}

	panic(fmt.Sprintf("Unhandled command %s", cmd))
}
