package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nllptr/triangles"
)

var a = flag.Int("a", 1, "Side A of the triangle")
var b = flag.Int("b", 1, "Side B of the triangle")
var c = flag.Int("c", 1, "Side C of the triangle")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nThis program tells you what type a triangle is, based on the lengths of its sides.\n\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.Parse()

	triangle := triangles.New(*a, *b, *c)
	if !triangle.Valid() {
		fmt.Printf("The sides a=%d,b=%d,c=%d do not form a valid triangle.\n", *a, *b, *c)
		flag.Usage()
	}
	fmt.Printf("Triangle type with sides a=%d,b=%d,c=%d is: %s\n", *a, *b, *c, triangle.Type())
}
