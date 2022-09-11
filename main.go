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

	// この下に初期設定を追加していく
	// ---------------------------

	// ログファイル
	var textLogFile, _ = os.OpenFile("kifuwarabe-uec14.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer textLogFile.Close() // ログファイル使用済み時にファイルを閉じる
	// ログファイル
	var jsonLogFile, _ = os.OpenFile("kifuwarabe-uec14-json.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer jsonLogFile.Close() // ログファイル使用済み時にファイルを閉じる
	// カスタマイズしたロガーを使うなら
	var logg = NewSugaredLoggerForGame(textLogFile, jsonLogFile) // customized LOGGer

	// この上に初期設定を追加していく
	// ---------------------------

	if name == "hello" { // [O1o1o0g9o0]
		fmt.Println("Hello, World!")

		// この下に分岐を挟んでいく
		// ---------------------

	} else if name == "welcome" { // [O1o1o0g11o__10o0]
		logg.c.Infof("Welcome! name:'%s' weight:%.1f x:%d", "nihon taro", 92.6, 3)
		logg.j.Infow("Welcome!",
			"name", "nihon taro", "weight", 92.6, "x", 3)
		// logg.c.Infof("Welcome! a:%d b:%d c:%d", 1, 2, 3)
		// logg.j.Infow("Welcome!",
		//  	"a", 1, "b", 2, "c", 3)

		// この上に分岐を挟んでいく
		// ---------------------

	} else {

		// fmt.Println("go run . {programName}")
		var board = NewBoard() // [O1o1o0g13o0]

		// [O1o1o0g11o_1o0] コンソール等からの文字列入力
		var scanner = bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			var command = scanner.Text()
			logg.c.Infof("# %s", command)
			logg.j.Infow("Input", "Command", command)

			var tokens = strings.Split(command, " ")
			switch tokens[0] {

			// この下にコマンドを挟んでいく
			// -------------------------

			case "board": // [O1o1o0g13o0]
				fmt.Print(`= board_test'''
. `)

				var setStone = func(s Stone) {
					fmt.Printf("%v", s)
				}
				var doNewline = func() {
					fmt.Printf("\n. ")
				}
				board.ForeachLikeText(setStone, doNewline)
				fmt.Print("\n. '''\n")

			case "coord": // [O1o1o0g17o0]
				// Example: "coord A7"
				var point = board.GetPointFromCode(tokens[1])
				fmt.Printf("= %d\n", point)

			case "file": // [O1o1o0g17o0]
				// Example: "file A7"
				var file = GetFileFromCode(tokens[1])
				fmt.Printf("= %s\n", file)

			case "quit": // [O1o1o0g11o_1o0]
				// os.Exit(0)
				return

			case "rank": // [O1o1o0g17o0]
				// Example: "rank J13"
				var rank = GetRankFromCode(tokens[1])
				fmt.Printf("= %s\n", rank)

			// この上にコマンドを挟んでいく
			// -------------------------

			default:
				logg.c.Infof("? unknown_command command:%s\n", tokens[0])
				logg.j.Infow("? unknown_command", "Command", tokens[0])
			}
		}
	}
}

// EOF [O1o1o0g9o0]
