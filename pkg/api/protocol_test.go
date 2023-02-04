package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBulkStringEncoding(t *testing.T) {
	assert.Equal(t, "$5\r\nhello\r\n", string(MakeBulkString("hello").Serialize()))
	assert.Equal(t, "$0\r\n\r\n", string(MakeBulkString("").Serialize()))
}
