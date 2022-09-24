# Go言語、Visual Studio Code でのデバッグ方法を調べる

Windows を使っているとする  

👇 とりあえずググって適当に記事を読む。大筋で、以下の記事の通りやるとできる  

📖 [Goのデバッグ環境 on VSCode](https://future-architect.github.io/articles/20201117/)  

とりあえず真似する  

# Step [O1o0] はじめに

## Step [O1o1o0] エディター

Visual studio code を使っているとする  

## Step [O1o2o0] プログラム言語

Go 言語をインストール済みとする  

## Step [O1o3o0] Go エクステンションを入れる

📖 [Go for Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=golang.Go)  

## Step [O1o4o0] 関連するツールをインストールする

`[Ctrl] + [Shift] + [P]` をキーにする  

`Go: Install/Update Tools` を選ぶ  

出てくるツールを全部選んでインストールする  

Output:  

```plaintext
All tools successfully installed. You are ready to Go. :)
```

このとき、 `dlv` もインストールしたものとする  

# Step [O2o0] デバッガーを使用する

## Step [O2o1o0] 設定ファイルの作成

VSCodeのツールバーの　虫みたいなアイコンをクリック  

`create a launch.json file` をクリック  

VSCode の上の方にフォルダー名か出てくるので、とりあえず ソースを置いているフォルダーを選ぶ  

👇 以下のファイルが自動生成される

```plaintext
    📂 .vscode
    └── 📄 launch.json
```

```json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}"
        }
    ]
}
```

## Step [O2o2o0] デバッグの実行

ブレークポイントを置く  

エントリーポイント（main関数）が書いている `*.go` ファイルを開く  

`F5` キーを押すとデバッグが開始される  

## Step [O3o0] 標準入力の差し替え

dlv では、以下のコードを実行できない  

```go
		var scanner = bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
            // ...略...
        }
```

👇 以下参照。 `delveは標準入力を受け付けられない` という致命的な不便がある  

📖 [golangのAtCoder向けデバック方法(VSCode)](https://qiita.com/tasmas/items/d2d5a8c95fa48e415702)  

解決方法としては、  
Go言語のテストの中で、標準入力 `os.Stdin` の中身を、テキストファイルから向けられているストリームに差し替え、  
テストを、デバッガー を使って実行する  
いうもののようだ  

しかし **入力のタイミングが１回しかない** 気がする。改造が要りそうだ  

## Step [O3o1o_1o0] ワークスペースズ モード使用

Go 1.18 からある ワークスペースズモードを使う。
説明は以前書いたから省略  

📖 [Go [O1o1o0] 目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜練習編＞](https://qiita.com/muzudho1/items/cea62be01f7418bbf150)  

## Step [O3o1o0] ファイル作成 - debugger/main.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
    📂
    └── 📂 debugger
👉      └── 📄 main.go
```

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// VirtualIO - 入出力を模擬したもの
type VirtualIO struct {
	scanner *bufio.Scanner
	writer  *bufio.Writer
}

// 新規作成
func newVirtualIO() *VirtualIO {
	// 実体をメモリ上に占有させ、そのアドレスを返す
	return &VirtualIO{
		scanner: bufio.NewScanner(os.Stdin),
		writer:  bufio.NewWriter(os.Stdout),
	}
}

// 次の文字列入力を読取る
func (io *VirtualIO) nextLine() string {
	io.scanner.Scan()
	return io.scanner.Text()
}

// 次の整数入力を読取る
func (io *VirtualIO) nextInt() int {
	i, e := strconv.Atoi(io.nextLine())
	if e != nil {
		panic(e)
	}
	return i
}

// 文字列出力
func (io *VirtualIO) printLn(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}

// 新規作成
var virtualIo = newVirtualIO()

func main() {
	virtualIo.scanner.Split(bufio.ScanWords)      // switch to separating by space
	virtualIo.scanner.Buffer([]byte{}, 100000007) // switch to read large size input
	defer virtualIo.writer.Flush()

	N := virtualIo.nextInt()
	hoge := fmt.Sprintf("%d is ok", N) // なんらかの処理
	virtualIo.printLn(hoge)            // 出力
}
```

## Step [O3o1o_2o0] モジュール作成

📂 debugger をカレントディレクトリーとする  

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
    📂
    └── 📂 debugger
👉      ├── 📄 go.mod
        └── 📄 main.go
```

```plaintext
module github.com/muzudho/kifuwarabe-uec14/debugger

go 1.19
```

## Step [O3o1o_2o0] ワークスペースズモードに登録

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:

```shell
go mod tidy
go work use .
```

## Step [O3o1o_3o0] 入力データ作成 - debugger/test.txt ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
    📂
    └── 📂 debugger
        ├── 📄 go.mod
        ├── 📄 main_test.mod
        ├── 📄 main.go
👉      └── 📄 test.txt
```

```go
10
```

## Step [O3o1o0] プログラム作成 - debugger/main_test.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
    📂
    └── 📂 debugger
        ├── 📄 go.mod
👉      ├── 📄 main_test.mod
        ├── 📄 main.go
        └── 📄 test.txt
```

```go
package main

import (
	"bufio"
	"os"
	"testing"
)

func TestAnswer(t *testing.T) {
	inbuf := readFile("./test.txt")
	stubStdin(inbuf, func() {
		main()
	})
}

func readFile(fileName string) string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// Stubs Stdin in 'fn'
// See also: 📖 [golangのAtCoder向けデバック方法(VSCode)](https://qiita.com/tasmas/items/d2d5a8c95fa48e415702)
//
// Examples
// --------
// inbuf := "入力されたつもりの文字列。テキストファイルから読み込んでくる"
//
//	stubStdin(inbuf, func() {
//	    main()
//	})
//
// Parameters
// ----------
// textToWrite - 書き込みたい文字列
func stubStdin(textToWrite string, fn func()) {
	// これより、ラムダ計算の専門用語で η（イータ）簡約 と呼ばれることと同じ考え方を利用する。
	// Input ストリームと使い勝手が同等になるよう、 Read モードと Write モードのファイル（メモリ上に存在する）を取得
	inr, inw, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	// Input ストリームに書き込んでいるつもりで、 Write モードのファイルに書き込む
	_, _ = inw.Write([]byte(textToWrite))
	// 書込みをフラッシュして終わる
	inw.Close()

	// Input ストリームから読込んでいるつもりで、 Read モードのファイルを `os.Stdin` と差し替える
	os.Stdin = inr
	// このスキャナーは、標準入力をスキャンしているように見えて、メモリ上に存在するファイルをスキャンしている
	virtualIo.scanner = bufio.NewScanner(os.Stdin)

	// あとは ふつうに処理を行う
	fn()

	// TODO `os.Stdin` を元に戻さなくていいのか？ fn() が main() プログラムと同等で、あとは終了するるだけなら 良いとはいえるが
}
```

## Step [O3o2o_1o0] 実行

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:

```shell
go run .
11
```

Output:  

```plaintext
11 is ok
```

## Step [O3o2o0] テスト実行

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:

```shell
go test
```

Output:  

```plaintext
10 is ok
PASS
ok      github.com/muzudho/kifuwarabe-uec14/debugger    0.206s
```

👇 カレントディレクトリーは元に戻してほしい  

```shell
cd ..
```

# 参考にした記事

## デバッグ環境

📖 [Goのデバッグ環境 on VSCode](https://future-architect.github.io/articles/20201117/)  

## デバッグと標準入力

📖 [scanner.Scan() hangs in GoLand debugger](https://stackoverflow.com/questions/53461228/scanner-scan-hangs-in-goland-debugger)  
📖 [golangのAtCoder向けデバック方法(VSCode)](https://qiita.com/tasmas/items/d2d5a8c95fa48e415702)  

EOF
