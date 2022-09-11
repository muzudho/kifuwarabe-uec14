// BOF [O1o1o0g9o0]

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/muzudho/kifuwarabe-uec14/kernel"
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
	var logg = kernel.NewSugaredLoggerForGame(textLogFile, jsonLogFile) // customized LOGGer

	// この上に初期設定を追加していく
	// ---------------------------

	if name == "hello" { // [O1o1o0g9o0]
		fmt.Println("Hello, World!")

		// この下に分岐を挟んでいく
		// ---------------------

	} else if name == "welcome" { // [O1o1o0g11o__10o0]
		logg.C.Infof("Welcome! name:'%s' weight:%.1f x:%d", "nihon taro", 92.6, 3)
		logg.J.Infow("Welcome!",
			"name", "nihon taro", "weight", 92.6, "x", 3)

		// この上に分岐を挟んでいく
		// ---------------------

	} else {

		// fmt.Println("go run . {programName}")

		// [O1o1o0g11o_3o0]
		var kernel1 = kernel.NewKernel()

		// [O1o1o0g11o_1o0] コンソール等からの文字列入力
		var scanner = bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			var command = scanner.Text()
			logg.C.Infof("# %s", command)             // 人間向けの出力
			logg.J.Infow("input", "command", command) // コンピューター向けの出力

			// [O1o1o0g11o_3o0]
			var isHandled = kernel1.Execute(command, logg)
			if isHandled {
				continue
			}

			// [O1o1o0g11o_1o0]
			var tokens = strings.Split(command, " ")
			switch tokens[0] {

			// この下にコマンドを挟んでいく
			// -------------------------

			case "quit": // [O1o1o0g11o_1o0]
				// os.Exit(0)
				return

			// この上にコマンドを挟んでいく
			// -------------------------

			default: // [O1o1o0g11o_1o0]
				logg.C.Infof("? unknown_command command:'%s'\n", tokens[0])
				logg.J.Infow("? unknown_command", "command", tokens[0])
			}
		}
	}
}

// EOF [O1o1o0g9o0]
