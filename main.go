// BOF [O1o1o0g9o0]

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	dbg "github.com/muzudho/kifuwarabe-uec14/debugger"
	"github.com/muzudho/kifuwarabe-uec14/kernel"
)

// [O1o1o0g11o_1o0] グローバル変数として、バーチャルIOを１つ新規作成
// アプリケーションの中では 標準入出力は これを使うようにする
var virtualIo = dbg.NewVirtualIO()

func main() {
	// [O1o1o0g11o__10o_5o0] 思考エンジン設定ファイル
	var (
		pEngineFilePath = flag.String("f", "engine.toml", "engine config file path")
		// [O1o1o0g11o__11o6o0] デバッグ用
		pIsDebug = flag.Bool("-d", false, "for debug")
	)
	flag.Parse()
	// プログラム名
	var name = flag.Arg(0)

	// この下に初期設定を追加していく
	// ---------------------------

	// [O1o1o0g11o__10o_5o0] 思考エンジン設定ファイル
	var onError = func(err error) *Config {
		// ログファイルには出力できません。ログファイルはまだ読込んでいません

		// 強制終了
		panic(err)
	}
	var engineConfig = LoadEngineConfig(*pEngineFilePath, onError)

	// [O1o1o0g11o__10o3o0] ログファイル
	var plainTextLogFile, _ = os.OpenFile(engineConfig.GetPlainTextLog(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer plainTextLogFile.Close() // ログファイル使用済み時にファイルを閉じる
	// ログファイル
	var jsonLogFile, _ = os.OpenFile(engineConfig.GetJsonLog(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer jsonLogFile.Close() // ログファイル使用済み時にファイルを閉じる
	// カスタマイズしたロガーを使うなら
	var logg = kernel.NewSugaredLoggerForGame(plainTextLogFile, jsonLogFile) // customized LOGGer

	// [O1o1o0g11o__11o6o0] デバッグ用
	if *pIsDebug {
		virtualIo.ReplaceInputToFileLines("./debug.input.txt")
	}

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

		// [O1o1o0g12o__11o_4o0] 棋譜の初期化に利用
		var onUnknownTurn = func() kernel.Stone {
			var errMsg = fmt.Sprintf("? unexpected play_first:%s", engineConfig.GetPlayFirst())
			logg.C.Info(errMsg)
			logg.J.Infow("error", "play_first", engineConfig.GetPlayFirst())
			panic(errMsg)
		}

		// [O1o1o0g11o_3o0]
		var kernel1 = kernel.NewKernel(engineConfig.GetBoardSize(), engineConfig.GetBoardSize(),
			// [O1o1o0g12o__11o_4o0] 棋譜の初期化
			engineConfig.GetMaxMovesNum(),
			kernel.GetStoneOrDefaultFromTurn(engineConfig.GetPlayFirst(), onUnknownTurn))
		// 設定ファイルの内容をカーネルへ反映
		kernel1.Board.Init(engineConfig.GetBoardSize(), engineConfig.GetBoardSize())

		// [O1o1o0g11o_1o0] コンソール等からの文字列入力
		for virtualIo.ScannerScan() {
			var command = virtualIo.ScannerText()
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
