package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	ps := NewmapStore()
	p1 := &Puppy{ID: 1, Breed: "Dogo", Colour: "white", Value: 50.0}
	_ = ps.CreatePuppy(p1)
	p2, _ := ps.ReadPuppy(p1.ID)
	fmt.Fprintf(out, "Read Puppy by ID: %v\n", p2)
}