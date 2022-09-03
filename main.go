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

	} else if name == "board_test" { // [O1o1o0g13o0]
		fmt.Print(`= board_test
.`)

		var b = NewBoard()
		var setStone = func(s Stone) {
			fmt.Printf("%v", s)
		}
		var doNewline = func() {
			fmt.Printf("\n.")
		}
		b.ForeachLikeText(setStone, doNewline)

		// この上に分岐を挟んでいく

	} else {
		fmt.Println("go run . {programName}")
	}
}

// EOF [O1o1o0g9o0]
