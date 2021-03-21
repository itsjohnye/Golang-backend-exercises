package main

import (
	"fmt"
	"os"
)

func main() {

	v := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	values := []int{2, 4, 6, 8}

	fmt.Println(v)
	fmt.Printf("total value is %v\n", sum(values...))

	linenum, name, ab := 15, "count", "not found"
	errorf(linenum, "undedined: %s", name, ab)
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)

}
