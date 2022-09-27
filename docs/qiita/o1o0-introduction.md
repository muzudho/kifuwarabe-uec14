# 目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O1o0] 導入

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# 本文

## Step [O1o0] はじめに

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

## Step [O2o0] ソースの置き場所（ローカル）

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

## Step [O3o0] Visual Studio Code を使う

がんばって、 `Visual Studio Code` を使えるようにしておいてほしい  

📖 [Visual Studio Code](https://code.visualstudio.com/)  

## Step [O4o0] Goエクステンションをインストールする

`Visual Studio Code` の `Extensions` から、 `Go` をインストールしておいてほしい  

## Step [O5o0] マルチ ワークスペース

モジュールを複数作れるよう、  
複数のワークスペースを作れるように対応しておく    

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go work init
```

👇 以下のファイルが自動生成される  

```plaintext
  	📂 kifuwarabe-uec14
👉  └── 📄 go.work
```

```go
go 1.19

```

## Step [O6o0] 設定 - .gitignore ファイル

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

# [O6o0]
go.work


# この上に kifuwarabe-uec14 でリポジトリにコミットしないものを追加する
# ---------------------------------------------------------------
# ...略...
```

## Step [O7o0] モジュール作成

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

## Step [O8o0] ワークスペースズへ登録

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

# Step [O9o0] エントリーポイント作成 - ハローワールド

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
    ├── 📄 .gitignore
    ├── 📄 go.mod
  	├── 📄 go.work
👉 	└── 📄 main.go
```

```go
// BOF [O9o0]

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

	if name == "hello" { // [O9o0]
		fmt.Println("Hello, World!")

		// この下に分岐を挟んでいく
		// ---------------------

		// この上に分岐を挟んでいく
		// ---------------------

	} else {
		fmt.Println("go run . {programName}")
	}
}

// EOF [O9o0]
```

## Step [O10o0] tidy

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
go mod tidy
```

## Step [O10o1o0] 実行

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go run . hello
```

Output:  

```plaintext
Hello, World!
```
