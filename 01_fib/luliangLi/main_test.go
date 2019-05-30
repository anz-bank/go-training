package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	m_input_output := make(map[int]string)

	m_input_output[7] = "1\n1\n2\n3\n5\n8\n13\n"
	m_input_output[-7] = "0\n1\n-1\n2\n-3\n5\n-8\n13\n"

	testCases := map[string]struct {
		input int
		want  string
	}{
		"fib":     {input: 7, want: "1\n1\n2\n3\n5\n8\n13\n"},
		"negafib":      {input: -7, want: "0\n1\n-1\n2\n-3\n5\n-8\n13\n"},
	}

	for name, test := range testCases {
		// t.Run creates a sub test and runs it like a normal test
		t.Run(name, func(t *testing.T) {
			fib(test.input)

			expected := strconv.Quote(test.want)
			actual := strconv.Quote(buf.String())

			if expected != actual {
				t.Errorf("runing n: %v, expected %v, got %v", test.input, expected, actual)
			}
			buf.Reset()
		})
	}
}
