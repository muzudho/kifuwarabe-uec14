// BOF [O1o1o0g9o0]

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
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
	var logFile, _ = os.OpenFile("kifuwarabe-uec14.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer logFile.Close()  // ログファイル使用済み時にファイルを閉じる
	log.SetOutput(logFile) // ロガーにログファイルを紐づけ
	// カスタマイズしたロガーを使うなら
	var slog = CreateSugaredLogger(logFile) // Sugared LOGger

	// この上に初期設定を追加していく
	// ---------------------------

	if name == "hello" { // [O1o1o0g9o0]
		fmt.Println("Hello, World!")

		// この下に分岐を挟んでいく
		// ---------------------

	} else if name == "welcome" { // [O1o1o0g11o__10o0]
		slog.Infow("Welcome!",
			"a", 1, "b", 2, "c", 3)

		// この上に分岐を挟んでいく
		// ---------------------

	} else {

		// fmt.Println("go run . {programName}")
		var board = NewBoard() // [O1o1o0g13o0]

		// [O1o1o0g11o_1o0] コンソール等からの文字列入力
		var scanner = bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			var command = scanner.Text()
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
				fmt.Printf("? unknown_command:%s\n\n", tokens[0])
			}
		}

	}
}

// EOF [O1o1o0g9o0]
