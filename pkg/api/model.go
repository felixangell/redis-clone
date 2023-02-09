package api

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

// DataTypePrefixByte specifies the prefix byte for determining
// the data type for a message
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
	IncrBy(value Value) Value
}

type Array struct {
	Length int
	Data   []Value
}

func (a Array) IncrBy(o Value) Value {
	panic("unsupported operation")
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

var legalSimpleStrings = map[string]bool{
	"OK": true,
}

type SimpleString struct {
	Length int
	Data   []byte
}

func (s SimpleString) IncrBy(o Value) Value {
	panic("unsupported operation")
}

func EncodeSimpleString(s string) SimpleString {
	if _, ok := legalSimpleStrings[s]; !ok {
		panic("bad simple string")
	}

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

func (b BulkString) IncrBy(o Value) Value {
	panic("unsupported operation")
}

func EncodeBulkString(s string) BulkString {
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

// TODO(FELIX): Delete me.
func (i IntegerValue) yolo() int {
	curr, err := strconv.Atoi(string(i.Data))
	if err != nil {
		panic(err)
	}
	return curr
}

func (i IntegerValue) IncrBy(o Value) Value {
	curr := i.yolo()

	switch other := o.(type) {
	case IntegerValue:
		otherNum := other.yolo()
		return EncodeInteger(curr + otherNum)
	case BulkString:
		otherNum, err := strconv.Atoi(string(other.Data))
		if err != nil {
			panic(err)
		}
		return EncodeInteger(curr + otherNum)
	}

	panic(fmt.Sprintf("unsupported operation incr by %v (%s)", o, reflect.TypeOf(o)))
}

func EncodeInteger(v int) IntegerValue {
	return IntegerValue{
		Data: []byte(strconv.Itoa(v)),
	}
}

func (i IntegerValue) Serialize() []byte {
	var buffer bytes.Buffer
	buffer.WriteByte(IntegerPrefix)
	buffer.Write(i.Data)
	buffer.Write(TerminatorByteSequence)
	return buffer.Bytes()
}
