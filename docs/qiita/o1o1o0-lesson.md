# Step [O1o1o0g1o0] はじめに

👇 Step の変な数字の説明  

📖 [電脳向量表記](https://qiita.com/muzudho1/items/fdbf31e41dd8c247081f)  

👇 練習編を読み終わってるものとする  

📖 [Go [O1o1o0] 目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜練習編＞](https://qiita.com/muzudho1/items/cea62be01f7418bbf150)  

👇 また、技術的でない内容を含むブログを別の場所に 開設する  

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会☆（＾ｑ＾）＜その２＞](http://grayscale2.dou-jin.com/go/%E7%9B%AE%E6%8C%87%E3%81%9B%EF%BC%81%E7%AC%AC%EF%BC%91%EF%BC%94%E5%9B%9E%EF%BC%B5%EF%BC%A5%EF%BC%A3%E6%9D%AF%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF%E3%83%BC%E5%9B%B2%E7%A2%81%E5%A4%A7%E4%BC%9A%E2%98%86%EF%BC%88%EF%BC%BE)  

アーキテクチャ:  

| What is           | This is            |
| ----------------- | ------------------ |
| OS                | Windows            |
| Editor            | Visual Studio Code |
| Program Language  | Go                 |
| Remote Repository | Git Hub            |

👇 まず、ソースの置き場所を決めておく  

📖 [リポジトリ](https://github.com/muzudho/kifuwarabe-uec14)  

## Step [O1o1o0g2o0] ソースの置き場所（ローカル）

Go言語では ローカルPCのどこにソースを置くかは自分で設定して決めておく。  
サンプルでは　ユーザーホームの下に置いているので、真似る  

👇 以下のコマンドを入力してほしい  

Input:  

```shell
# Windows
echo %HOMEPATH%
```

Output:  

```plaintext
\Users\むずでょ
```

ユーザーホームのパスが分かった。この下に `go\src` で始まるディレクトリーを作っていく。  
私は以下の場所にした  

`C:\Users\むずでょ\Documents\GitHub\kifuwarabe-uec14`

以降の文章では、あなたのリポジトリに読み替えてほしい  

## Step [O1o1o0g3o0] Visual Studio Code を使う

がんばって、 `Visual Studio Code` を使えるようにしておいてほしい  

📖 [Visual Studio Code](https://code.visualstudio.com/)  

## Step [O1o1o0g4o0] Goエクステンションをインストールする

`Visual Studio Code` の `Extensions` から、 `Go` をインストールしておいてほしい  

## Step [O1o1o0g5o0] マルチ ワークスペース

べつに ワークスペースは１つでいいが、  
一応 複数のワークスペースを作れるように対応しておく    

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go work init
```

👇 以下のファイルが自動生成される  

```plaintext
  	📂 kifuwarabe-uec14
👉  └── 📄 go.work

```go
go 1.19

```

## Step [O1o1o0g6o0] 設定 - .gitignore ファイル

👇 以下の既存ファイルを編集（無ければ新規作成）してほしい  

```plaintext
  	📂 kifuwarabe-uec14
👉  ├── 📄 .gitignore
  	└── 📄 go.work
```

👇 例えば冒頭に追加  

```plaintext
# この下に kifuwarabe-uec14 でリポジトリにコミットしないものを追加する
# ---------------------------------------------------------------

go.work


# この上に kifuwarabe-uec14 でリポジトリにコミットしないものを追加する
# ---------------------------------------------------------------
# ...略...
```

## Step [O1o1o0g7o0] モジュール作成

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go mod init github.com/muzudho/kifuwarabe-uec14
#           -----------------------------------
#           1
# 1. モジュール名。この部分はあなたのレポジトリに合わせて変えてほしい
```

Output:  

```plaintext
go: creating new go.mod: module github.com/muzudho/kifuwarabe-uec14
go: to add module requirements and sums:
        go mod tidy
```

👇 以下のファイルが自動生成される  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📄 .gitignore
👉  ├── 📄 go.mod
  	└── 📄 go.work
```

```go
module github.com/muzudho/kifuwarabe-uec14

go 1.19
```

## Step [O1o1o0g8o0] ワークスペースズへ登録

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go work use .
```

👇 以下のファイルが自動で編集されている  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📄 .gitignore
    ├── 📄 go.mod
👉 	└── 📄 go.work
```

```go
// ...略...


// * 以下の行が自動追加
use .
```

# Step [O1o1o0g9o0] エントリーポイント作成 - ハローワールド

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📄 .gitignore
    ├── 📄 go.mod
  	├── 📄 go.work
👉 	└── 📄 main.go
```

```go
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

	// この下に初期設定を追加していく
	// ---------------------------

	// この上に初期設定を追加していく
	// ---------------------------

	if name == "hello" { // [O1o1o0g9o0]
		fmt.Println("Hello, World!")

		// この下に分岐を挟んでいく
		// ---------------------

		// この上に分岐を挟んでいく
		// ---------------------

	} else {
		fmt.Println("go run . {programName}")
	}
}

// EOF [O1o1o0g9o0]
```

## Step [O1o1o0g10o0] tidy

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
go mod tidy
```

## Step [O1o1o0g10o1o0] 実行

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go run . hello
```

Output:  

```plaintext
Hello, World!
```

# Step [O1o1o0g11o__10o0] ロガー設定

## Step [O1o1o0g11o__10o1o0] インストール

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go get -u go.uber.org/zap
```

## Step [O1o1o0g11o__10o2_1o0] 設定 - .gitignore ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
👉	├── 📄 .gitignore
    ├── 📄 go.mod
  	├── 📄 go.work
 	├── 📄 logger.go
 	└── 📄 main.go
```

```plaintext
# この下に kifuwarabe-uec14 でリポジトリにコミットしないものを追加する
# ---------------------------------------------------------------
# ...略...

*.log

# この上に kifuwarabe-uec14 でリポジトリにコミットしないものを追加する
# ---------------------------------------------------------------
# ...略...
```

## Step [O1o1o0g11o__10o2o0] ファイル作成

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📄 .gitignore
    ├── 📄 go.mod
  	├── 📄 go.work
👉 	├── 📄 logger.go
 	└── 📄 main.go
```

```go
// BOF [O1o1o0g11o__10o2o0]

package main

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type SugaredLoggerForGame struct {
	c *zap.SugaredLogger // for Console
	j *zap.SugaredLogger // for Json
}

func NewSugaredLoggerForGame(textLogFile *os.File, jsonLogFile *os.File) *SugaredLoggerForGame {
	var slog = new(SugaredLoggerForGame) // Sugared LOGger
	slog.c = createSugaredLoggerForConsole(textLogFile)
	slog.j = createSugaredLoggerAsJson(jsonLogFile)
	return slog
}

// ロガーを作成します，コンソール形式
func createSugaredLoggerForConsole(textLogFile *os.File) *zap.SugaredLogger {
	// 設定，コンソール用
	var configC = zapcore.EncoderConfig{
		MessageKey: "message",

		// LevelKey:    "level",
		// EncodeLevel: zapcore.CapitalLevelEncoder,

		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder, // 日本時間のタイムスタンプ

		// CallerKey:    "caller",
		// EncodeCaller: zapcore.ShortCallerEncoder,
	}

	// 設定、ファイル用
	var configF = zapcore.EncoderConfig{
		MessageKey: "message",

		// LevelKey:    "level",
		// EncodeLevel: zapcore.CapitalLevelEncoder,

		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder, // 日本時間のタイムスタンプ

		// CallerKey:    "caller",
		// EncodeCaller: zapcore.ShortCallerEncoder,
	}

	// コア
	var core = zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(configC), // コンソール形式
			zapcore.Lock(os.Stderr),            // 出力先は標準エラー
			zapcore.DebugLevel),                // ログレベル
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(configF), // コンソール形式
			zapcore.AddSync(textLogFile),       // 出力先はファイル
			zapcore.DebugLevel),                // ログレベル
	)

	// ロガーのビルド
	var logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	// 糖衣構文のインターフェースを取得
	return logger.Sugar()
}

// ロガーを作成します，JSON複数行形式
func createSugaredLoggerAsJson(jsonLogFile *os.File) *zap.SugaredLogger {
	// 設定 > 製品用
	var config = zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder // 日本時間のタイムスタンプ

	// コア
	var core = zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(config), // JSON形式
			zapcore.AddSync(jsonLogFile),   // 出力先はファイル
			zapcore.DebugLevel),            // ログレベル
	)

	// ロガーのビルド
	var logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	// 糖衣構文のインターフェースを取得
	return logger.Sugar()
}

func (logg *SugaredLoggerForGame) Infow(msg string, keysAndValues ...interface{}) {
	var sb strings.Builder
	if msg != "" {
		sb.WriteString(msg)
		sb.WriteString(" ")
	}

	for i, v := range keysAndValues {
		if i == 0 {
			// pass
		} else if i%2 == 0 {
			sb.WriteString(" ")
		} else {
			sb.WriteString(":")
		}

		sb.WriteString(fmt.Sprintf("%v", v))
	}
	logg.c.Infof("%s", sb.String())

	logg.j.Infow(msg, keysAndValues...)
}

// EOF [O1o1o0g11o__10o2o0]
```

## Step [O1o1o0g11o__10o3o0] ファイル編集

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📄 .gitignore
    ├── 📄 go.mod
  	├── 📄 go.work
 	├── 📄 logger.go
👉 	└── 📄 main.go
```

```go
	// ...略...
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
		// ...略...
		// この下に分岐を挟んでいく
		// ---------------------
		// ...略...


	} else if name == "welcome" { // [O1o1o0g11o__10o0]
		logg.c.Infof("Welcome! name:'%s' weight:%.1f x:%d", "nihon taro", 92.6, 3)
		logg.j.Infow("Welcome!",
			"name", "nihon taro", "weight", 92.6, "x", 3)


		// ...略...
		// この上に分岐を挟んでいく
		// ---------------------
		// ...略...
	}
```

# Step [O1o1o0g11o__10o4o0] 実行

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go run . welocome
```

Output:  

```shell
2022-09-11T14:42:53.258+0900    Welcome! a:1 b:2 c:3
```

* 標準出力は、大会サーバーにメッセージを送るのに利用されることがある。従って 標準出力にログを出力すると反則負けになることがある
  * 従って、ログをコンソールに表示したいときは、標準エラーに出力するようにする

👇 以下のファイルが新規作成された  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📄 .gitignore
    ├── 📄 go.mod
  	├── 📄 go.work
👉	├── 📄 kifuwarabe-uec14-json.log
👉	├── 📄 kifuwarabe-uec14.log
 	├── 📄 logger.go
 	└── 📄 main.go
```

👇 📄 `kifuwarabe-uec14-json.log`  

```json
{"level":"info","ts":"2022-09-11T14:43:54.145+0900","caller":"kifuwarabe-uec14/main.go:42","msg":"Welcome!","a":1,"b":2,"c":3}
```

* 作成されるログファイルは JSON形式ではない。 ワンライナーのJSONが複数行並ぶ

👇 📄 `kifuwarabe-uec14.log`  

```plaintext
2022-09-11T14:43:54.112+0900	info	kifuwarabe-uec14/main.go:41	Welcome! a:1 b:2 c:3
```

# Step [O1o1o0g11o_1o0] インタープリター 作成

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📄 .gitignore
    ├── 📄 go.mod
  	├── 📄 go.work
👉 	└── 📄 main.go
```

```go
package main

import (
	"bufio"		// * 自動で追加される
	"flag"
	"fmt"
	"os"		// * 自動で追加される
	"strings"	// * 自動で追加される
)

// ...略...


	if name == "hello" { // [O1o1o0g9o0]
		// ...略...
	} else {

		// * 消す
		// fmt.Println("go run . {programName}")

		// * 追加
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

			case "quit": // [O1o1o0g11o_1o0]
				// os.Exit(0)
				return

			// この上にコマンドを挟んでいく
			// -------------------------

			default:
				logg.c.Infof("? unknown_command command:'%s'\n", tokens[0])
				logg.j.Infow("? unknown_command", "command", tokens[0])
			}
		}
	}


// ...略...
```

# Step [O1o1o0g11o_2o0] 実行

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
abc
```

Output:  

```
? unknown_command:abc
```

👆 そのようなコマンドは定義されていないことが示される  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
quit
```

アプリケーションは終了した  

強制終了したいときは、 `[Ctrl]` キーを押しながら `[C]` キーを押してほしい。  
これを以後 `[Ctrl] + [C]` と表記する  

# Step [O1o1o0g11o0] 石定義 作成

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📄 .gitignore
    ├── 📄 go.mod
  	├── 📄 go.work
 	├── 📄 main.go
👉 	└── 📄 stone.go
```

```go
// BOF [O1o1o0g11o0]

package main

import "fmt"

type Stone uint

const (
	Empty Stone = iota
	Black
	White
	Wall
)

// String - 文字列化
func (s Stone) String() string {
	switch s {
	case Empty:
		return "."
	case Black:
		return "x"
	case White:
		return "o"
	case Wall:
		return "+"
	default:
		panic(fmt.Sprintf("%d", int(s)))
	}
}

// EOF [O1o1o0g11o0]
```

# Step [O1o1o0g12o0] 盤定義 作成

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📄 .gitignore
👉  ├── 📄 board.go
    ├── 📄 go.mod
  	├── 📄 go.work
 	├── 📄 main.go
 	└── 📄 stone.go
```

```go
// BOF [O1o1o0g12o0]

package main

// Board - 盤
type Board struct {
	// 枠付きの横幅
	memoryWidth int

	// 枠付きの縦幅
	memoryHeight int

	// 交点
	nodes []Stone
}

// NewBoard - 新規作成
func NewBoard() *Board {
	var b = new(Board)

	b.memoryWidth = 19 + 2
	b.memoryHeight = 19 + 2
	b.nodes = make([]Stone, b.getMemoryArea())

	// 枠を設定する
	// 上辺、下辺を引く
	{
		var y = 0
		var y2 = b.memoryHeight - 1
		for x := 0; x < b.memoryWidth; x++ {
			var i = (y * b.memoryWidth) + x
			b.nodes[i] = Wall

			i = (y2 * b.memoryWidth) + x
			b.nodes[i] = Wall
		}
	}
	// 左辺、右辺を引く
	{
		var x = 0
		var x2 = b.memoryWidth - 1
		for y := 1; y < b.memoryHeight-1; y++ {
			var i = (y * b.memoryWidth) + x
			b.nodes[i] = Wall

			i = (y * b.memoryWidth) + x2
			b.nodes[i] = Wall
		}
	}

	return b
}

// ForeachLikeText - 枠を含めた各セル
func (b *Board) ForeachLikeText(setStone func(Stone), doNewline func()) {
	for y := 0; y < b.memoryHeight; y++ {
		if y != 0 {
			doNewline()
		}

		for x := 0; x < b.memoryWidth; x++ {
			var i = (y * b.memoryWidth) + x
			var stone = b.nodes[i]
			setStone(stone)
		}
	}
}

// 枠付き盤の面積
func (b *Board) getMemoryArea() int {
	return b.memoryWidth * b.memoryHeight
}

// EOF [O1o1o0g12o0]
```

# Step [O1o1o0g13o0] 盤表示コマンド 作成

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📄 .gitignore
    ├── 📄 board.go
    ├── 📄 go.mod
  	├── 📄 go.work
👉  ├── 📄 main.go
 	└── 📄 stone.go
```

```go
// ...略...


		var board = NewBoard()	// [O1o1o0g13o0]
		// * 以下の行より上
		// var scanner = bufio.NewScanner(os.Stdin)
		// for scanner.Scan() {


			//...略...


			// この下にコマンドを挟んでいく
			// -------------------------

			// * アルファベット順になる位置に、以下のケース文を挿入
			case "board": // [O1o1o0g13o0]
				// 人間向けの出力
				{
					var sb strings.Builder
					sb.WriteString(`= board:'''
. `)

					var setStone = func(s Stone) {
						sb.WriteString(fmt.Sprintf("%v", s))
					}
					var doNewline = func() {
						sb.WriteString("\n. ")
					}
					board.ForeachLikeText(setStone, doNewline)
					sb.WriteString("\n. '''\n")
					logg.c.Info(sb.String())
				}
				// コンピューター向けの出力
				{
					var sb strings.Builder

					var setStone = func(s Stone) {
						sb.WriteString(fmt.Sprintf("%v", s))
					}
					var doNewline = func() {
						// pass
					}
					board.ForeachLikeText(setStone, doNewline)
					logg.j.Infow("output", "board", sb.String())
				}

			// この上にコマンドを挟んでいく
			// -------------------------


// ...略...
```

# Step [O1o1o0g14o0] 実行

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
board
```

Output:  

```plaintext
2022-09-11T19:04:58.192+0900    = board:'''
. +++++++++++++++++++++
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +++++++++++++++++++++
. '''
```

* この出力書式は 私の方法であって、公式大会のものではないことに注意されたい
	* 📖 [思考エンジンの思考ログ仕様（きふわらべ2022年以降）](https://qiita.com/muzudho1/items/ceb6130cf558cd373dd7)

👇 以下のファイルが自動生成された  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📄 .gitignore
    ├── 📄 board.go
    ├── 📄 go.mod
  	├── 📄 go.work
👉 	├── 📄 kifuwarabe-uec14-json.log
👉 	├── 📄 kifuwarabe-uec14.log
  	├── 📄 logger.go
  	├── 📄 main.go
 	└── 📄 stone.go
```

👇 📄 `kifuwarabe-uec14-json.log`  

```plaintext
{"level":"info","ts":"2022-09-11T19:04:58.192+0900","caller":"kifuwarabe-uec14/main.go:57","msg":"input","command":"board"}
{"level":"info","ts":"2022-09-11T19:04:58.195+0900","caller":"kifuwarabe-uec14/main.go:81","msg":"output","board":"++++++++++++++++++++++...................++...................++...................++...................++...................++...................++...................++...................++...................++...................++...................++...................++...................++...................++...................++...................++...................++...................++...................++++++++++++++++++++++"}
```

👇 📄 `kifuwarabe-uec14.log`  

```plaintext
2022-09-11T19:04:58.157+0900	# board
2022-09-11T19:04:58.192+0900	= board:'''
. +++++++++++++++++++++
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +...................+
. +++++++++++++++++++++
. '''
```

`quit` コマンドで 思考エンジンを終了してほしい  

# Step [O1o1o0g15o0] 座標の定義

👇 以下の既存ファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📄 .gitignore
    ├── 📄 board.go
    ├── 📄 go.mod
  	├── 📄 go.work
  	├── 📄 main.go
👉  ├── 📄 point.go
 	└── 📄 stone.go
```

```go
// BOF [O1o1o0g15o0]

package main

// Point - 交点の座標。いわゆる配列のインデックス。壁を含む盤の左上を 0 とします
type Point int

// GetXFromFile - `A` ～ `Z` を 0 ～ 24 へ変換します。 国際囲碁連盟のルールに倣い、筋の符号に `I` は使いません
func GetXFromFile(file string) int {
	// 筋
	var x = file[0] - 'A'
	if file[0] >= 'I' {
		x--
	}
	return int(x)
}

// GetYFromRank - '1' ～ '99' を 0 ～ 98 へ変換します
func GetYFromRank(rank string) int {
	// 段
	var y = int(rank[0] - '0')
	if 1 < len(rank) {
		y *= 10
		y += int(rank[1] - '0')
	}
	return y - 1
}

// GetFileFromCode - 座標の符号の筋の部分を抜き出します
//
// * `code` - 座標の符号。 Example: "A7" や "J13"
func GetFileFromCode(code string) string {
	return code[0:1]
}

// GetRankFromCode - 座標の符号の段の部分を抜き出します
//
// * `code` - 座標の符号。 Example: "A7" や "J13"
func GetRankFromCode(code string) string {
	if 2 < len(code) {
		return code[1:3]
	}

	return code[1:2]
}

// EOF [O1o1o0g15o0]
```

# Step [O1o1o0g16o0] 座標の算出

👇 以下の既存ファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📄 .gitignore
👉  ├── 📄 board_coord.go
    ├── 📄 board.go
    ├── 📄 go.mod
  	├── 📄 go.work
  	├── 📄 main.go
  	├── 📄 point.go
 	└── 📄 stone.go
```

```go
// BOF [O1o1o0g16o0]

package main

// GetPointFromXy - 座標変換 (x,y) → Point
func (b *Board) GetPointFromXy(x int, y int) Point {
	// 枠の厚み 1 を考慮
	return Point((y+1)*b.memoryWidth + x + 1)
}

// GetPointFromCode - "A7" や "J13" といった符号を Point へ変換します
//
// * `code` - 座標の符号。 Example: "A7" や "J13"
func (b *Board) GetPointFromCode(code string) Point {
	return b.GetPointFromXy(
		GetXFromFile(GetFileFromCode(code)),
		GetYFromRank(GetRankFromCode(code)))
}

// EOF [O1o1o0g16o0]
```

# Step [O1o1o0g17o0] 符号変換 作成

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📄 .gitignore
    ├── 📄 board.go
    ├── 📄 go.mod
  	├── 📄 go.work
👉  ├── 📄 main.go
 	└── 📄 stone.go
```

```go
// ...略...


			// この下にコマンドを挟んでいく
			// -------------------------

			// * アルファベット順になる位置に、以下のケース文を挿入
			case "coord": // [O1o1o0g17o0]
				// Example: "coord A7"
				var point = board.GetPointFromCode(tokens[1])
				fmt.Printf("= %d\n", point)

			case "file": // [O1o1o0g17o0]
				// Example: "file A7"
				var file = GetFileFromCode(tokens[1])
				fmt.Printf("= %s\n", file)

			case "rank": // [O1o1o0g17o0]
				// Example: "rank J13"
				var rank = GetRankFromCode(tokens[1])
				fmt.Printf("= %s\n", rank)

			// この上にコマンドを挟んでいく
			// -------------------------


// ...略...
```

# Step [O1o1o0g18o0] 実行

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
file A1
```

Output:  

```plaintext
= A
```

Input:  

```shell
rank A1
```

Output:  

```plaintext
= 1
```

Input:  

```shell
coord A1
```

Output:  

```plaintext
= 22
```

# 参考にした記事

## Go言語

### ファイル分割

📖 [[Go言語] ファイル分割とローカルパッケージ](https://zenn.dev/fm_radio/articles/ca2ff1dfcf89b5)  

### 列挙型

📖 [How to make Go print enum fields as string?](https://stackoverflow.com/questions/41480543/how-to-make-go-print-enum-fields-as-string)  

### 出力

📖 [Go で値を出力する方法](https://golang.keicode.com/basics/go-print-basics.php)  

### 可変長引数

📖 [Concatenating and Building Strings in Go 1.10+](https://www.calhoun.io/concatenating-and-building-strings-in-go/)  
📖 [Convert interface to string](https://yourbasic.org/golang/interface-to-string/)  

.
