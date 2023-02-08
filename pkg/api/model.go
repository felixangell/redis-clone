package api

import (
	"bytes"
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

type Value interface {
	Serialize() []byte
}

type Array struct {
	Length int
	Data   []Value
}

func (a Array) Serialize() []byte {
	buffer := new(bytes.Buffer)
	buffer.WriteByte(ArrayPrefix)
	buffer.WriteString(strconv.Itoa(a.Length))
	buffer.Write(TerminatorByteSequence)

	for _, val := range a.Data {
		buffer.Write(val.Serialize())
	}

	return buffer.Bytes()
}

func NewArray(length int) *Array {
	return &Array{Length: length}
}

type SimpleString struct {
	Length int
	Data   []byte
}

func MakeSimpleString(s string) SimpleString {
	// FIXME simple string is a status or
	// a fixed string, e.g. OK
	return SimpleString{
		Length: len(s),
		Data:   []byte(s),
	}
}

func (s SimpleString) Serialize() []byte {
	var buffer bytes.Buffer
	buffer.WriteByte(byte(SimpleStringPrefix))
	buffer.Write(s.Data)
	buffer.Write(TerminatorByteSequence)
	return buffer.Bytes()
}

type BulkString struct {
	Length int
	Data   []byte
}

func MakeBulkString(s string) BulkString {
	return BulkString{
		Length: len(s),
		Data:   []byte(s),
	}
}

func (b BulkString) Serialize() []byte {
	var buffer bytes.Buffer
	buffer.WriteByte(byte(BulkStringPrefix))
	buffer.WriteString(strconv.Itoa(len(b.Data)))
	buffer.Write(TerminatorByteSequence)

	buffer.Write(b.Data)
	buffer.Write(TerminatorByteSequence)
	return buffer.Bytes()
}

type IntegerValue struct {
	Data []byte
}

func MakeInteger(v int) IntegerValue {
	return IntegerValue{
		Data: []byte(strconv.Itoa(v)),
	}
}

func NewIntegerValue(data []byte) *IntegerValue {
	return &IntegerValue{Data: data}
}

func (i IntegerValue) Serialize() []byte {
	var buffer bytes.Buffer
	buffer.WriteByte(IntegerPrefix)
	buffer.Write(i.Data)
	buffer.Write(TerminatorByteSequence)
	return buffer.Bytes()
}
