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

	cmd := strings.ToLower(string(cmdNameMsg.Data[:cmdNameMsg.Length]))

	switch cmd {

	// internal

	case "hello":
		return NewHelloCommand(bs.Data[1:]), nil

	// key based operations

	case "set":
		return NewSetCommand(bs.Data[1:]), nil

	case "get":
		return NewGetCommand(bs.Data[1:]), nil

	case "del":
		panic("not yet implemented")

	// unordered hash set operations

	case "hset":
		// hset hash-key k1 v1
		panic("not yet implemented")

	case "hget":
		// hget hash-key k1
		panic("not yet implemented")

	case "hgetall":
		// entire hash content
		panic("not yet implemented")

	case "hdel":
		// hdel hash-key k1
		panic("not yet implemented")

		// sets

		// bitmap

		// zsets
	}

	panic(fmt.Sprintf("Unhandled command %s", cmd))
}
