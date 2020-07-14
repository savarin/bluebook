// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	// "testing"
)

func main1() {
	s, sep := "", ""

	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

func main2() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}

func main3() {
	fmt.Println(strings.Join(os.Args, " "))
}

func main() {
	// main1()
	// main2()
	main3()
}
