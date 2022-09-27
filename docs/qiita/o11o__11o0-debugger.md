目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O11o__11o0] デバッグ可能標準入力 作成

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# Step [O11o__11o0] デバッグ可能標準入力 作成

## Step [O11o__11o1o0] git向け対応 - .gitignore ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	└── 📄 logger.go
👉  ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```plaintext
# この下に kifuwarabe-uec14 でリポジトリにコミットしないものを追加する
# ---------------------------------------------------------------
# ...略...

# [O11o__11o1o0]
*.input.txt

# この上に kifuwarabe-uec14 でリポジトリにコミットしないものを追加する
# ---------------------------------------------------------------
# ...略...
```

## Step [O11o__11o2o_1o0] データファイル作成 - debugger/test.input.txt ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📂 debugger
👉  │   └── 📄 test.input.txt
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	└── 📄 logger.go
  	├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```plaintext
10
```

`*.input.txt` というファイル名は、内容が読み取られるとともに空っぽに消される目印にしている。消えて困る内容を書かないように注意してほしい  

## Step [O11o__11o2o_1o0] バーチャルIO作成 - debugger/virtual_io.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📂 debugger
  	│   ├── 📄 test.input.txt
👉  │   └── 📄 virtual_io.go
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	└── 📄 logger.go
  	├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```go
// BOF [O11o__11o2o_1o0]

package debugger

import (
	"bufio"
	"os"
	"regexp"
	"time"
)

// VirtualIO - 入出力を模擬したもの
//
// デバッガーの dlv が標準入力を読めないので、ファイル入力に置き換えるための仕組みです
type VirtualIO struct {
	scanner *bufio.Scanner
	writer  *bufio.Writer

	inputFilePath string
	inputLines    []string
	pollingTime   time.Duration
}

// 新規作成
//
// - 行読取
//
// Parameters
// ----------
// setVIO - 初期化に使える
func NewVirtualIO() *VirtualIO {
	// 実体をメモリ上に占有させる
	//
	// - 規定値：標準入出力
	var virtualIo = VirtualIO{
		scanner:       bufio.NewScanner(os.Stdin),
		writer:        bufio.NewWriter(os.Stdout),
		inputFilePath: "",
		inputLines:    []string{},
		// 1秒は長いが、しかたない
		pollingTime: 1 * time.Second,
	}

	// virtualIo.Scanner.Split(bufio.ScanWords) // 空白で区切る
	virtualIo.scanner.Split(bufio.ScanLines) // 改行で区切る
	// 入力バッファーのサイズを巨大にする
	virtualIo.scanner.Buffer([]byte{}, 100000007)

	// バーチャルIOのアドレスを返す
	return &virtualIo
}

// IsEmpty - 空っぽか
func (vio *VirtualIO) IsEmpty() bool {
	// １行以上存在し、０行目が空行なら、詰める
	for len(vio.inputLines) != 0 && vio.inputLines[0] == "" {
		vio.inputLines = vio.inputLines[1:len(vio.inputLines)]
	}

	// ０行なら空っぽ
	return len(vio.inputLines) == 0
}

// ReplaceInputToFileLines - 標準入力を使うのをやめ、ファイルの先頭行から１行ずつ切り取る方法に変えます
//
// Parameters
// ----------
// inputFilePath - ファイルパス
func (vio *VirtualIO) ReplaceInputToFileLines(inputFilePath string) {
	vio.inputFilePath = inputFilePath
}

var re = regexp.MustCompile("\r\n|\n")

func (vio *VirtualIO) ScannerScan() bool {

	// テキストファイルから読み込むなら
	if vio.inputFilePath != "" {

		var popAllLines = func() []string {
			// ファイル読込
			var bytes, err = os.ReadFile(vio.inputFilePath)
			if err != nil {
				panic(err)
			}

			var text = string(bytes)

			// ファイルを空にする
			os.Truncate(vio.inputFilePath, 0)

			// 全文を改行でスプリット
			return re.Split(text, -1)
		}

		// バッファーが空なら、ファイルから取ってくる
		if vio.IsEmpty() {
			// 全行取得
			vio.inputLines = popAllLines()
		}

		// バッファーが空の間ブロック（繰り返し）する
		for vio.IsEmpty() {
			// スリープする。なぜなら、入力がないときブロックするという機能を入れないと、無限に空文字列を読み続けてしまうから
			time.Sleep(vio.pollingTime)

			// 全行取得
			vio.inputLines = popAllLines()
		}

		return true
	}

	return vio.scanner.Scan()
}

func (vio *VirtualIO) ScannerText() string {

	// テキストファイルから読み込むなら
	if vio.inputFilePath != "" {
		// 先頭の１行を取り出し
		var firstLine = vio.inputLines[0]

		// 繰り上がり
		vio.inputLines = vio.inputLines[1:len(vio.inputLines)]

		return firstLine
	}

	return vio.scanner.Text()
}

func (vio *VirtualIO) WriterFlush() {
	virtualIo.writer.Flush()
}

// EOF [O11o__11o2o_1o0]
```

## Step [O11o__11o2o_2o0] バーチャルIO作成 - debugger/virtual_io_fmt.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📂 debugger
  	│   ├── 📄 test.input.txt
👉 	│   ├── 📄 virtual_io_fmt.go
  	│   └── 📄 virtual_io.go
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	└── 📄 logger.go
  	├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```go
// BOF [O11o__11o2o_2o0]

package debugger

import "fmt"

// 文字列出力
func (vio *VirtualIO) Printf(format string, a ...interface{}) {
	fmt.Fprintf(vio.writer, format, a...)
}

// EOF [O11o__11o2o_2o0]
```

## Step [O11o__11o2o0] ファイル作成 - debugger/main.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📂 debugger
👉  │   ├── 📄 main.go
  	│   ├── 📄 test.input.txt
 	│   ├── 📄 virtual_io_fmt.go
  	│   └── 📄 virtual_io.go
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	└── 📄 logger.go
  	├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```go
// BOF [O11o__11o2o0]

package debugger

import (
	"strconv"
)

// グローバル変数として、バーチャルIOを１つ新規作成
// アプリケーションの中では 標準入出力は これを使うようにする
var virtualIo = NewVirtualIO()

func main() {
	// この関数を抜けるときに、バーチャルIOの出力バッファーをフラッシュする
	defer virtualIo.WriterFlush()

	// 入力を読取る
	if virtualIo.ScannerScan() {
		var text = virtualIo.ScannerText()
		var i, err = strconv.Atoi(text)
		if err != nil {
			panic(err)
		}

		virtualIo.Printf("%d is ok\n", i) // 出力
	}
}

// EOF [O11o__11o2o0]
```

## Step [O11o__11o3o0] ファイル作成 - debugger/main_test.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📂 debugger
👉  │   ├── 📄 main_test.go
  	│   ├── 📄 main.go
  	│   ├── 📄 test.input.txt
 	│   ├── 📄 virtual_io_fmt.go
  	│   └── 📄 virtual_io.go
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	└── 📄 logger.go
  	├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```go
// BOF [O11o__11o3o0]

package debugger

import (
	"testing"
)

func TestMain(t *testing.T) {
	virtualIo.ReplaceInputToFileLines("./test.input.txt")
	main()
}

// EOF [O11o__11o3o0]
```

## Step [O11o__11o4o0] モジュール作成 - デバッガー

👇 📂 debugger をカレントディレクトリーとする  

```shell
cd debugger
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go mod init github.com/muzudho/kifuwarabe-uec14/debugger
#           --------------------------------------------
#           1
# 1. モジュール名。この部分はあなたのレポジトリに合わせて変えてほしい
```

👇 以下のファイルが自動生成される  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📂 debugger
👉  │   ├── 📄 go.mod
  	│   ├── 📄 main_test.go
  	│   ├── 📄 main.go
  	│   ├── 📄 test.input.txt
 	│   ├── 📄 virtual_io_fmt.go
  	│   └── 📄 virtual_io.go
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	└── 📄 logger.go
  	├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```plaintext
module github.com/muzudho/kifuwarabe-uec14/debugger

go 1.19
```

## Step [O11o__11o5o0] ワークスペースズモードへ登録

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go work use .
go mod tidy
```

👇 カレントディレクトリーは戻しておいてほしい  

```shell
cd ..
```

# Step [O11o__11o6o_1o0] デバッグ用コマンドライン引数作成

## Step [O11o__11o6o_2o0] データ作成 - debug.input.txt

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📂 debugger
  	│   ├── 📄 go.mod
  	│   ├── 📄 main_test.go
  	│   ├── 📄 main.go
  	│   ├── 📄 test.input.txt
 	│   ├── 📄 virtual_io_fmt.go
  	│   └── 📄 virtual_io.go
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	└── 📄 logger.go
    ├── 📄 .gitignore
👉	├── 📄 debug.input.txt
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```plaintext
quit
```

## Step [O11o__11o6o0] ファイル編集 - main.go

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📂 debugger
  	│   ├── 📄 go.mod
  	│   ├── 📄 main_test.go
  	│   ├── 📄 main.go
  	│   ├── 📄 test.input.txt
 	│   ├── 📄 virtual_io_fmt.go
  	│   └── 📄 virtual_io.go
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	└── 📄 logger.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
👉 	└── 📄 main.go
```

```go
// func main() {
	// // [O11o__10o_5o0] 思考エンジン設定ファイル
	// var (
	// 	pEngineFilePath = flag.String("f", "engine.toml", "engine config file path")
		// [O11o__11o6o0] デバッグ用
		pIsDebug = flag.Bool("d", false, "for debug")
	// )
	// ...略...

	// この下に初期設定を追加していく
	// ---------------------------
	// ...略...

	// [O11o__11o6o0] デバッグ用
	if *pIsDebug {
		virtualIo.ReplaceInputToFileLines("./debug.input.txt")
	}
	// ...略...

	// この上に初期設定を追加していく
	// ---------------------------
	// ...略...
```

## ~~Step [O11o__11o6o1o0]~~

Removed.  

## Step [O11o__11o6o2o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
go run . -d=true
```

# Step [O11o__12o0] インタープリター 作成

## Step [O11o_1o0] コマンド実装 - ファイル編集 - main.go

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📂 debugger
  	│   ├── 📄 go.mod
  	│   ├── 📄 main_test.go
  	│   ├── 📄 main.go
  	│   ├── 📄 test.input.txt
 	│   ├── 📄 virtual_io_fmt.go
  	│   └── 📄 virtual_io.go
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	└── 📄 logger.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
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

// [O11o_1o0] グローバル変数として、バーチャルIOを１つ新規作成
// アプリケーションの中では 標準入出力は これを使うようにする
var virtualIo = dbg.NewVirtualIO()

// ...略...

// func main() {
	// ...略...

	if name == "hello" { // [O9o0]
		// ...略...
	} else {

		// * 消す
		// fmt.Println("go run . {programName}")

		// * 追加
		// [O11o_1o0] コンソール等からの文字列入力
		for virtualIo.ScannerScan() {
			var command = virtualIo.ScannerText()
			logg.C.Infof("# %s", command)             // 人間向けの出力
			logg.J.Infow("input", "command", command) // コンピューター向けの出力

			// [O11o_1o0]
			var tokens = strings.Split(command, " ")
			switch tokens[0] {

			// この下にコマンドを挟んでいく
			// -------------------------

			case "quit": // [O11o_1o0]
				// os.Exit(0)
				return

			// この上にコマンドを挟んでいく
			// -------------------------

			default: // [O11o_1o0]
				logg.C.Infof("? unknown_command command:'%s'\n", tokens[0])
				logg.J.Infow("? unknown_command", "command", tokens[0])
			}
		}
	}


// ...略...
// }
```

## Step [O11o_2o0] 動作確認

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
[2022-09-12 00:36:47]   # abc
[2022-09-12 00:36:47]   ? unknown_command command:'abc'
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
