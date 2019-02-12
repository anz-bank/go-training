package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fib(7)
}

func fib(n int) {
	n1, n2 := 0, 1
	for i := 0; i < n; i++ {
		n1, n2 = n2, n1+n2
		fmt.Fprint(out, " ", n1)
	}
}
