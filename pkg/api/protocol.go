package api

import (
	"bytes"
	"github.com/bat-labs/krake/pkg/data"
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

type RESP interface {
	Serialize() []byte
}

type SimpleString struct {
	Value string
}

func MakeSimpleString(s string) SimpleString {
	return SimpleString{s}
}

func (s SimpleString) Serialize() []byte {
	var buffer bytes.Buffer
	buffer.WriteByte(byte(SimpleStringPrefix))
	buffer.WriteString(s.Value)
	buffer.Write(TerminatorByteSequence)
	return buffer.Bytes()
}

type BulkString struct {
	Value string
}

func (b BulkString) AsModel() data.BulkString {
	return data.BulkString{
		Length: len(b.Value),
		Data:   []byte(b.Value),
	}
}

func (b BulkString) Serialize() []byte {
	var buffer bytes.Buffer
	buffer.WriteByte(byte(BulkStringPrefix))
	buffer.WriteString(strconv.Itoa(len(b.Value)))
	buffer.Write(TerminatorByteSequence)

	buffer.Write([]byte(b.Value))
	buffer.Write(TerminatorByteSequence)
	return buffer.Bytes()
}

func MakeBulkString(s string) BulkString {
	return BulkString{
		Value: s,
	}
}
