// BOF [O1o1o0g9o0]

package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	// プログラム名
	var name = flag.Arg(0)

	if name == "hello" { // [O1o1o0g9o0]
		fmt.Println("Hello, World!")

		// この上に分岐を挟んでいく

	} else {
		fmt.Println("go run . {programName}")
	}
}

// EOF [O1o1o0g9o0]
