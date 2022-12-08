package main

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type Writer string

func (x *Writer) Write(p []byte) (n int, err error) {
	*x = Writer(p)
	return len(p), nil
}

func TestSolution(t *testing.T) {
	tc := []struct {
		input, output string
	}{}
	w := new(Writer)
	for i := range tc {
		in, out = bufio.NewReader(strings.NewReader(tc[i].input)), bufio.NewWriterSize(w, len(tc[i].output))
		main()
		assert.Equal(t, string(*w), tc[i].output)
	}
}
