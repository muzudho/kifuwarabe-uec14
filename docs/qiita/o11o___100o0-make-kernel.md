目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O11o___100o0] カーネル作成 ～ Step [O11o__10o_1o0] 思考エンジン設定ファイル

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# Step [O11o___100o0] カーネル作成

👇 ここでは、カーネルは以下の意味を指す  

* 思考エンジンのプログラムのうち、おおまかに言って **ゲームの知識（ドメイン）以外の部分**
* １つのカーネルは、１つの対局に対応する

## Step [O11o___100o1p0] フォルダー作成

👇 以下のフォルダーを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
👉	├── 📂 kernel
    ├── 📄 .gitignore
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

## Step [O11o___100o2p0] カレントディレクトリーを移動

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd kernel
```

## Step [O11o___100o3o_1o0] ワークスペースズへ登録

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go work use .
```

👇 以下のファイルが自動で編集されている  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
    ├── 📄 .gitignore
    ├── 📄 go.mod
👉 	├── 📄 go.work
 	└── 📄 main.go
```

```go
// ...略...


// * 以下の行が自動削除
// use .
// * 以下が自動追加
use (
	.
	./kernel
)
```

## Step [O11o___100o3o0] Goモジュールの作成

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go mod init github.com/muzudho/kifuwarabe-uec14/kernel
```

Output:  

```shell
go: creating new go.mod: module github.com/muzudho/kifuwarabe-uec14/kernel
```

👇 以下のファイルが自動生成された  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
👉	│	└── 📄 go.mod
    ├── 📄 .gitignore
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```go
module github.com/muzudho/kifuwarabe-uec14/kernel

go 1.19
```

## Step [O11o___100o4p0] カレントディレクトリーを戻す

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd ..
```

# Step [O11o__10o_1o0] 思考エンジン設定ファイル

## Step [O11o__10o_2o0] ファイル作成 - engine.toml

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	└── 📄 go.mod
    ├── 📄 .gitignore
👉  ├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```toml
# BOF [O11o__10o_2o0]
# 思考エンジンのデフォルト値です。 CgfGoBan などの GUI はこのファイルを見ません

# Game - 対局１つ分に相当
[Game]

# Komi - コミ
Komi = 6.5

# BoardSize - 何路盤
BoardSize = 19

# MaxMoves - 最大手数
MaxMoves = 400

# BoardData - 盤面
#
# * '.' - 空点
# * 'x' - 黒石
# * 'o' - 白石
# * '+' - 枠
# * 半角空白は無視します
BoardData = '''
+ + + + + + + + + + + + + + + + + + + + +
+ . . . . . . . . . . . . . . . . . . . +
+ . . . . . . . . . . . . . . . . . . . +
+ . . . . . . . . . . . . . . . . . . . +
+ . . . . . . . . . . . . . . . . . . . +
+ . . . . . . . . . . . . . . . . . . . +
+ . . . . . . . . . . . . . . . . . . . +
+ . . . . . . . . . . . . . . . . . . . +
+ . . . . . . . . . . . . . . . . . . . +
+ . . . . . . . . . . . . . . . . . . . +
+ . . . . . . . . . . . . . . . . . . . +
+ . . . . . . . . . . . . . . . . . . . +
+ . . . . . . . . . . . . . . . . . . . +
+ . . . . . . . . . . . . . . . . . . . +
+ . . . . . . . . . . . . . . . . . . . +
+ . . . . . . . . . . . . . . . . . . . +
+ . . . . . . . . . . . . . . . . . . . +
+ . . . . . . . . . . . . . . . . . . . +
+ . . . . . . . . . . . . . . . . . . . +
+ . . . . . . . . . . . . . . . . . . . +
+ + + + + + + + + + + + + + + + + + + + +
'''

# PlayFirst - 先行。 black か white
PlayFirst = "black"

# ファイルやフォルダーのパスの設定
[Paths]

# PlainTextLog - コンソールのより詳細なログ
PlainTextLog = "kifuwarabe-uec14.log"

# JsonLog - コンピューター向けのログ
JsonLog = "kifuwarabe-uec14-json.log"

# EOF [O11o__10o_2o0]
```

## Step [O11o__10o_3o0] インストール - go-toml

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go get github.com/pelletier/go-toml
go mod tidy
```

Output:  

```plaintext
go: added github.com/pelletier/go-toml v1.9.5
```

## Step [O11o__10o_4o0] ファイル作成 - engine_config.go

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	└── 📄 go.mod
    ├── 📄 .gitignore
👉 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```go
// BOF [O11o__10o_4o0]

package main

import (
	"os"

	"github.com/pelletier/go-toml"
)

// LoadEngineConfig - 思考エンジン設定ファイルを読み込む
func LoadEngineConfig(
	path string,
	onError func(error) *Config) *Config {

	// ファイル読込
	var fileData, err = os.ReadFile(path)
	if err != nil {
		return onError(err)
	}

	// Toml解析
	var binary = []byte(string(fileData))
	var config = new(Config)
	// Go言語の struct に合わせてデータを読み込む
	toml.Unmarshal(binary, config)

	return config
}

// Config - 設定ファイル
type Config struct {
	// Game - 対局
	Game Game
	// Paths - ファイルやフォルダーのパスの設定
	Paths Paths
}

// GetBoardSize - 何路盤か
func (c *Config) GetBoardSize() int {
	return int(c.Game.BoardSize)
}

// GetKomi - コミ
//
// * float 32bit で足りるが、実行速度優先で float 64bit に変換して返す
func (c *Config) GetKomi() float64 {
	return float64(c.Game.Komi)
}

// GetMaxPositionNumber - 最大手数
func (c *Config) GetMaxPositionNumber() int {
	return int(c.Game.MaxMoves)
}

// GetPlayFirst - 先行。 black か white
func (c *Config) GetPlayFirst() string {
	return c.Game.PlayFirst
}

// GetPlainTextLog - PlainTextLog - コンソールのより詳細なログ
func (c *Config) GetPlainTextLog() string {
	return c.Paths.PlainTextLog
}

// GetJsonLog - コンピューター向けのログ
func (c *Config) GetJsonLog() string {
	return c.Paths.JsonLog
}

// Game - 対局
type Game struct {
	// Komi - コミ
	Komi float32

	// BoardSize - 盤の一辺の長さ
	BoardSize int8

	// MaxMoves - 手数の上限
	MaxMoves int16

	// BoardData - 盤面データ
	BoardData string

	// PlayFirst - 先行。 black か white
	PlayFirst string
}

// Paths - ファイルやフォルダーのパスの設定
type Paths struct {
	// PlainTextLog - コンソールのより詳細なログ
	PlainTextLog string

	// JsonLog - コンピューター向けのログ
	JsonLog string
}

// EOF [O11o__10o_4o0]
```

## Step [O11o__10o_5o0] ファイル編集 - main.go

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	└── 📄 go.mod
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
👉 	└── 📄 main.go
```

```go
// ...略...

// func main() {
	// [O11o__10o_5o0] 思考エンジン設定ファイル
	var (
		pEngineFilePath = flag.String("f", "engine.toml", "engine config file path")
	)
	// * この上に追加
	// flag.Parse()
	// var name = flag.Arg(0)
	// この下に初期設定を追加していく
	// ---------------------------
	// * この下に追加
	// [O11o__10o_5o0] 思考エンジン設定ファイル
	var onError = func(err error) *Config {
		// ログファイルには出力できません。ログファイルはまだ読込んでいません

		// 強制終了
		panic(err)
	}
	var engineConfig = LoadEngineConfig(*pEngineFilePath, onError)

// ...略...
```

## Step [O11o__10o_6o0] 実行についての備考

設定ファイルを使った例は、これ以降のステップで示す  

これ以降のステップでは、 📄 `engine.toml` ファイルの名前や内容を任意に編集して、がんばってテストしてほしい。  
テストが終わったらもとに戻してほしい  

👇 思考エンジン設定ファイルを読込むには、以下のようにコマンドに `-f` 引数を付けてほしい  

Input:  

```shell
go run . -f engine.toml
```

* `-f engine.toml` を省略すると、デフォルトで `./engine.toml` ファイルを読みに行く
