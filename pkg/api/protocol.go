package api

import (
	"strconv"
)

// DataTypePrefixByte specifies the prefix byte for determining
// the data type of a message
type DataTypePrefixByte byte

const (
	SimpleStringPrefix DataTypePrefixByte = '+'
	BulkStringPrefix                      = '$'
	IntegerPrefix                         = ':'
	ArrayPrefix                           = '*'
	ErrorPrefix                           = '-'
)

var TerminatorByteSequence = []byte{'\r', '\n'}

func MakeSimpleString(s string) SimpleString {
	// FIXME simple string is a status or
	// a fixed string, e.g. OK
	return SimpleString{
		Length: len(s),
		Data:   []byte(s),
	}
}

func MakeBulkString(s string) BulkString {
	return BulkString{
		Length: len(s),
		Data:   []byte(s),
	}
}

func MakeInteger(v int) IntegerValue {
	return IntegerValue{
		Data: []byte(strconv.Itoa(v)),
	}
}

// TODO map...?
