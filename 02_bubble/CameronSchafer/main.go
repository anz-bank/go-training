package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	//testing some bubble sorting ideas
	newS, check := s, false

	for !check {
		newS, check = bubbleLoop(newS)
	}

	return s
}

//this function loops through the slice and bubble sorts it.
func bubbleLoop(nextS []int) ([]int, bool) {
	bubbleCheck := true
	newS := nextS

	//bubbleCheck = true

	return newS, bubbleCheck
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
