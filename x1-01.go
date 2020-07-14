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
