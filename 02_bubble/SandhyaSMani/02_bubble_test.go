package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBubbleSortOutput(t *testing.T) {
	// Given
	r := require.New(t)
	// Then

	expected := []int{1, 2, 3, 4, 5}
	actual := bubble([]int{3, 4, 1, 5, 2})
	r.Equal(expected, actual)

}
func TestBubbleSortOutputForNegativeValues(t *testing.T) {
	// Given
	r := require.New(t)
	// Then

	expected := []int{-10, -5, 0, 20, 100}
	actual := bubble([]int{20, 0, -10, 100, -5})
	r.Equal(expected, actual)

}
func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := strconv.Quote("[1 2 3 4 5]\n")
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in main()")
}