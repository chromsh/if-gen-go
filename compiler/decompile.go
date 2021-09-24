package compiler

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
)

func Decompile(r io.Reader) ([]byte, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	origData, err := convertToOriginal(data)
	if err != nil {
		return nil, err
	}

	zr, err := zip.NewReader(bytes.NewReader(origData), int64(len(data)))
	if err != nil {
		return nil, err
	}
	for _, f := range zr.File {
		if f.Name != "main.go" {
			continue
		}
		r, err := f.Open()
		if err != nil {
			return nil, err
		}
		extracted, err := io.ReadAll(r)
		if err != nil {
			return nil, err
		}

		return extracted, nil
	}
	return nil, fmt.Errorf("not supported")
}

func convertToOriginal(data []byte) ([]byte, error) {
	ret := make([]byte, 0)
	bitR := NewDecoder(data)
	byteIdx := 7
	byteData := byte(0)

	for {
		bit, err := bitR.ReadBit()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		byteData = byteData | (bit << byteIdx)
		byteIdx--
		if byteIdx < 0 {
			ret = append(ret, byteData)
			byteData = byte(0)
			byteIdx = 7
		}
	}

	return ret, nil
}
