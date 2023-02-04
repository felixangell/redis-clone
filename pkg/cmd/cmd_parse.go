package cmd

import (
	"fmt"
	"github.com/bat-labs/krake/pkg/data"
	"log"
	"reflect"
)

func ParseCommand(v data.Value) (interface{}, error) {
	log.Println("parsing command", v)
	bs, ok := v.(data.Array)
	if !ok {
		panic(fmt.Sprintf("Did not expect value as command %s (type %s)", v, reflect.TypeOf(v)))
	}

	commandName, ok := bs.Data[0].(data.BulkString)
	if !ok {
		panic(fmt.Sprintf("Did not expect value as command %s (type %s)", v, reflect.TypeOf(v)))
	}

	switch string(commandName.Data[:commandName.Length]) {
	case "HELLO":
		fallthrough
	case "hello":
		break
	}

	return nil, nil
}
