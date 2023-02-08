package api

type KeyValueMap map[string]string

func (k KeyValueMap) Serialize() []byte {
	a := NewArray(len(k) * 2)

	for k, v := range k {
		key := EncodeBulkString(k)
		value := EncodeBulkString(v)

		a.Data = append(a.Data, key)
		a.Data = append(a.Data, value)
	}

	return a.Serialize()
}
