package parse

import (
	"fmt"
	"github.com/bat-labs/krake/pkg/api"
	"log"
	"strconv"
)

type parserContext struct {
	pos  int
	data []byte
}

func (p *parserContext) peek(offs ...int) byte {
	offset := 0
	if len(offs) > 0 {
		offset = offs[0]
	}
	return p.data[p.pos+offset]
}

func (p *parserContext) peekAheadBy(offset int) byte {
	return p.data[p.pos+offset]
}

func (p *parserContext) consume() byte {
	curr := p.peek()
	p.pos++
	return curr
}

func (p *parserContext) hasNext() bool {
	return p.pos < len(p.data)
}

func (p *parserContext) parseBulkString() api.Value {
	p.expect(api.BulkStringPrefix)

	strLen := p.parseExpectedNumericValue()

	// parse the string content
	content := make([]byte, strLen)
	for i := 0; i < strLen; i++ {
		content[i] = p.consume()
	}
	p.expectTerminator()

	return api.BulkString{
		Length: strLen,
		Data:   content,
	}
}

func (p *parserContext) parseArray() api.Array {
	p.expect('*')
	arrayLen := p.parseExpectedNumericValue()

	var values []api.Value
	for i := 0; i < arrayLen; i++ {
		values = append(values, p.parseValue())
	}

	return api.Array{
		Length: arrayLen,
		Data:   values,
	}
}

func (p *parserContext) parseInteger() api.Value {
	p.expect(':')
	bytes := p.consumeUntilTerminator()
	p.expectTerminator()
	return api.IntegerValue{
		Data: bytes,
	}
}

func (p *parserContext) parseExpectedNumericValue() int {
	bytes := p.consumeUntilTerminator()
	arrayLen, err := strconv.Atoi(string(bytes))
	if err != nil {
		panic(fmt.Sprintf("bad input: '%s'", string(bytes)))
	}
	p.expectTerminator()
	return arrayLen
}

func (p *parserContext) consumeUntilTerminator() []byte {
	var result []byte
	for p.hasNext() {
		if p.peek() == '\r' && p.hasNext() {
			if p.peek(1) == '\n' {
				return result
			}
		}
		result = append(result, p.consume())
	}

	// End of input?
	return result
}

func (p *parserContext) expect(b byte) byte {
	curr := p.peek()
	if curr == b {
		return p.consume()
	}
	panic(fmt.Sprintf("expected '%d' got '%d'", b, curr))
}

func (p *parserContext) expectTerminator() {
	p.expect('\r')
	p.expect('\n')
}

func (p *parserContext) consumeWhile(pred func(val byte, futureBytes ...byte) bool) []byte {
	var datas []byte

	// FIXME(FELIX): bug waiting to happen where we look ahead
	// by +1 but only check for 1 more char.
	for p.hasNext() && pred(p.peek(), p.peekAheadBy(1)) {
		datas = append(datas, p.consume())
	}
	return datas
}

func (p *parserContext) parseValue() api.Value {
	curr := p.peek()

	switch curr {
	case api.ArrayPrefix:
		return p.parseArray()

	case api.BulkStringPrefix:
		return p.parseBulkString()

	case api.IntegerPrefix:
		return p.parseInteger()
	}

	panic(fmt.Sprintf("Invalid value prefix %q (%d) at [%d]\n%s",
		curr, int(curr), p.pos, p.contextualBlame(p.pos)))
}

func (p *parserContext) contextualBlame(pos int) string {
	underlineSpan := func(s []byte, span [2]int) string {
		// first we replace the double chars
		// (carriage return) with a single one
		for idx, c := range s {
			switch c {
			case '\r':
				s[idx] = '_'
			case '\n':
				s[idx] = '_'
			}
		}

		var underline string
		for i := 0; i < span[0]; i++ {
			underline += " "
		}
		for i := span[0]; i <= span[1]; i++ {
			underline += "^"
		}

		result := fmt.Sprintln(string(s))
		result += underline
		return result
	}

	// FIXME we only underline one character.
	return underlineSpan(p.data, [2]int{pos, pos})
}

func ParseMessage(data []byte) api.Value {
	log.Printf("%q\n", string(data))

	p := parserContext{
		pos:  0,
		data: data,
	}

	if p.hasNext() {
		return p.parseValue()
	}
	panic("Reached end of input")
}
