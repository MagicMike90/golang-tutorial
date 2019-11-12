package main

import (
	"fmt"
)

func main() {
	const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
	fmt.Println(sLiteral)
	// Using %x in fmt.Printf() will return the AB part of a \xAB sequence.
	fmt.Printf("x: %x\n", sLiteral)

	fmt.Printf("sLiteral length: %d\n", len(sLiteral))
	for i := 0; i < len(sLiteral); i++ {
		fmt.Printf("%x ", sLiteral[i])
	}
	fmt.Println()

	// Using %q in fmt.Printf() with a string argument will print a double-quoted string safely escaped with Go syntax
	fmt.Printf("q: %q\n", sLiteral)
	// Using %+q in fmt.Printf() with a string argument will guarantee ASCII-only output.
	fmt.Printf("+q: %+q\n", sLiteral)
	// using % x (note the space between the % character and the x character) in fmt.Printf() will put spaces between the printed bytes
	fmt.Printf(" x: % x\n", sLiteral)

	fmt.Printf("s: As a string: %s\n", sLiteral)

	s2 := "€£³"
	for x, y := range s2 {
		//  %#U will print the characters in the U+0058 format
		fmt.Printf("%#U starts at byte position %d\n", y, x)
	}

	fmt.Printf("s2 length: %d\n", len(s2))
	const s3 = "ab12AB"
	fmt.Println("s3:", s3)
	fmt.Printf("x: % x\n", s3)

	fmt.Printf("s3 length: %d\n", len(s3))

	for i := 0; i < len(s3); i++ {
		fmt.Printf("%x ", s3[i])
	}
	fmt.Println()
}
