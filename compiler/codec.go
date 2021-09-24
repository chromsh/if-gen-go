package compiler

import (
	"bytes"
	"fmt"
	"io"
)

type Decoder struct {
	data   []byte
	cursor int

	zero []byte
	one  []byte
}

func NewDecoder(data []byte) Decoder {
	zero := " "
	one := "if"
	return Decoder{
		data:   data,
		cursor: 0,
		zero:   []byte(zero),
		one:    []byte(one),
	}
}

// ReadBit read 1bit
func (d *Decoder) ReadBit() (byte, error) {
	if len(d.data) <= d.cursor {
		return 0, io.EOF
	}

	remainBytes := len(d.data) - d.cursor
	// check zero
	if len(d.zero) <= remainBytes {
		bs := d.data[d.cursor : d.cursor+len(d.zero)]
		if bytes.Equal(d.zero, bs) {
			d.cursor += len(d.zero)
			return 0, nil
		}
	}

	// check one
	if len(d.one) <= remainBytes {
		bs := d.data[d.cursor : d.cursor+len(d.one)]
		if bytes.Equal(d.one, bs) {
			d.cursor += len(d.one)
			return 1, nil
		}
	}
	return 0, fmt.Errorf("illegal file format, cursor=%v", d.cursor)
}
