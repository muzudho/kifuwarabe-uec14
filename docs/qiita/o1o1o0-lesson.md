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

# Step [O1o1o0g2o0] ソースの置き場所（ローカル）

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

# Step [O1o1o0g3o0] Visual Studio Code を使う

がんばって、 `Visual Studio Code` を使えるようにしておいてほしい  

📖 [Visual Studio Code](https://code.visualstudio.com/)  

# Step [O1o1o0g4o0] Goエクステンションをインストールする

`Visual Studio Code` の `Extensions` から、 `Go` をインストールしておいてほしい  

# Step [O1o1o0g5o0] マルチ ワークスペース

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

# Step [O1o1o0g6o0] 設定 - .gitignore ファイル

👇 以下の既存ファイルを編集（無ければ新規作成）してほしい  

```plaintext
  	📂 kifuwarabe-uec14
👉  ├── 📄 .gitignore
  	└── 📄 go.work
```

例えば冒頭に追加  

```plaintext
# * ここから、以下を追加
# (^q^)

go.work
# * ここまで


# ...略...
```

# Step [O1o1o0g7o0] モジュール作成

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

# Step [O1o1o0g8o0] ワークスペースズへ登録

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

# Step [O1o1o0g9o0] エントリーポイント作成

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

	if name == "hello" { // [O1o1o0g9o0]
		fmt.Println("Hello, World!")

		// この上に分岐を挟んでいく

	} else {
		fmt.Println("go run . {programName}")
	}
}

// EOF [O1o1o0g9o0]
```

# Step [O1o1o0g10o0] tidy

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
go mod tidy
```

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

// MemoryWidth - 枠付きの横幅
const MemoryWidth int = 19 + 2

// MemoryHeight - 枠付きの縦幅
const MemoryHeight int = 19 + 2

// MemoryArea - 枠付きの面積
const MemoryArea int = MemoryWidth * MemoryHeight

// Board - 盤
type Board struct {
	// 交点
	nodes []Stone
}

// NewBoard - 新規作成
func NewBoard() *Board {
	var b = new(Board)

	b.nodes = make([]Stone, MemoryArea)

	// 枠を設定する
	// 上辺、下辺を引く
	{
		var y = 0
		var y2 = MemoryHeight - 1
		for x := 0; x < MemoryWidth; x++ {
			var i = (y * MemoryWidth) + x
			b.nodes[i] = Wall

			i = (y2 * MemoryWidth) + x
			b.nodes[i] = Wall
		}
	}
	// 左辺、右辺を引く
	{
		var x = 0
		var x2 = MemoryWidth - 1
		for y := 1; y < MemoryHeight-1; y++ {
			var i = (y * MemoryWidth) + x
			b.nodes[i] = Wall

			i = (y * MemoryWidth) + x2
			b.nodes[i] = Wall
		}
	}

	return b
}

// ForeachLikeText - 枠を含めた各セル
func (b *Board) ForeachLikeText(setStone func(Stone), doNewline func()) {
	for y := 0; y < MemoryHeight; y++ {
		if y != 0 {
			doNewline()
		}

		for x := 0; x < MemoryWidth; x++ {
			var i = (y * MemoryWidth) + x
			var stone = b.nodes[i]
			setStone(stone)
		}
	}
}

// EOF [O1o1o0g12o0]
```

# Step [O1o1o0g13o0] 盤表示 作成

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


	// * 以下を追加
	} else if name == "board_test" { // [O1o1o0g13o0]
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
		fmt.Print("\n. '''")

		// この上に分岐を挟んでいく


// ...略...
```

# Step [O1o1o0g14o0] 実行

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run . board_test
```

Output:  

```plaintext
= board_test'''
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

この出力書式は 私の方法であって、公式大会のものではないことに注意されたい  

📖 [思考エンジンの思考ログ仕様（きふわらべ2022年以降）](https://qiita.com/muzudho1/items/ceb6130cf558cd373dd7)  

# 参考にした記事

## Go言語

### 列挙型

📖 [How to make Go print enum fields as string?](https://stackoverflow.com/questions/41480543/how-to-make-go-print-enum-fields-as-string)  

.

