package parse

import (
	"github.com/bat-labs/krake/pkg/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseArray(t *testing.T) {
	p := parserContext{
		pos: 0,
		data: []byte{
			'*', '1', '\r', '\n',
			'$', '3', '\r', '\n',
			'S', 'E', 'T', '\r', '\n',
		},
	}

	result := p.parseArray()
	assert.Equal(t, 1, result.Length)
	assert.Equal(t, api.MakeBulkString("SET"), result.Data[0])
}
