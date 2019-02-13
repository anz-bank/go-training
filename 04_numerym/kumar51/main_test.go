package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "[a11y K8s abc]\n"

	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestNumeroniumsNoArgs(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := numeronyms()

	// Then
	expected := 0
	actual := len(m)
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestNumeroniumsSingleArg(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := numeronyms("ahajka")

	// Then
	expected := 1
	actual := len(m)
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestGetNumeroniumsEmptyString(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := getResponse("")

	// Then
	expected := ""
	actual := m
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestGetNumeroniumsSingleCharString(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := getResponse("a")

	// Then
	expected := "a"
	actual := m
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestGetNumeroniumsDoubleCharString(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := getResponse("ab")

	// Then
	expected := "ab"
	actual := m
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestGetNumeroniumsThreeCharString(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := getResponse("abc")

	// Then
	expected := "abc"
	actual := m
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestGetNumeroniumsFourCharString(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := getResponse("abcd")

	// Then
	expected := "a2d"
	actual := m
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestGetNumeroniumsMultiCharString(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := getResponse("Global Warming")

	// Then
	expected := "G12g"
	actual := m
	r.Equalf(expected, actual, "Unexpected output in main()")
}
