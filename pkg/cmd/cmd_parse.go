package cmd

import (
	"fmt"
	"github.com/bat-labs/krake/pkg/api"
	"github.com/bat-labs/krake/pkg/data"
	"reflect"
)

func ParseCommand(v data.Value) (api.RESP, error) {
	bs, ok := v.(data.Array)
	if !ok {
		panic(fmt.Sprintf("Did not expect value as command %s (type %s)", v, reflect.TypeOf(v)))
	}

	cmdNameMsg, ok := bs.Data[0].(data.BulkString)
	if !ok {
		panic(fmt.Sprintf("Did not expect value as command %s (type %s)", v, reflect.TypeOf(v)))
	}

	cmd := string(cmdNameMsg.Data[:cmdNameMsg.Length])
	switch cmd {
	case "HELLO":
		fallthrough
	case "hello":
		return api.MakeSimpleString("OK"), nil
	}

	panic(fmt.Sprintf("Unhandled command %s", cmd))
}
