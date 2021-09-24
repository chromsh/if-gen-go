package compiler

import (
	"archive/zip"
	"bytes"
	"io"
	"strings"
)

func Compile(r io.Reader) (string, error) {
	buf := bytes.NewBuffer(nil)
	w := zip.NewWriter(buf)

	f, err := w.Create("main.go")
	if err != nil {
		return "", err
	}

	data, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	_, err = f.Write(data)
	if err != nil {
		return "", err
	}
	w.Close()
	return convert(buf.Bytes())
}

func convert(data []byte) (string, error) {
	var sb strings.Builder
	for _, b := range data {
		ifs := convertToIf(b)
		sb.WriteString(ifs)
	}

	return sb.String(), nil
}

func convertToIf(b byte) string {
	var sb strings.Builder
	for i := 0; i < 8; i++ {
		if 0x01&(b>>i) == 1 {
			sb.WriteString("if")
		} else {
			sb.WriteString(" ")
		}
	}
	return sb.String()
}
