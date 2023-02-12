package exec

import (
	"errors"
	"fmt"
	"github.com/bat-labs/krake/pkg/api"
	"reflect"
	"strings"
)

var supportedCommands = map[string]Command{
	"set": SetCommand{},
	"get": GetCommand{},
	"del": DelCommand{},

	"incrby": IncryByCommand{},

	"hset": HSetComamnd{},
	"hget": HGetCommand{},
	"hdel": HDelCommand{},

	"hello": HelloCommand{},

	"exists":  ExistsCommand{},
	"persist": PersistCommand{},

	"expire": ExpireCommand{},
	"ttl":    nil,
	"incr":   nil,
	"decr":   nil,

	"rpush":  nil,
	"lpush":  nil,
	"lrange": nil,
	"lindex": nil,

	"hmget": nil,
	"hmset": nil,

	"sadd":      nil,
	"smembers":  nil,
	"sismember": nil,

	"zadd":   nil,
	"zrange": nil,
}

type ParsedCommandResult struct {
	Command   Command
	Arguments []api.Value
}

func ParseCommand(v api.Value) (*ParsedCommandResult, error) {
	bs, ok := v.(api.Array)
	if !ok {
		panic(fmt.Sprintf("Did not expect value as command %s (type %s)", v, reflect.TypeOf(v)))
	}

	cmdNameMsg, ok := bs.Data[0].(api.BulkString)
	if !ok {
		panic(fmt.Sprintf("Did not expect value as command %s (type %s)", v, reflect.TypeOf(v)))
	}

	commandName := strings.ToLower(string(cmdNameMsg.Data[:cmdNameMsg.Length]))

	command, ok := supportedCommands[commandName]
	if ok {
		return &ParsedCommandResult{
			Command:   command,
			Arguments: bs.Data[1:],
		}, nil
	}

	return nil, errors.New("Unsupported command!")
}
