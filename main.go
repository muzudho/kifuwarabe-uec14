// BOF [O1o1o0g9o0]

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	// プログラム名
	var name = flag.Arg(0)

	if name == "hello" { // [O1o1o0g9o0]
		fmt.Println("Hello, World!")

		// この上に分岐を挟んでいく

	} else {

		// fmt.Println("go run . {programName}")

		// [O1o1o0g11o_1o0] コンソール等からの文字列入力
		var scanner = bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			var command = scanner.Text()
			var tokens = strings.Split(command, " ")
			switch tokens[0] {

			// この下にコマンドを挟んでいく
			// -------------------------

			// ★
			case "board": // [O1o1o0g13o0]
				fmt.Print(`= board_test'''
. `)

				var b = NewBoard()
				var setStone = func(s Stone) {
					fmt.Printf("%v", s)
				}
				var doNewline = func() {
					fmt.Printf("\n. ")
				}
				b.ForeachLikeText(setStone, doNewline)
				fmt.Print("\n. '''\n")
				// ★

			case "quit": // [O1o1o0g11o_1o0]
				// os.Exit(0)
				return

			// この上にコマンドを挟んでいく
			// -------------------------

			default:
				fmt.Printf("? unknown_command:%s\n\n", tokens[0])
			}
		}

	}
}

// EOF [O1o1o0g9o0]
