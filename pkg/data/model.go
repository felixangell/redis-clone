package data

type Value interface{}

type Array struct {
	Length int
	Data   []Value
}

type BulkString struct {
	Length int
	Data   []byte
}

type IntegerValue struct {
	Data []byte
}
