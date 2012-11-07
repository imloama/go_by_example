package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

var pf = fmt.Printf

func main() {
	p := point{1, 2}
	// Straight ahead variable print
	pf("%v\n", p)
	// %+v includes field names
	pf("%+v\n", p)
	// %#v prints the go syntax representation
	pf("%#v\n", p)
	// %T is the type of a given value
	pf("%T\n", p)
	// Boolean print
	pf("%t\n", true)

	// Integer / Binary / Char / Hex / float
	pf("%d\n", 123)
	pf("%b\n", 14)
	pf("%c\n", 33)
	pf("%x\n", 456)
	pf("%f\n", 78.9)

	// Scientific Notation
	pf("%e\n", 123400000.0)
	pf("%E\n", 123400000.0)

	// String printing - q for goSource double quote
	pf("%s\n", "\"string\"")
	pf("%q\n", "\"string\"")

	// Hex rep of a string
	pf("%x\n", "hex this")

	// Print a pointer
	pf("%p\n", &p)

	// Formatted int
	pf("|%6d|%6d|\n", 12, 345)

	// Formatted float
	pf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	pf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// Formatted String
	pf("|%6s|%6s|\n", "foo", "b")
	pf("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s \n", "string")
	pf(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}
