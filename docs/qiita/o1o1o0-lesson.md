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

# Step [O1o1o0g11o___100o0] カーネル作成

👇 ここでは、カーネルは以下の意味を指す  

* 思考エンジンのプログラムのうち、おおまかに言って **ゲームの知識（ドメイン）以外の部分**
* １つのカーネルは、１つの対局に対応する

## Step [O1o1o0g11o___100o1p0] フォルダー作成

👇 以下のフォルダーを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
👉	├── 📂 kernel
    ├── 📄 .gitignore
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

## Step [O1o1o0g11o___100o2p0] カレントディレクトリーを移動

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd kernel
```

## Step [O1o1o0g11o___100o3o_1o0] ワークスペースズへ登録

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

## Step [O1o1o0g11o___100o3o0] Goモジュールの作成

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

## Step [O1o1o0g11o___100o4p0] カレントディレクトリーを戻す

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd ..
```

# Step [O1o1o0g11o__10o_1o0] 思考エンジン設定ファイル

## Step [O1o1o0g11o__10o_2o0] ファイル作成 - engine.toml

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
# BOF [O1o1o0g11o__10o_2o0]
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
# * '+' - 壁
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

# EOF [O1o1o0g11o__10o_2o0]
```

## Step [O1o1o0g11o__10o_3o0] インストール - go-toml

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

## Step [O1o1o0g11o__10o_4o0] ファイル作成 - engine_config.go

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
// BOF [O1o1o0g11o__10o_4o0]

package main

import (
	"os"

	"github.com/pelletier/go-toml"
)

// LoadEngineConfig - 思考エンジン設定ファイルを読み込む
func LoadEngineConfig(
	path string,
	onError func(error) Config) Config {

	// ファイル読込
	var fileData, err = os.ReadFile(path)
	if err != nil {
		return onError(err)
	}

	// Toml解析
	var binary = []byte(string(fileData))
	var config = Config{}
	// Go言語の struct に合わせてデータを読み込む
	toml.Unmarshal(binary, &config)

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

// GetMaxMovesNum - 最大手数
func (c *Config) GetMaxMovesNum() int {
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

// EOF [O1o1o0g11o__10o_4o0]
```

## Step [O1o1o0g11o__10o_5o0] ファイル編集 - main.go

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
	// [O1o1o0g11o__10o_5o0] 思考エンジン設定ファイル
	var (
		pEngineFilePath = flag.String("f", "engine.toml", "engine config file path")
	)
	// * この上に追加
	// flag.Parse()
	// var name = flag.Arg(0)
	// この下に初期設定を追加していく
	// ---------------------------
	// * この下に追加
	// [O1o1o0g11o__10o_5o0] 思考エンジン設定ファイル
	var onError = func(err error) {
		// ログファイルには出力できません。ログファイルはまだ読込んでいません
		panic(err)
	}
	var engineConfig = LoadEngineConfig(*pEngineFilePath, onError)

// ...略...
```

## Step [O1o1o0g11o__10o_6o0] 実行についての備考

設定ファイルを使った例は、これ以降のステップで示す  

これ以降のステップでは、 📄 `engine.toml` ファイルの名前や内容を任意に編集して、がんばってテストしてほしい。  
テストが終わったらもとに戻してほしい  

👇 思考エンジン設定ファイルを読込むには、以下のようにコマンドに `-f` 引数を付けてほしい  

Input:  

```shell
go run . -f engine.toml
```

* `-f engine.toml` を省略すると、デフォルトで `./engine.toml` ファイルを読みに行く

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
	├── 📂 kernel
	│	└── 📄 go.mod
👉	├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
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
	├── 📂 kernel
	│	├── 📄 go.mod
👉 	│	└── 📄 logger.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```go
// BOF [O1o1o0g11o__10o2o0]

package kernel

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	// C is sugared logger for Console
	C *zap.SugaredLogger
	// J is sugared logger as JSON
	J *zap.SugaredLogger
}

func NewSugaredLoggerForGame(plainTextLogFile *os.File, jsonLogFile *os.File) *Logger {
	var slog = new(Logger) // Sugared LOGger
	slog.C = createSugaredLoggerForConsole(plainTextLogFile)
	slog.J = createSugaredLoggerAsJson(jsonLogFile)
	return slog
}

// ロガーを作成します，コンソール形式
func createSugaredLoggerForConsole(plainTextLogFile *os.File) *zap.SugaredLogger {
	// 設定，コンソール用
	var configC = zapcore.EncoderConfig{
		MessageKey: "message",

		// LevelKey:    "level",
		// EncodeLevel: zapcore.CapitalLevelEncoder,

		TimeKey:    "time",
		EncodeTime: encodeTimeSimpleInJapan, // 簡略化したタイムスタンプ

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
			zapcore.AddSync(plainTextLogFile),       // 出力先はファイル
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

// 簡略化したタイムスタンプ
// 📖 [golang zap v1.0.0 でログの日付をJSTで表示する方法](https://qiita.com/fuku2014/items/c6501c187c8161336485)
func encodeTimeSimpleInJapan(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	// // JST形式
	// const layout = "2006-01-02T15:04:05+09:00"
	// jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	// enc.AppendString(t.In(jst).Format(layout))

	// 簡略化したタイムスタンプ
	const layout = "[2006-01-02 15:04:05]"
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	enc.AppendString(t.In(jst).Format(layout))
}

// EOF [O1o1o0g11o__10o2o0]
```

## Step [O1o1o0g11o__10o3o_1o0] リモートリポジトリにプッシュ

がんばって git などを使い、 `github.com/muzudho/kifuwarabe-uec14/kernel` モジュールの各パッケージのソースを  
リモートリポジトリにプッシュしてほしい  

## Step [O1o1o0g11o__10o3o0] ファイル編集

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	└── 📄 logger.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	├── 📄 logger.go
👉 	└── 📄 main.go
```

```go
// ...略...

import (
	// ...略...
	"github.com/muzudho/kifuwarabe-uec14/kernel" // * あとでリポジトリからダウンロードする
)

func main() {
	// ...略...
	// この下に初期設定を追加していく
	// ---------------------------

	// [O1o1o0g11o__10o3o0] ログファイル
	var plainTextLogFile, _ = os.OpenFile(engineConfig.GetPlainTextLog(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer plainTextLogFile.Close() // ログファイル使用済み時にファイルを閉じる
	// ログファイル
	var jsonLogFile, _ = os.OpenFile(engineConfig.GetJsonLog(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer jsonLogFile.Close() // ログファイル使用済み時にファイルを閉じる
	// カスタマイズしたロガーを使うなら
	var logg = kernel.NewSugaredLoggerForGame(plainTextLogFile, jsonLogFile) // customized LOGGer

	// この上に初期設定を追加していく
	// ---------------------------

	if name == "hello" { // [O1o1o0g9o0]
		// ...略...
		// この下に分岐を挟んでいく
		// ---------------------
		// ...略...


	} else if name == "welcome" { // [O1o1o0g11o__10o0]
		logg.C.Infof("Welcome! name:'%s' weight:%.1f x:%d", "nihon taro", 92.6, 3)
		logg.J.Infow("Welcome!",
			"name", "nihon taro", "weight", 92.6, "x", 3)


		// ...略...
		// この上に分岐を挟んでいく
		// ---------------------
		// ...略...
	}
```

## Step [O1o1o0g11o__10o4o0] tidy

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go get -u
go mod tidy
```

Output:  

```plaintext
C:\Users\むずでょ\Documents\GitHub\kifuwarabe-uec14>go mod tidy
go: finding module for package github.com/muzudho/kifuwarabe-uec14/kernel
go: downloading github.com/muzudho/kifuwarabe-uec14/kernel v0.0.0-20220911105808-5a17a869516a
go: found github.com/muzudho/kifuwarabe-uec14/kernel in github.com/muzudho/kifuwarabe-uec14/kernel v0.0.0-20220911105808-5a17a869516a
go: downloading github.com/stretchr/testify v1.8.0
go: downloading github.com/benbjohnson/clock v1.1.0
go: downloading go.uber.org/goleak v1.1.11
go: downloading gopkg.in/yaml.v3 v3.0.1
```

👇 自分のパッケージはローカルPCにダウンロードされている  

Example: 📂 `C:\Users\むずでょ\go\pkg\mod\github.com\muzudho`  

## Step [O1o1o0g11o__10o5o0] 実行

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go run . welcome
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
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	└── 📄 logger.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
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

# Step [O1o1o0g11o__10o0] インタープリター 作成

## Step [O1o1o0g11o_1o0] コマンド実装 - ファイル編集 - main.go

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
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
			logg.C.Infof("# %s", command)             // 人間向けの出力
			logg.J.Infow("input", "command", command) // コンピューター向けの出力

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


// ...略...
```

## Step [O1o1o0g11o_2o0] 動作確認

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

# Step [O1o1o0g11o_3o0] カーネルのインタープリター

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
👉 	│	├── 📄 kernel.go
 	│	└── 📄 logger.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```go
// BOF [O1o1o0g11o_3o0]

package kernel

import "strings"

type Kernel struct {
	// Board - 盤
	Board *Board
}

// NewKernel - カーネルの新規作成
func NewKernel() *Kernel {
	var k = new(Kernel)
	k.Board = NewBoard()
	return k
}

// Execute - 実行
//
// Returns
// -------
// isHandled : bool
// 処理済なら真
func (k *Kernel) Execute(command string, logg *Logger) bool {

	var tokens = strings.Split(command, " ")
	switch tokens[0] {

	// この下にコマンドを挟んでいく
	// -------------------------

	// この上にコマンドを挟んでいく
	// -------------------------

	default:
	}

	return false
}

// EOF [O1o1o0g11o_3o0]
```

## Step [O1o1o0g11o_3o1o0] 外側のインタープリターから呼び出す

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	└── 📄 logger.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
👉 	└── 📄 main.go
```

👇 がんばって挿入してほしい  

```go
// ...略...

import (
	// ...略...

	"github.com/muzudho/kifuwarabe-uec14/kernel"
)
// ...略...

		// [O1o1o0g12o__11o_4o0] 棋譜の初期化に利用
		var onUnknownTurn = func() kernel.Stone {
			var errMsg = fmt.Sprintf("? unexpected play_first:%s", engineConfig.GetPlayFirst())
			logg.C.Info(errMsg)
			logg.J.Infow("error", "play_first", engineConfig.GetPlayFirst())
			panic(errMsg)
		}

		// [O1o1o0g11o_3o0]
		var kernel1 = kernel.NewKernel()
		// 設定ファイルの内容をカーネルへ反映
		kernel1.Board.Init(engineConfig.GetBoardSize(), engineConfig.GetBoardSize())

		/*
		...以下略...
		// [O1o1o0g11o_1o0] コンソール等からの文字列入力
		var scanner = bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			var command = scanner.Text()
			logg.C.Infof("# %s", command)             // 人間向けの出力
			logg.J.Infow("input", "command", command) // コンピューター向けの出力
		...以上略...
		*/

			// * これを追加する
			// [O1o1o0g11o_3o0]
			var isHandled = kernel1.Execute(command, logg)
			if isHandled {
				continue
			}

			/*
			...以下略...
			// [O1o1o0g11o_1o0]
			var tokens = strings.Split(command, " ")
			switch tokens[0] {
			*/
```

* カーネルと、自作の部分で コマンドが被ったなら、カーネルの方を優先する
  * これにより カーネルのアップデートにより 自作のコマンドが避けられることから、アップデート時は動作テストしてほしい

# Step [O1o1o0g11o_4o0] 石の色定義

`石` を定義していないが、先に `石の色` を定義する  

石の色の組み合わせを定義する。  
石の色の組み合わせは以下の４通りある。これらの集合を `Color` と名付けることにする   

* `Color_None` - 隣接する連は１つもない
* `Color_Black` - 黒石の連とだけ隣接する
* `Color_White` - 白石の連とだけ隣接する
* `Color_Mixed` - 黒石と白石の連の両方に隣接する

## Step [O1o1o0g11o_4o1o0] ファイル作成 - color.go

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
👉	│	├── 📄 color.mod
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	└── 📄 logger.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```go
// BOF [O1o1o0g11o_4o1o0]

package kernel

import "fmt"

type Color uint

const (
	Color_None Color = iota
	Color_Black
	Color_White
	Color_Mixed
)

// String - 文字列化
func (c Color) String() string {
	switch c {
	case Color_None:
		return ""
	case Color_Black:
		return "x"
	case Color_White:
		return "o"
	case Color_Mixed:
		return "xo"
	default:
		panic(fmt.Sprintf("unexpected color:%d", int(c)))
	}
}

// GetAdded - 色の加算
func (c1 Color) GetAdded(c2 Color) Color {
	switch c1 {
	case Color_None:
		return c2
	case Color_Black:
		switch c2 {
		case Color_None:
			return Color_Black
		case Color_Black:
			return Color_Black
		case Color_White:
			return Color_Mixed
		case Color_Mixed:
			return Color_Mixed
		default:
			panic(fmt.Sprintf("unexpected my_color:%s adds_color:%s", c1, c2))
		}
	case Color_White:
		switch c2 {
		case Color_None:
			return Color_White
		case Color_Black:
			return Color_Mixed
		case Color_White:
			return Color_White
		case Color_Mixed:
			return Color_Mixed
		default:
			panic(fmt.Sprintf("unexpected my_color:%s adds_color:%s", c1, c2))
		}
	case Color_Mixed:
		return Color_Mixed
	default:
		panic(fmt.Sprintf("unexpected my_color:%s adds_color:%s", c1, c2))
	}
}

// GetOpponent - 色の反転
func (c Color) GetOpponent() Color {
	switch c {
	case Color_None:
		return Color_Mixed
	case Color_Black:
		return Color_White
	case Color_White:
		return Color_Black
	case Color_Mixed:
		return Color_None
	default:
		panic(fmt.Sprintf("unexpected color:%d", int(c)))
	}
}

// EOF [O1o1o0g11o_4o1o0]
```

# Step [O1o1o0g11o_4o2o0] 連の定義

`石` を定義していないが、先に `連` （れん）を定義する。  
`連` とは何かの説明は、ここでは省く  

### Step [O1o1o0g11o_4o2o1o0] ファイル作成 - ren.go

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 color.mod
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
👉 	│	└── 📄 ren.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```go
// BOF [O1o1o0g11o_4o2o1o0]

package kernel

// Ren - 連，れん
type Ren struct {
	// Color - 色
	Color Color
	// AdjacentColor - 隣接する石の色
	AdjacentColor Color
	// LibertyArea - 呼吸点の面積
	LibertyArea int
	// Elements - 要素の石の位置
	Elements []Point
}

// GetArea - 面積
func (r *Ren) GetArea() int {
	return len(r.Elements)
}

// EOF [O1o1o0g11o_4o2o1o0]
```

# Step [O1o1o0g11o_5o0] 石定義

## Step [O1o1o0g11o0] ファイル作成 - stone.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
👉 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```go
// BOF [O1o1o0g11o0]

package kernel

import "fmt"

type Stone uint

const (
	Space Stone = iota
	Black
	White
	Wall
)

// GetStoneFromName - 文字列の名前を与えると、Stone値を返します
//
// Returns
// -------
// isOk : bool
// stone : Stone
func GetStoneFromName(stoneName string, getDefaultStone func() (bool, Stone)) (bool, Stone) {
	switch stoneName {
	case "space":
		return true, Space
	case "black":
		return true, Black
	case "white":
		return true, White
	case "wall":
		return true, Wall
	default:
		return getDefaultStone()
	}
}

// GetStoneOrDefaultFromTurn - black または white を与えると、Stone値を返します
//
// Returns
// -------
// stone : Stone
func GetStoneOrDefaultFromTurn(stoneName string, getDefaultStone func() Stone) Stone {
	switch stoneName {
	case "black":
		return Black
	case "white":
		return White
	default:
		return getDefaultStone()
	}
}

// GetStoneFromChar - １文字与えると、Stone値を返します
//
// Returns
// -------
// isOk : bool
// stone : Stone
func GetStoneFromChar(stoneChar string, getDefaultStone func() (bool, Stone)) (bool, Stone) {
	switch stoneChar {
	case ".":
		return true, Space
	case "x":
		return true, Black
	case "o":
		return true, White
	case "+":
		return true, Wall
	default:
		return getDefaultStone()
	}
}

// String - 文字列化
func (s Stone) String() string {
	switch s {
	case Space:
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

// GetColor - 色の取得
func (s Stone) GetColor() Color {
	switch s {
	case Space:
		return Color_None
	case Black:
		return Color_Black
	case White:
		return Color_White
	case Wall:
		return Color_None
	default:
		panic(fmt.Sprintf("%d", int(s)))
	}
}

// EOF [O1o1o0g11o0]
```

## Step [O1o1o0g11o1o0] リモートリポジトリにプッシュ

がんばって git などを使い、 `github.com/muzudho/kifuwarabe-uec14/kernel` モジュールの各パッケージのソースを  
リモートリポジトリにプッシュしてほしい  

## Step [O1o1o0g11o2o0] カレントディレクトリーを移動

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd kernel
```

## Step [O1o1o0g11o3o0] tidy

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go get -u
go mod tidy
```

## Step [O1o1o0g11o4o0] カレントディレクトリーを戻す

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd ..
```

## Step [O1o1o0g11o5o0] tidy

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go mod tidy
```

# Step [O1o1o0g12o__10o0] 点定義、またはその盤座標符号定義

盤を作る前に、これから盤座標符号を定義する  

## Step [O1o1o0g12o__10o1o0] ファイル作成 - point.go ファイル

👇 以下の既存ファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
👉  │	├── 📄 point.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
  	└── 📄 main.go
```

```go
// BOF [O1o1o0g12o__10o1o0]

package kernel

import (
	"fmt"
	"strconv"
)

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

// GetFileFromX - GetXFromFile の逆関数
func GetFileFromX(x int) string {
	// ABCDEFGHI
	// 012345678
	if 7 < x {
		// 'I' を飛ばす
		x++
	}
	// 筋
	return fmt.Sprintf("%c", 'A'+x)
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

// GetRankFromY - GetYFromRank の逆関数
//
// Parameters
// ----------
// y : int
//
//	0 .. 98
//
// Returns
// -------
// rank : string
//
//	"1" .. "99"
func GetRankFromY(y int) string {
	return strconv.Itoa(y + 1)
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

// EOF [O1o1o0g12o__10o1o0]
```

## Step [O1o1o0g12o__10o2o0] コマンド実装 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
👉 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
  	└── 📄 main.go
```

👇 がんばって、 Execute メソッドに挿入してほしい  

```go
// ...略...


	// この下にコマンドを挟んでいく
	// -------------------------

	// ...略...

	// * アルファベット順になる位置に、以下のケース文を挿入
	case "test_file": // [O1o1o0g12o__10o2o0]
		// Example: "test_file A"
		var file = GetFileFromCode(tokens[1])
		logg.C.Infof("= %s\n", file)
		logg.J.Infow("output", "file", file)
		return true

	case "test_rank": // [O1o1o0g12o__10o2o0]
		// Example: "test_rank 13"
		var rank = GetRankFromCode(tokens[1])
		logg.C.Infof("= %s\n", rank)
		logg.J.Infow("output", "rank", rank)
		return true

	case "test_x": // [O1o1o0g12o__10o2o0]
		// Example: "test_x 18"
		var x, err = strconv.Atoi(tokens[1])
		if err != nil {
			logg.C.Infof("? unexpected x:%s\n", tokens[1])
			logg.J.Infow("error", "x", tokens[1])
			return true
		}
		var file = GetFileFromX(x)
		logg.C.Infof("= %s\n", file)
		logg.J.Infow("output", "file", file)
		return true

	case "test_y": // [O1o1o0g12o__10o2o0]
		// Example: "test_y 18"
		var y, err = strconv.Atoi(tokens[1])
		if err != nil {
			logg.C.Infof("? unexpected y:%s\n", tokens[1])
			logg.J.Infow("error", "y", tokens[1])
			return true
		}
		var rank = GetRankFromY(y)
		logg.C.Infof("= %s\n", rank)
		logg.J.Infow("output", "rank", rank)
		return true

	// この上にコマンドを挟んでいく
	// -------------------------


// ...略...
```

## Step [O1o1o0g12o__10o3o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
test_file A1
```

Output > Console:  

```plaintext
[2022-09-11 23:27:52]   # test_file A1
[2022-09-11 23:27:52]   = A
```

Output > Log > PlainText:  

```plaintext
2022-09-11T23:27:52.547+0900	# test_file A1
2022-09-11T23:27:52.583+0900	= A
```

Output > Log > JSON:  

```json
{"level":"info","ts":"2022-09-11T23:27:52.583+0900","caller":"kifuwarabe-uec14/main.go:61","msg":"input","command":"test_file A1"}
{"level":"info","ts":"2022-09-11T23:27:52.584+0900","caller":"kernel/kernel.go:73","msg":"output","file":"A"}
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
test_rank A1
```

Output > Console:  

```plaintext
[2022-09-11 23:31:11]   # test_rank A1
[2022-09-11 23:31:11]   = 1
```

Output > Log > PlainText:  

```plaintext
2022-09-11T23:31:11.020+0900	# test_rank A1
2022-09-11T23:31:11.020+0900	= 1
```

Output > Log > JSON:  

```json
{"level":"info","ts":"2022-09-11T23:31:11.020+0900","caller":"kifuwarabe-uec14/main.go:61","msg":"input","command":"test_rank A1"}
{"level":"info","ts":"2022-09-11T23:31:11.021+0900","caller":"kernel/kernel.go:80","msg":"output","rank":"1"}
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
test_coord A1
```

Output > Console:  

```plaintext
[2022-09-11 23:32:42]   # test_coord A1
[2022-09-11 23:32:42]   = 22
```

Output > Log > PlainText:  

```plaintext
2022-09-11T23:32:42.228+0900	# test_coord A1
2022-09-11T23:32:42.229+0900	= 22
```

Output > Log > JSON:  

```json
{"level":"info","ts":"2022-09-11T23:32:42.229+0900","caller":"kifuwarabe-uec14/main.go:61","msg":"input","command":"test_coord A1"}
{"level":"info","ts":"2022-09-11T23:32:42.229+0900","caller":"kernel/kernel.go:66","msg":"output","point":22}
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
test_x 18
```

* x は 0 から始まるので、 19列目は 18

Output > Console:  

```plaintext
[2022-09-13 23:53:40]   # test_x 18
[2022-09-13 23:53:40]   = T
```

Output > Log > PlainText:  

```plaintext
2022-09-13T23:53:40.906+0900	# test_x 18
2022-09-13T23:53:40.943+0900	= T
```

Output > Log > JSON:  

```json
{"level":"info","ts":"2022-09-13T23:53:40.943+0900","caller":"kifuwarabe-uec14/main.go:76","msg":"input","command":"test_x 18"}
{"level":"info","ts":"2022-09-13T23:53:40.943+0900","caller":"kernel/kernel.go:115","msg":"output","file":"T"}
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
test_y 18
```

* y は 0 から始まるので、 19列目は 18

Output > Console:  

```plaintext
[2022-09-13 23:58:42]   # test_y 18
[2022-09-13 23:58:42]   = 19
```

Output > Log > PlainText:  

```plaintext
2022-09-13T23:58:42.739+0900	# test_y 18
2022-09-13T23:58:42.781+0900	= 19
```

Output > Log > JSON:  

```json
{"level":"info","ts":"2022-09-13T23:58:42.781+0900","caller":"kifuwarabe-uec14/main.go:76","msg":"input","command":"test_y 18"}
{"level":"info","ts":"2022-09-13T23:58:42.782+0900","caller":"kernel/kernel.go:128","msg":"output","rank":"19"}
```

# Step [O1o1o0g12o__11o_1o0] 棋譜定義

## Step [O1o1o0g12o__11o_2o0] ファイル作成 - record.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
	│	├── 📄 point.go
👉	│	├── 📄 record.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
  	└── 📄 main.go
```

```go
// BOF [O1o1o0g12o__11o_2o0]

package kernel

// Record - 棋譜
type Record struct {
	// 先行
	playFirst Stone

	// 着手点
	points []Point
}

// NewRecord - 棋譜の新規作成
func NewRecord(maxMoves int, playFirst Stone) *Record {
	var r = new(Record)
	r.playFirst = playFirst
	r.points = make([]Point, maxMoves)
	return r
}

// Push - 末尾に追加
func (r *Record) Push(placePlay Point) {
	r.points = append(r.points, placePlay)
}

// Push - 末尾を削除
func (r *Record) Pop(placePlay Point) Point {
	var lastIndex = len(r.points) - 1
	var tail = r.points[lastIndex]
	r.points = r.points[:lastIndex]
	return tail
}

// EOF [O1o1o0g12o__11o_2o0]
```

## Step [O1o1o0g12o__11o_3o0] ファイル編集 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
👉 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
	│	├── 📄 point.go
	│	├── 📄 record.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
  	└── 📄 main.go
```

```go
// type Kernel struct {
	// ...略...

	// Record - [O1o1o0g12o__11o_3o0] 棋譜
	Record Record

// }

// NewKernel - カーネルの新規作成
// func NewKernel(
	// [O1o1o0g12o__11o_2o0] 棋譜の初期化
	maxMoves int, playFirst Stone//) *Kernel {
	// ...略...

	// * 以下を追加
	// [O1o1o0g12o__11o_2o0] 棋譜の初期化
	k.Record = *NewRecord(maxMoves, playFirst)

	// ...略...
	// return k
// }
```

## Step [O1o1o0g12o__11o_4o0] ファイル編集 - main.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
	│	├── 📄 point.go
	│	├── 📄 record.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
👉 	└── 📄 main.go
```

```go
		// [O1o1o0g11o_3o0]
		//var kernel1 = kernel.NewKernel(
			// [O1o1o0g12o__11o_4o0] 棋譜の初期化
			engineConfig.GetMaxMovesNum(),
			kernel.GetStoneOrDefaultFromTurn(engineConfig.GetPlayFirst(), onUnknownTurn)//)
```

## Step [O1o1o0g12o__11o_5o0] コマンド実装 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
👉 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
	│	├── 📄 point.go
	│	├── 📄 record.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

👇 がんばって、 Execute メソッドに挿入してほしい  

```go
	// ...略...
	// この下にコマンドを挟んでいく
	// -------------------------
	// ...略...

	// * アルファベット順になる位置に、以下のケース文を挿入
	case "record": // [O1o1o0g12o__11o_5o0]
		// Example: "record"
		var sb strings.Builder
		for i, point := range k.Record.points {
			var ordinals = i+1	// 基数を序数へ変換
			sb.WriteString(fmt.Sprintf("%d.%s ", ordinals, k.Board.GetCodeFromPoint(point)))
		}
		var text = sb.String()
		text = text[:len(text)-1]
		logg.C.Infof("= record:'%s'\n", text)
		logg.J.Infow("ok", "record", text)
		return true

	// ...略...
	// この上にコマンドを挟んでいく
	// -------------------------
	// ...略...
```

## Step [O1o1o0g12o__11o_6o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
play black A1
play black B2
play black C3
play black D4
play black E5
record
```

Output > Console:  

```plaintext
[2022-09-17 18:04:45]   # record
[2022-09-17 18:04:45]   = record:'0.@0 1.@0 2.@0 3.@0 4.@0 5.@0 6.@0 7.@0 8.@0 9.@0 10.@0 11.@0 12.@0 13.@0 14.@0 15.@0 16.@0 17.@0 18.@0 
19.@0 20.@0 21.@0 22.@0 23.@0 24.@0 25.@0 26.@0 27.@0 28.@0 29.@0 30.@0 31.@0 32.@0 33.@0 34.@0 35.@0 36.@0 37.@0 38.@0 39.@0 40.@0 41.@0 
42.@0 43.@0 44.@0 45.@0 46.@0 47.@0 48.@0 49.@0 50.@0 51.@0 52.@0 53.@0 54.@0 55.@0 56.@0 57.@0 58.@0 59.@0 60.@0 61.@0 62.@0 63.@0 64.@0 
65.@0 66.@0 67.@0 68.@0 69.@0 70.@0 71.@0 72.@0 73.@0 74.@0 75.@0 76.@0 77.@0 78.@0 79.@0 80.@0 81.@0 82.@0 83.@0 84.@0 85.@0 86.@0 87.@0 
88.@0 89.@0 90.@0 91.@0 92.@0 93.@0 94.@0 95.@0 96.@0 97.@0 98.@0 99.@0 100.@0 101.@0 102.@0 103.@0 104.@0 105.@0 106.@0 107.@0 108.@0 109.@0 110.@0 111.@0 112.@0 113.@0 114.@0 115.@0 116.@0 117.@0 118.@0 119.@0 120.@0 121.@0 122.@0 123.@0 124.@0 125.@0 126.@0 127.@0 128.@0 129.@0 130.@0 131.@0 132.@0 133.@0 134.@0 135.@0 136.@0 137.@0 138.@0 139.@0 140.@0 141.@0 142.@0 143.@0 144.@0 145.@0 146.@0 147.@0 148.@0 149.@0 150.@0 151.@0 152.@0 153.@0 154.@0 155.@0 156.@0 157.@0 158.@0 159.@0 160.@0 161.@0 162.@0 163.@0 164.@0 165.@0 166.@0 167.@0 168.@0 169.@0 170.@0 171.@0 172.@0 173.@0 174.@0 175.@0 176.@0 177.@0 178.@0 179.@0 180.@0 181.@0 182.@0 183.@0 184.@0 185.@0 186.@0 187.@0 188.@0 189.@0 190.@0 191.@0 192.@0 193.@0 194.@0 195.@0 196.@0 197.@0 198.@0 199.@0 200.@0 201.@0 202.@0 203.@0 204.@0 205.@0 206.@0 207.@0 
208.@0 209.@0 210.@0 211.@0 212.@0 213.@0 214.@0 215.@0 216.@0 217.@0 218.@0 219.@0 220.@0 221.@0 222.@0 223.@0 224.@0 225.@0 226.@0 227.@0 228.@0 229.@0 230.@0 231.@0 232.@0 233.@0 234.@0 235.@0 236.@0 237.@0 238.@0 239.@0 240.@0 241.@0 242.@0 243.@0 244.@0 245.@0 246.@0 247.@0 248.@0 249.@0 250.@0 251.@0 252.@0 253.@0 254.@0 255.@0 256.@0 257.@0 258.@0 259.@0 260.@0 261.@0 262.@0 263.@0 264.@0 265.@0 266.@0 267.@0 268.@0 269.@0 270.@0 271.@0 272.@0 273.@0 274.@0 275.@0 276.@0 277.@0 278.@0 279.@0 280.@0 281.@0 282.@0 283.@0 284.@0 285.@0 286.@0 287.@0 288.@0 289.@0 290.@0 291.@0 292.@0 293.@0 294.@0 295.@0 296.@0 297.@0 298.@0 299.@0 300.@0 301.@0 302.@0 303.@0 304.@0 305.@0 306.@0 307.@0 308.@0 309.@0 310.@0 311.@0 312.@0 313.@0 314.@0 315.@0 316.@0 317.@0 318.@0 319.@0 320.@0 321.@0 322.@0 323.@0 324.@0 325.@0 326.@0 327.@0 328.@0 329.@0 330.@0 331.@0 332.@0 333.@0 334.@0 335.@0 336.@0 337.@0 338.@0 339.@0 340.@0 341.@0 342.@0 343.@0 344.@0 345.@0 
346.@0 347.@0 348.@0 349.@0 350.@0 351.@0 352.@0 353.@0 354.@0 355.@0 356.@0 357.@0 358.@0 359.@0 360.@0 361.@0 362.@0 363.@0 364.@0 365.@0 366.@0 367.@0 368.@0 369.@0 370.@0 371.@0 372.@0 373.@0 374.@0 375.@0 376.@0 377.@0 378.@0 379.@0 380.@0 381.@0 382.@0 383.@0 384.@0 385.@0 386.@0 387.@0 388.@0 389.@0 390.@0 391.@0 392.@0 393.@0 394.@0 395.@0 396.@0 397.@0 398.@0 399.@0 400.A1 401.B2 402.C3 403.D4 404.E5' 
```

# Step [O1o1o0g12o__11o0] 盤定義（土台）

これから盤を作っていく前に、土台を作る  

## Step [O1o1o0g12o__11o1o0] ファイル作成 - board.go

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
👉  │	├── 📄 board.go
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```go
// BOF [O1o1o0g12o__11o1o0]

package kernel

// Board - 盤
type Board struct {
	// 枠付きの横幅
	memoryWidth int

	// 枠付きの縦幅
	memoryHeight int

	// 交点
	//
	// * 英語で交点は node かも知れないが、表計算でよく使われる cell の方を使う
	cells []Stone
}

// NewBoard - 新規作成
func NewBoard() *Board {
	var b = new(Board)

	// 盤のサイズ指定と、盤面の初期化
	b.resize(19, 19)

	return b
}

// GetMemoryWidth - 枠の厚みを含んだ横幅
func (b *Board) GetMemoryWidth() int {
	return b.memoryWidth
}

// GetMemoryHeight - 枠の厚みを含んだ縦幅
func (b *Board) GetMemoryHeight() int {
	return b.memoryHeight
}

// GetWidth - 枠の厚みを含まない横幅
func (b *Board) GetWidth() int {
	return b.memoryWidth - 2
}

// GetHeight - 枠の厚みを含まない縦幅
func (b *Board) GetHeight() int {
	return b.memoryHeight - 2
}

// GetStoneAt - 指定座標の石を取得
func (b *Board) GetStoneAt(i Point) Stone {
	return b.cells[i]
}

// SetStoneAt - 指定座標の石を設定
func (b *Board) SetStoneAt(i Point, s Stone) {
	b.cells[i] = s
}

// GetColorAt - 指定座標の石の色を取得
func (b *Board) GetColorAt(i Point) Color {
	return b.cells[i].GetColor()
}

// GetPointFromXy - 座標変換 (x,y) → Point
//
// Parameters
// ----------
// x : int
//	筋番号。 Example: 19路盤なら0～18
// y : int
//	段番号。 Example: 19路盤なら0～18
//
// Returns
// -------
// point : Point
//  配列インデックス。 Example: 2,3 なら 65
func (b *Board) GetPointFromXy(x int, y int) Point {
	return Point(y*b.memoryWidth + x)
}

// GetXyFromPoint - `GetPointFromXy` の逆関数
func (b *Board) GetXyFromPoint(point Point) (int, int) {
	var p = int(point)
	return p % b.memoryWidth, p / b.memoryWidth
}

// サイズ変更
func (b *Board) resize(width int, height int) {
	b.memoryWidth = width + 2
	b.memoryHeight = height + 2
	b.cells = make([]Stone, b.getMemoryArea())
}

// 枠付き盤の面積
func (b *Board) getMemoryArea() int {
	return b.memoryWidth * b.memoryHeight
}

// EOF [O1o1o0g12o__11o1o0]
```

## Step [O1o1o0g12o__11o2o0] コマンド実装 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
  	│	├── 📄 board.go
	│	├── 📄 go.mod
👉 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

👇 がんばって、 Execute メソッドに挿入してほしい  

```go
	// ...略...
	// この下にコマンドを挟んでいく
	// -------------------------
	// ...略...

	// * アルファベット順になる位置に、以下のケース文を挿入
	case "test_get_point_from_xy": // [O1o1o0g12o__11o2o0]
		// Example: "test_get_point_from_xy 2 3"
		var x, errX = strconv.Atoi(tokens[1])
		if errX != nil {
			logg.C.Infof("? unexpected x:%s\n", tokens[1])
			logg.J.Infow("error", "x", tokens[1], "err", errX)
			return true
		}
		var y, errY = strconv.Atoi(tokens[2])
		if errY != nil {
			logg.C.Infof("? unexpected y:%s\n", tokens[2])
			logg.J.Infow("error", "y", tokens[2], "err", errY)
			return true
		}

		var point = k.Board.GetPointFromXy(x, y)
		logg.C.Infof("= %d\n", point)
		logg.J.Infow("output", "point", point)
		return true

	// この上にコマンドを挟んでいく
	// -------------------------
	// ...略...
```

## Step [O1o1o0g12o__11o3o0] 動作確認

19路盤とする  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
test_get_point_from_xy 2 3
```

Output > Console:  

```plaintext
[2022-09-14 22:37:42]   # test_get_point_from_xy 2 3
[2022-09-14 22:37:42]   = 65
```

Output > Log > PlainText:  

```plaintext
2022-09-14T22:37:42.600+0900	# test_get_point_from_xy 2 3
2022-09-14T22:37:42.637+0900	= 65
```

Output > Log > JSON:  

```json
{"level":"info","ts":"2022-09-14T22:37:42.637+0900","caller":"kifuwarabe-uec14/main.go:76","msg":"input","command":"test_get_point_from_xy 2 3"}
{"level":"info","ts":"2022-09-14T22:37:42.638+0900","caller":"kernel/kernel.go:119","msg":"output","point":65}
```

# Step [O1o1o0g12o_1o0] 盤定義（盤面）

## Step [O1o1o0g12o0] ファイル作成 - board_area.go

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
👉  │	├── 📄 board_area.go
  	│	├── 📄 board.go
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```go
// BOF [O1o1o0g12o0]

package kernel

// Init - 盤面初期化
func (b *Board) Init(width int, height int) {
	// 盤面のサイズが異なるなら、盤面を作り直す
	if b.memoryWidth != width || b.memoryHeight != height {
		b.resize(width, height)
	}

	// 枠の上辺、下辺を引く
	{
		var y = 0
		var y2 = b.memoryHeight - 1
		for x := 0; x < b.memoryWidth; x++ {
			var i = b.GetPointFromXy(x, y)
			b.cells[i] = Wall

			i = b.GetPointFromXy(x, y2)
			b.cells[i] = Wall
		}
	}
	// 枠の左辺、右辺を引く
	{
		var x = 0
		var x2 = b.memoryWidth - 1
		for y := 1; y < b.memoryHeight-1; y++ {
			var i = b.GetPointFromXy(x, y)
			b.cells[i] = Wall

			i = b.GetPointFromXy(x2, y)
			b.cells[i] = Wall
		}
	}
	// 枠の内側を空点で埋める
	{
		var height = b.GetHeight()
		var width = b.GetWidth()
		for y := 1; y < height; y++ {
			for x := 1; x < width; x++ {
				var i = b.GetPointFromXy(x, y)
				b.cells[i] = Space
			}
		}
	}
}

// ForeachLikeText - 枠を含めた各セル
func (b *Board) ForeachLikeText(setStone func(Stone), doNewline func()) {
	for y := 0; y < b.memoryHeight; y++ {
		if y != 0 {
			doNewline()
		}

		for x := 0; x < b.memoryWidth; x++ {
			var i = b.GetPointFromXy(x, y)
			var stone = b.cells[i]
			setStone(stone)
		}
	}
}

// ForeachNeumannNeighborhood - [O1o1o0g13o__10o0] 隣接する４方向の定義
func (b *Board) ForeachNeumannNeighborhood(here Point, setAdjacentPoint func(int, Point)) {
	// 東、北、西、南
	for dir := 0; dir < 4; dir++ {
		var adjacentP = here + Point(b.Direction[dir]) // 隣接する交点

		// 範囲外チェック
		if adjacentP < 1 || b.getMemoryArea() <= int(adjacentP) {
			continue
		}

		setAdjacentPoint(dir, adjacentP)
	}
}

// EOF [O1o1o0g12o0]
```

## Step [O1o1o0g12o1o0] カレントディレクトリーを移動

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd kernel
```

## Step [O1o1o0g12o2o0] tidy

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go get -u
go mod tidy
```

## Step [O1o1o0g12o3o0] カレントディレクトリーを戻す

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd ..
```

## Step [O1o1o0g12o4o0] リモートリポジトリにプッシュ

がんばって git などを使い、 `github.com/muzudho/kifuwarabe-uec14/kernel` モジュールの各パッケージのソースを  
リモートリポジトリにプッシュしてほしい  

## Step [O1o1o0g12o5o0] tidy

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go get -u
```

Output:  

```plaintext
go: downloading github.com/muzudho/kifuwarabe-uec14/kernel v0.0.0-20220911114704-f68bc2cc5ece
go: upgraded github.com/muzudho/kifuwarabe-uec14/kernel v0.0.0-20220911112135-f237cc5d1832 => v0.0.0-20220911114704-f68bc2cc5ece
```

Input:  

```
go mod tidy
```

# Step [O1o1o0g13o_1o0] 盤表示 - board コマンド

## Step [O1o1o0g13o0] コマンド実装 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
  	│	├── 📄 board_area.go
  	│	├── 📄 board.go
	│	├── 📄 go.mod
👉 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
	└── 📄 main.go
```

👇 がんばって、 Execute メソッドに挿入してほしい  

```go
	// ...略...
	// この下にコマンドを挟んでいく
	// -------------------------

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
			k.Board.ForeachLikeText(setStone, doNewline)
			sb.WriteString("\n. '''\n")
			logg.C.Info(sb.String())
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
			k.Board.ForeachLikeText(setStone, doNewline)
			logg.J.Infow("output", "board", sb.String())
		}
		return true

	// この上にコマンドを挟んでいく
	// -------------------------
	// ...略...
```

## Step [O1o1o0g14o0] 動作確認

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
	├── 📂 kernel
  	│	├── 📄 board_area.go
  	│	├── 📄 board.go
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
👉 	├── 📄 kifuwarabe-uec14-json.log
👉 	├── 📄 kifuwarabe-uec14.log
  	└── 📄 main.go
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

# Step [O1o1o0g15o__10o0] 盤サイズの変更 - resize コマンド

## Step [O1o1o0g15o__11o0] コマンド実装 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
  	│	├── 📄 board_area.go
  	│	├── 📄 board.go
	│	├── 📄 go.mod
👉 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
	└── 📄 main.go
```

👇 がんばって、 Execute メソッドに挿入してほしい  

```go
// ...略...

	// この下にコマンドを挟んでいく
	// -------------------------

	case "boardsize": // [O1o1o0g15o__11o0]
		// Example: `boardsize 19`
		var sideLength, err = strconv.Atoi(tokens[1])

		if err != nil {
			logg.C.Infof("? unexpected sideLength:%s\n", tokens[1])
			logg.J.Infow("error", "sideLength", tokens[1])
			return true
		}

		k.Board.Init(sideLength, sideLength)
		logg.C.Info("=\n")
		logg.J.Infow("ok")

		return true

	// ...略...

	// この上にコマンドを挟んでいく
	// -------------------------

// ...略...
```

## Step [O1o1o0g15o__12o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
boardsize 9
```

Output > Console:  

```plaintext
[2022-09-12 20:44:12]   # boardsize 9
[2022-09-12 20:44:12]   =
```

Output > Log > Plain:  

```plaintext
2022-09-12T20:44:12.860+0900	# boardsize 9
2022-09-12T20:44:12.896+0900	=
```

Output > Log > Json:  

```json
{"level":"info","ts":"2022-09-12T20:44:12.894+0900","caller":"kifuwarabe-uec14/main.go:61","msg":"input","command":"boardsize 9"}
{"level":"info","ts":"2022-09-12T20:44:12.896+0900","caller":"kernel/kernel.go:76","msg":"ok"}
```

Input:  

```shell
board
```

Output:  

```plaintext
[2022-09-12 20:56:20]   # board
[2022-09-12 20:56:20]   = board:'''
. +++++++++++
. +.........+
. +.........+
. +.........+
. +.........+
. +.........+
. +.........+
. +.........+
. +.........+
. +.........+
. +++++++++++
. '''
```

# ~~Step [O1o1o0g15o__13o0]~~

Removed  

## ~~Step [O1o1o0g15o__13o1o0]~~

Moved to `O1o1o0g11o__10o_2o0`  

## ~~Step [O1o1o0g15o__13o2o_1o0]~~

Moved to `[O1o1o0g11o__10o_3o0]`  

## ~~Step [O1o1o0g15o__13o2o_2o0]~~

Moved to `[O1o1o0g11o__10o_4o0]`  

## ~~Step [O1o1o0g15o__13o2o_3o0]~~

Merged to `[O1o1o0g11o_3o0]`  

## ~~Step [O1o1o0g15o__13o2o_4o0]~~

Moved to `[O1o1o0g11o__10o_6o0]`  

# Step [O1o1o0g15o__14o_1o0] データファイル作成 - data/board1.txt ファイル

あとで使うファイルを先に作成する  

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
👉 	│	└── 📄 board1.txt
	├── 📂 kernel
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
 	│	├── 📄 set_board.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```plaintext
    ABCDEFGHJKLMNOPQRST
   +++++++++++++++++++++
 1 +...................+
 2 +.xxx...............+
 3 +.x.x...............+
 4 +.xxx...............+
 5 +...................+
 6 +...................+
 7 +...................+
 8 +...................+
 9 +...................+
10 +...................+
11 +...................+
12 +...................+
13 +...................+
14 +...................+
15 +...................+
16 +...................+
17 +...................+
18 +...................+
19 +...................+
   +++++++++++++++++++++
```

# Step [O1o1o0g15o__14o0] 初期盤面を設定する - set_board コマンド

## Step [O1o1o0g15o__14o1o0] ファイル作成 - set_board.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
👉 	│	├── 📄 set_board.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```go
// BOF [O1o1o0g15o__14o1o0]

package kernel

import (
	"os"
	"strings"
)

// DoSetBoard - 盤面を設定する
//
// コマンドラインの複数行入力は難しいので、ファイルから取ることにする
// * `command` - Example: `set_board file data/board1.txt`
// ........................--------- ---- ---------------
// ........................0         1    2
func (k *Kernel) DoSetBoard(command string, logg *Logger) {
	var tokens = strings.Split(command, " ")

	if tokens[1] == "file" {
		var filePath = tokens[2]

		var fileData, err = os.ReadFile(filePath)
		if err != nil {
			logg.C.Infof("? unexpected file:%s\n", filePath)
			logg.J.Infow("error", "file", filePath)
			return
		}

		var getDefaultStone = func() (bool, Stone) {
			return false, Space
		}

		var size = k.Board.getMemoryArea()
		var i Point = 0
		for _, c := range string(fileData) {
			var str = string([]rune{c})
			var isOk, stone = GetStoneFromChar(str, getDefaultStone)

			if isOk {
				if size <= int(i) {
					// 配列サイズ超過
					logg.C.Infof("? board out of bounds i:%d size:%d\n", i, size)
					logg.J.Infow("error board out of bounds", "i", i, "size", size)
					return
				}

				k.Board.SetStoneAt(i, stone)
				i++
			}
		}

		// サイズが足りていないなら、エラー
		if int(i) != size {
			logg.C.Infof("? not enough size i:%d size:%d\n", i, size)
			logg.J.Infow("error not enough size", "i", i, "size", size)
			return
		}
	}

}

// EOF [O1o1o0g15o__14o1o0]
```

## Step [O1o1o0g15o__14o2o0] コマンド実装 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
👉 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
 	│	├── 📄 set_board.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

👇 がんばって、 Execute メソッドに挿入してほしい  

```go
	// ...略...
	// この下にコマンドを挟んでいく
	// -------------------------
	// ...略...

	case "set_board": // [O1o1o0g15o__14o2o0]
		// Example: `set_board file data/board1.txt`
		k.DoSetBoard(command, logg)
		return true

	// ...略...
	// この上にコマンドを挟んでいく
	// -------------------------
	// ...略...
```

## Step [O1o1o0g15o__14o3o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
set_board file data/board1.txt
board
```

出力は略  

# Step [O1o1o0g15o_1o0] 座標の定義

## ~~Step [O1o1o0g15o0]~~

Moved to `[O1o1o0g12o__10o1o0]`  

## Step [O1o1o0g16o0] ファイル作成 - board_coord.go ファイル

👇 以下の既存ファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
👉  │	├── 📄 board_coord.go
  	│	├── 📄 board_area.go
  	│	├── 📄 board.go
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
  	└── 📄 main.go
```

```go
// BOF [O1o1o0g16o0]

package kernel

// GetPointFromCode - "A7" や "J13" といった符号を Point へ変換します
//
// * `code` - 座標の符号。 Example: "A7" や "J13"
func (b *Board) GetPointFromCode(code string) Point {
	// 枠の厚み
	var left_wall = 1
	var top_wall = 1
	return b.GetPointFromXy(
		GetXFromFile(GetFileFromCode(code))+left_wall,
		GetYFromRank(GetRankFromCode(code))+top_wall)
}

// GetCodeFromPoint - `GetPointFromCode` の逆関数
func (b *Board) GetCodeFromPoint(point Point) string {
	// 枠の厚み
	var left_wall = 1
	var top_wall = 1
	var x, y = b.GetXyFromPoint(point)
	var file = GetFileFromX(x - left_wall)
	var rank = GetRankFromY(y - top_wall)
	return fmt.Sprintf("%s%s", file, rank)
}

// EOF [O1o1o0g16o0]
```

### Step [O1o1o0g16o1o0] コマンド実装 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 board_coord.go
  	│	├── 📄 board_area.go
  	│	├── 📄 board.go
	│	├── 📄 go.mod
👉 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
  	└── 📄 main.go
```

👇 がんばって、 Execute メソッドに挿入してほしい  

```go
	// ...略...
	// この下にコマンドを挟んでいく
	// -------------------------
	// ...略...

	// * アルファベット順になる位置に、以下のケース文を挿入
	case "test_get_point_from_code": // [O1o1o0g16o1o0]
		// Example: "test_get_point_from_code A1"
		var point = k.Board.GetPointFromCode(tokens[1])
		var code = k.Board.GetCodeFromPoint(point)
		logg.C.Infof("= %d %s", point, code)
		logg.J.Infow("ok", "point", point, "code", code)
		return true

	// ...略...
	// この上にコマンドを挟んでいく
	// -------------------------
	// ...略...
```

## Step [O1o1o0g16o2o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
test_get_point_from_code A1
```

Output > Console:  

```plaintext
[2022-09-17 17:18:46]   # test_get_point_from_code A1
[2022-09-17 17:18:46]   = 22 A1
```

## ~~Step [O1o1o0g17o0]~~

Moved to `[O1o1o0g12o__10o2o0]`  

## ~~Step [O1o1o0g18o0]~~

Moved to `[O1o1o0g12o__10o3o0]`  

# Step [O1o1o0g19o_1o0] 石を打つ - play コマンド

## Step [O1o1o0g19o0] ファイル作成 - play.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 board.go
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
👉 	│	├── 📄 play.go
 	│	├── 📄 point.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```go
// BOF [O1o1o0g19o0]

package kernel

import "strings"

// DoPlay - 打つ
//
// * `command` - Example: `play black A19`
//                         ---- ----- ---
//                         0    1     2
func (k *Kernel) DoPlay(command string, logg *Logger) {
	var tokens = strings.Split(command, " ")
	var stoneName = tokens[1]

	var getDefaultStone = func() (bool, Stone) {
		logg.C.Infof("? unexpected stone:%s\n", stoneName)
		logg.J.Infow("error", "stone", stoneName)
		return false, Space
	}

	var isOk1, stone = GetStoneFromName(stoneName, getDefaultStone)
	if !isOk1 {
		return
	}

	var coord = tokens[2]
	var point = k.Board.GetPointFromCode(coord)

	// [O1o1o0g22o1o2o0] 石（または壁）の上に石を置こうとしました
	var onMasonry = func() bool {
		logg.C.Infof("? masonry my_stone:%s point:%s\n", stone, k.Board.GetCodeFromPoint(point))
		logg.J.Infow("error", "my_stone", stone, "point", k.Board.GetCodeFromPoint(point))
		return false
	}

	var isOk = k.Play(stone, point, logg,
		// [O1o1o0g22o1o2o0] ,onMasonry
		onMasonry)

	if isOk {
		k.Record.Push(point) // 棋譜に追加
		logg.C.Info("=\n")
		logg.J.Infow("ok")
	}
}

// Play - 石を打つ
//
// Returns
// -------
// isOk : bool
//		石を置けたら真、置けなかったら偽
func (k *Kernel) Play(stoneA Stone, pointB Point, logg *Logger,
	// [O1o1o0g22o1o2o0] onMasonry
	onMasonry func() bool) bool {

	// [O1o1o0g22o1o2o0]
	if k.IsMasonryError(stoneA, pointB) {
		return onMasonry()
	}

	// 石を置く
	k.Board.SetStoneAt(pointB, stoneA)

	return true
}

// EOF [O1o1o0g19o0]
```

## Step [O1o1o0g20o0] 実装 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 board.go
	│	├── 📄 go.mod
👉 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
	└── 📄 main.go
```

👇 がんばって、 Execute メソッドに挿入してほしい  

```go
	// ...略...
	// この下にコマンドを挟んでいく
	// -------------------------
	// ...略...

	// * アルファベット順になる位置に、以下のケース文を挿入
	case "play": // [O1o1o0g20o0]
		// Example: `play black A19`
		k.DoPlay(command, logg)
		return true

	// ...略...
	// この上にコマンドを挟んでいく
	// -------------------------
	// ...略...
```

## Step [O1o1o0g21o0] 実行

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
play black B2
```

Output > Console:  

```plaintext
[2022-09-12 21:46:48]   # play black B2
[2022-09-12 21:46:48]   =
```

Output > Log > Plain:  

```plaintext
2022-09-12T21:46:48.701+0900	# play black B2
2022-09-12T21:46:48.739+0900	=
```

Output > Log > Json:  

```json
{"level":"info","ts":"2022-09-12T21:46:48.739+0900","caller":"kifuwarabe-uec14/main.go:61","msg":"input","command":"play black B2"}
{"level":"info","ts":"2022-09-12T21:46:48.739+0900","caller":"kernel/play.go:32","msg":"ok"}
```

Input:  

```shell
board
```

Output > Console:  

```plaintext
[2022-09-12 21:49:36]   # board
[2022-09-12 21:49:36]   = board:'''
. +++++++++++++++++++++
. +...................+
. +.x.................+
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

# Step [O1o1o0g22o_1o0] データファイル作成 - data/board2.txt ファイル

あとで使うファイルを先に作成する  

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	├── 📄 board1.txt
👉 	│	└── 📄 board2.txt
	├── 📂 kernel
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
 	│	├── 📄 set_board.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```plaintext
    ABCDEFGHJKLMNOPQRST
   +++++++++++++++++++++
 1 +...................+
 2 +...x...............+
 3 +..xox..............+
 4 +.xo.ox.............+
 5 +..xox..............+
 6 +...x...............+
 7 +...................+
 8 +...................+
 9 +...................+
10 +...................+
11 +...................+
12 +...................+
13 +...................+
14 +...................+
15 +...................+
16 +...................+
17 +...................+
18 +...................+
19 +...................+
   +++++++++++++++++++++
```

# Step [O1o1o0g22o0] 囲碁の石を打つルールの実装

## Step [O1o1o0g22o1o0] 空点以外のところ（石または壁の上）に石を置くことの禁止 - IsMasonryError関数作成

とりあえず、 `石または壁の上に石を置く行為` に `Masonry` （メイスンリー）という名前を付ける。  
従って この主のエラーは `Masonry error` と呼ぶことにする。  
そのようなエラーであるかどうか判定する関数の名前は `IsMasonryError` と呼ぶことにする  

### Step [O1o1o0g22o1o1o0] ファイル作成 - masonry.go

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📂 play_rule
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
👉 	│	├── 📄 masonry.go
 	│	├── 📄 play.go
 	│	├── 📄 point.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
	└── 📄 main.go
```

```go
// BOF [O1o1o0g22o1o1o0]

package kernel

import "fmt"

func (k *Kernel) IsMasonryError(stone Stone, point Point) bool {
	var target = k.Board.cells[point]

	switch target {
	case Space:
		return false
	case Black, White, Wall:
		return true
	default:
		panic(fmt.Sprintf("unexpected target cell:%s", target))
	}
}

// EOF [O1o1o0g22o1o1o0]
```

### Step [O1o1o0g22o1o2o0] 呼出し

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📂 play_rule
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
👉 	│	├── 📄 play.go
 	│	├── 📄 point.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
	└── 📄 main.go
```

👇 がんばって以下の箇所に挿入してほしい  

```go
// func (k *Kernel) DoPlay(command string, logg *Logger) {
	// ...略...
	// var point = k.Board.GetPointFromCode(tokens[2])

	// * 以下を追加
	// [O1o1o0g22o1o2o0]
	var onMasonry = func() bool {
		logg.C.Infof("? masonry my_stone:%s point:%s\n", stone, point)
		logg.J.Infow("error", "my_stone", stone, "point", point)
		return false
	}

	// var isOk = k.Play(stone, point, logg,
		// * 以下を追加
		// [O1o1o0g22o1o2o0] ,onMasonry
		onMasonry//)
	// ...略...
// }

// func (k *Kernel) Play(stone Stone, point Point, logg *Logger,
	// * 以下を追加
	// [O1o1o0g22o1o2o0] onMasonry
	onMasonry func() bool//) bool {

	// * 以下を追加
	// [O1o1o0g22o1o2o0]
	if k.IsMasonryError(stone, point) {
		return onMasonry()
	}

//	k.Board.cells[point] = stone
//	return true
// }
```

## Step [O1o1o0g22o2o0] 連の認識と、呼吸点のカウント - GetLiberty 関数作成

盤上の座標を指定することで、そこにある `連` の `呼吸点` の数を算出したい  

* `連` - Ren、れん。コンピューター囲碁用語。説明は省略
* `呼吸点` - Liberty、こきゅうてん。コンピューター囲碁用語。説明は省略

👇 呼吸点を数えるために探索すると、一緒に以下のことも行える  

* 連の認識
* 隣接する連の色の取得

このような探索を行う関数に名前を付ける。  
`GetRen` がふさわしいが、慣習を優先して `GetLiberty` と名付けることにする  

### ~~Step [O1o1o0g22o2o1o0]~~

Moved to `[O1o1o0g11o_4o2o1o0]`  

### Step [O1o1o0g22o2o2o0] ファイル作成 - check_board.go

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📂 play_rule
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 board.go
👉 	│	├── 📄 check_board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
 	│	├── 📄 play.go
 	│	├── 📄 point.go
 	│	├── 📄 ren.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
	└── 📄 main.go
```

```go
// BOF [O1o1o0g22o2o2o0]

package kernel

// CheckBoard - チェック盤
type CheckBoard struct {
	// 枠付きの横幅
	memoryWidth int

	// 枠付きの縦幅
	memoryHeight int

	// 交点
	//
	// * 英語で交点は node かも知れないが、表計算でよく使われる cell の方を使う
	cells []bool
}

// NewCheckBoard - 新規作成
//
// * 使用する前に Init 関数を呼び出してほしい
func NewCheckBoard() *CheckBoard {
	var b = new(CheckBoard)
	return b
}

// Init - 盤面初期化
func (b *CheckBoard) Init(width int, height int) {
	// 盤面のサイズが異なるなら、盤面を作り直す
	if b.memoryWidth != width || b.memoryHeight != height {
		b.Resize(width, height)
		return
	}

	// 盤面のクリアー
	for i := 0; i < len(b.cells); i++ {
		b.cells[i] = false
	}
}

// Resize - サイズ変更
func (b *CheckBoard) Resize(width int, height int) {
	b.memoryWidth = width + 2
	b.memoryHeight = height + 2
	b.cells = make([]bool, b.getMemoryArea())
}

// Check - チェックを付けます
func (b *CheckBoard) Check(point Point) {
	b.cells[point] = true
}

// IsChecked - チェックが付いているか？
func (b *CheckBoard) IsChecked(point Point) bool {
	return b.cells[point]
}

// 枠付き盤の面積
func (b *CheckBoard) getMemoryArea() int {
	return b.memoryWidth * b.memoryHeight
}

// EOF [O1o1o0g22o2o2o0]
```

### Step [O1o1o0g22o2o3o_1o0] ファイル編集 - board.go

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📂 play_rule
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
👉 	│	├── 📄 board.go
 	│	├── 📄 check_board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
 	│	├── 📄 play.go
 	│	├── 📄 point.go
 	│	├── 📄 ren.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
	└── 📄 main.go
```

```go
// type Board struct {
// ...略...

	// Direction - ４方向（東、北、西、南）の番地への相対インデックス
	Direction [4]int

// }
// ...略...

// func (b *Board) resize(width int, height int) {
	// ...略...

	// ４方向（東、北、西、南）の番地への相対インデックス
	b.Direction = [4]int{1, -b.GetMemoryWidth(), -1, b.GetMemoryWidth()}

// }
```

### Step [O1o1o0g22o2o3o0] ファイル編集 - kernel.go

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📂 play_rule
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 board.go
 	│	├── 📄 check_board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
👉 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
 	│	├── 📄 play.go
 	│	├── 📄 point.go
 	│	├── 📄 ren.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
	└── 📄 main.go
```

👇 がんばって挿入してほしい  

```go
// ...略...
// type Kernel struct {
//	Board *Board

	// * 以下を追加
	// [O1o1o0g22o2o3o0]
	// CheckBoard - 呼吸点の探索時に使います
	CheckBoard *CheckBoard
	// Ren - 呼吸点の探索時に使います
	Ren *Ren
//}

// func NewKernel() *Kernel {
//	var k = new(Kernel)
//	k.Board = NewBoard()

	// * 以下を追加
	// [O1o1o0g22o2o3o0]
	k.CheckBoard = NewCheckBoard()

//	return k
// }
```

### Step [O1o1o0g22o2o4o0] ファイル作成 - liberty.go

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📂 play_rule
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 board.go
 	│	├── 📄 check_board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	├── 📄 kernel.go
👉 	│	├── 📄 liberty.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
 	│	├── 📄 play.go
 	│	├── 📄 point.go
 	│	├── 📄 ren.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
	└── 📄 main.go
```

```go
// BOF [O1o1o0g22o2o4o0]

package kernel

// GetLiberty - 呼吸点の数え上げ。連の数え上げ
//
// Parameters
// ----------
// * `arbitraryPoint` - 連に含まれる任意の一点
func (k *Kernel) GetLiberty(arbitraryPoint Point) *Ren {
	// チェックボードの初期化
	k.CheckBoard.Init(k.Board.GetWidth(), k.Board.GetHeight())
	// 連の初期化
	k.Ren = new(Ren)
	// 連の色
	k.Ren.Color = k.Board.GetColorAt(arbitraryPoint)
	// 隣接する連の色
	k.Ren.AdjacentColor = Color_None

	k.searchRen(arbitraryPoint)

	return k.Ren
}

// 再帰関数。連の探索
func (k *Kernel) searchRen(here Point) {
	k.CheckBoard.Check(here)
	k.Ren.Elements = append(k.Ren.Elements, here)

	var setAdjacentPoint = func(dir int, adjacentP Point) {
		// 探索済みならスキップ
		if k.CheckBoard.IsChecked(adjacentP) {
			return
		}

		var adjacentS = k.Board.GetStoneAt(adjacentP)
		if adjacentS == Space { // 空点
			k.CheckBoard.Check(adjacentP)
			k.Ren.LibertyArea++
			return
		} else if adjacentS == Wall { // 壁
			return
		}

		var adjacentC = adjacentS.GetColor()
		// 隣接する色、追加
		k.Ren.AdjacentColor = k.Ren.AdjacentColor.GetAdded(adjacentC)

		if adjacentC == k.Ren.Color { // 同色の石
			k.searchRen(adjacentP) // 再帰
		}
	}

	// 隣接する４方向
	k.Board.ForeachNeumannNeighborhood(here, setAdjacentPoint)
}

// EOF [O1o1o0g22o2o4o0]
```

### Step [O1o1o0g22o2o5o0] コマンド実装 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📂 play_rule
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 board.go
 	│	├── 📄 check_board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
👉 	│	├── 📄 kernel.go
 	│	├── 📄 liberty.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
 	│	├── 📄 play.go
 	│	├── 📄 point.go
 	│	├── 📄 ren.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
	└── 📄 main.go
```

👇 がんばって、 Execute メソッドに挿入してほしい  

```go
// ...略...


	// この下にコマンドを挟んでいく
	// -------------------------

	// ...略...

	// * アルファベット順になる位置に、以下のケース文を挿入
	case "test_get_liberty": // [O1o1o0g22o2o5o0]
		// Example: "test_get_liberty B2"
		var point = k.Board.GetPointFromCode(tokens[1])
		var ren = k.GetLiberty(point)
		logg.C.Infof("= ren color:%s area:%d libertyArea:%d adjacentColor:%s\n", ren.Color, ren.GetArea(), ren.LibertyArea, ren.AdjacentColor)
		logg.J.Infow("output ren", "color", ren.Color, "area", ren.GetArea(), "libertyArea", ren.LibertyArea, "adjacentColor", ren.AdjacentColor)
		return true

	// この上にコマンドを挟んでいく
	// -------------------------


// ...略...
```

### Step [O1o1o0g22o2o6o0] 動作確認

19路盤とする  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
play black B2
test_get_liberty B2
```

Output > Console:  

```plaintext
[2022-09-14 23:36:15]   # play black B2
[2022-09-14 23:36:15]   =

[2022-09-14 23:36:21]   # test_get_liberty B2
[2022-09-14 23:36:21]   = ren color:x area:1 libertyArea:4 adjacentColor:
```

Output > Log > PlainText:  

```plaintext
2022-09-14T23:36:15.521+0900	# play black B2
2022-09-14T23:36:15.556+0900	=

2022-09-14T23:36:21.463+0900	# test_get_liberty B2
2022-09-14T23:36:21.464+0900	= ren color:x area:1 libertyArea:4 adjacentColor:
```

Output > Log > JSON:  

```json
{"level":"info","ts":"2022-09-14T23:36:15.556+0900","caller":"kifuwarabe-uec14/main.go:76","msg":"input","command":"play black B2"}
{"level":"info","ts":"2022-09-14T23:36:15.556+0900","caller":"kernel/play.go:43","msg":"ok"}
{"level":"info","ts":"2022-09-14T23:36:21.464+0900","caller":"kifuwarabe-uec14/main.go:76","msg":"input","command":"test_get_liberty B2"}
{"level":"info","ts":"2022-09-14T23:36:21.465+0900","caller":"kernel/kernel.go:115","msg":"output ren","color":"x","area":1,"libertyArea":4,"adjacentColor":""}
```

## Step [O1o1o0g22o3o0] 相手の眼に石を置くことの禁止 - OpponentEye

囲碁のルールでは、相手の眼へは石を置けない。これを判定する  

とりあえず、このルールへ直訳で短い名前を付ける。 仮に `OpponentEye` とでも呼ぶことにする  

このルールは、あとで出てくる `Captured` のルールよりは優先度が低いとする  

### Step [O1o1o0g22o3o1o0] ファイル編集 - kernel.go

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	└── 📄 board1.txt
	├── 📂 kernel
	│	├── 📂 play_rule
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 board.go
 	│	├── 📄 check_board.go
 	│	├── 📄 color.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	├── 📄 kernel.go
 	│	├── 📄 liberty.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
👉 	│	├── 📄 play.go
 	│	├── 📄 point.go
 	│	├── 📄 ren.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
	└── 📄 main.go
```

👇 がんばって挿入してほしい  

```go
// func (k *Kernel) DoPlay(command string, logg *Logger) {

	// ...略...
	// [O1o1o0g22o3o1o0] 相手の眼に石を置こうとしました
	var onOpponentEye = func() bool {
		logg.C.Infof("? opponent_eye my_stone:%s point:%s\n", stone, k.Board.GetCodeFromPoint(point))
		logg.J.Infow("error opponent_eye", "my_stone", stone, "point", k.Board.GetCodeFromPoint(point))
		return false
	}

//	var isOk = k.Play(stone, point, logg,
//		// [O1o1o0g22o1o2o0] ,onMasonry
//		onMasonry,
		// [O1o1o0g22o3o1o0] ,onOpponentEye
		onOpponentEye//)
//
//	if isOk {
//		k.Record.Push(point)
//		logg.C.Info("=\n")
//		logg.J.Infow("ok")
//	}
// }

// func (k *Kernel) Play(stoneA Stone, pointB Point, logg *Logger,
	// // [O1o1o0g22o1o2o0] onMasonry
	// onMasonry func() bool,
	// [O1o1o0g22o3o1o0] onOpponentEye
	onOpponentEye func() bool//) bool {

	// ...略...
	// // [O1o1o0g22o1o2o0]
	// if k.IsMasonryError(stone, point) {
	//	return onMasonry()
	// }

	// [O1o1o0g22o3o1o0]
	var renC = k.GetLiberty(pointB)
	if renC.GetArea() == 1 { // 石Aを置いた交点を含む連Cについて、連Cの面積が1である（眼）
		if stoneA.GetColor() == renC.AdjacentColor.GetOpponent() {
			// かつ、連Cに隣接する連の色が、石Aのちょうど反対側の色であったなら、
			// 相手の眼に石を置こうとしたとみなす
			return onOpponentEye()
		}
	}

	// ...略...
	// k.Board.cells[point] = stone
	// return true
```

### Step [O1o1o0g22o3o2o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
set_board file data/board1.txt
play white C3
```

Output > Console:  

```plaintext
[2022-09-17 00:41:29]   # play white C3
[2022-09-17 00:41:29]   ? opponent_eye my_stone:o point:C3
```

## Step [O1o1o0g22o4o0] 自分の眼に石を置くことの任意の禁止

囲碁のルール上可能だが、明らかに損な手は、プレイアウトから除外したい。  
対局時には許可し、プレイアウト時には禁止するよう、選択できるようにする  

### Step [O1o1o0g22o4o1o0] ファイル編集 - play.go

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	└── 📄 board1.txt
	├── 📂 kernel
	│	├── 📂 play_rule
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 board.go
 	│	├── 📄 check_board.go
 	│	├── 📄 color.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	├── 📄 kernel.go
 	│	├── 📄 liberty.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
👉 	│	├── 📄 play.go
 	│	├── 📄 point.go
 	│	├── 📄 ren.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
	└── 📄 main.go
```

👇 がんばって挿入してほしい  

```go
// ...略...
// type Kernel struct {
	// ...略...

	// CanNotPutOnMyEye - [O1o1o0g22o4o1o0] 自分の眼に石を置くことはできません
	CanNotPutOnMyEye bool
// }
// ...略...

// func (k *Kernel) DoPlay(command string, logg *Logger) {

	// ...略...
	// [O1o1o0g22o4o1o0] 自分の眼に石を置こうとしました
	var onForbiddenMyEye = func() bool {
		logg.C.Infof("? my_eye my_stone:%s point:%s\n", stone, k.Board.GetCodeFromPoint(point))
		logg.J.Infow("error my_eye", "my_stone", stone, "point", k.Board.GetCodeFromPoint(point))
		return false
	}

//	var isOk = k.Play(stone, point, logg,
//		// [O1o1o0g22o1o2o0] ,onMasonry
//		onMasonry,
//		// [O1o1o0g22o3o1o0] ,onOpponentEye
//		onOpponentEye,
		// [O1o1o0g22o4o1o0] ,onForbiddenMyEye
		onForbiddenMyEye//)
//
//	if isOk {
//		k.Record.Push(point)
//		logg.C.Info("=\n")
//		logg.J.Infow("ok")
//	}
// }

// func (k *Kernel) Play(stoneA Stone, pointB Point, logg *Logger,
	// // [O1o1o0g22o1o2o0] onMasonry
	// onMasonry func() bool,
	// [O1o1o0g22o3o1o0] onOpponentEye
	onOpponentEye func() bool,
	// [O1o1o0g22o4o1o0]
	onForbiddenMyEye func() bool//) bool {

	// ...略...
	// // [O1o1o0g22o3o1o0]
	// var renC = k.GetLiberty(pointB)
	// if renC.GetArea() == 1 { // 石Aを置いた交点を含む連Cについて、連Cの面積が1である（眼）
	// 	if stoneA.GetColor() == renC.AdjacentColor.GetOpponent() {
			// かつ、連Cに隣接する連の色が、石Aのちょうど反対側の色であったなら、
			// 相手の眼に石を置こうとしたとみなす
	// 		return onOpponentEye()

		} else if k.CanNotPutOnMyEye && stoneA.GetColor() == renC.AdjacentColor {
			// [O1o1o0g22o4o1o0]
			// かつ、連Cに隣接する連の色が、石Aの色であったなら、
			// 自分の眼に石を置こうとしたとみなす
			return onForbiddenMyEye()

	// }

	// ...略...
	// k.Board.cells[point] = stone
	// return true
```

### Step [O1o1o0g22o4o2o_1o0] ファイル編集 - kernel.go

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	└── 📄 board1.txt
	├── 📂 kernel
	│	├── 📂 play_rule
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 board.go
 	│	├── 📄 check_board.go
 	│	├── 📄 color.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
👉 	│	├── 📄 kernel.go
 	│	├── 📄 liberty.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
 	│	├── 📄 play.go
 	│	├── 📄 point.go
 	│	├── 📄 ren.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
	└── 📄 main.go
```

👇 がんばって挿入してほしい  

```go
	// ...略...
	// この下にコマンドを挟んでいく
	// -------------------------
	// ...略...

	// * アルファベット順になる位置に、以下のケース文を挿入
	case "can_not_put_on_my_eye": // [O1o1o0g22o4o2o_1o0]
		// Example 1: "can_not_put_on_my_eye get"
		// Example 2: "can_not_put_on_my_eye set true"
		var method = tokens[1]
		switch method {
		case "get":
			var value = k.CanNotPutOnMyEye
			logg.C.Infof("= %t\n", value)
			logg.J.Infow("ok", "value", value)
			return true

		case "set":
			var value = tokens[2]
			switch value {
			case "true":
				k.CanNotPutOnMyEye = true
				return true
			case "false":
				k.CanNotPutOnMyEye = false
				return true
			default:
				logg.C.Infof("? unexpected method:%s value:%s\n", method, value)
				logg.J.Infow("error", "method", method, "value", value)
				return true
			}

		default:
			logg.C.Infof("? unexpected method:%s\n", method)
			logg.J.Infow("error", "method", method)
			return true
		}

	// ...略...
	// この上にコマンドを挟んでいく
	// -------------------------
	// ...略...
```

### Step [O1o1o0g22o4o2o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
set_board file data/board1.txt
play black C3
```

Output > Console:  

```plaintext
[2022-09-17 09:11:48]   # play black C3
[2022-09-17 09:11:48]   =
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
set_board file data/board1.txt
can_not_put_on_my_eye set true
play black C3
```

Output > Console:  

```plaintext
[2022-09-17 09:11:48]   # play black C3
[2022-09-17 09:11:48]   ? my_eye my_stone:x point:C3
```

## Step [O1o1o0g22o5o0] 任意の連の打ち上げ - RemoveRen 関数作成

### Step [O1o1o0g22o5o1o0] ファイル作成 - kernel_facade.go

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	└── 📄 board1.txt
	├── 📂 kernel
	│	├── 📂 play_rule
	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 board.go
 	│	├── 📄 check_board.go
 	│	├── 📄 color.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
👉 	│	├── 📄 kernel_facade.go
 	│	├── 📄 kernel.go
 	│	├── 📄 liberty.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
 	│	├── 📄 play.go
 	│	├── 📄 point.go
 	│	├── 📄 ren.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
	└── 📄 main.go
```

```go
// BOF [O1o1o0g22o5o1o0]

package kernel

// RemoveRen - 石の連を打ち上げます
func (k *Kernel) RemoveRen(ren *Ren) {
	for _, point := range ren.Elements {
		k.Board.SetStoneAt(point, Space)
	}
}

// EOF [O1o1o0g22o5o1o0]
```

### Step [O1o1o0g22o5o2o0] コマンド実装 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	└── 📄 board1.txt
	├── 📂 kernel
	│	├── 📂 play_rule
	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 board.go
 	│	├── 📄 check_board.go
 	│	├── 📄 color.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	├── 📄 kernel_facade.go
👉 	│	├── 📄 kernel.go
 	│	├── 📄 liberty.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
 	│	├── 📄 play.go
 	│	├── 📄 point.go
 	│	├── 📄 ren.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
	└── 📄 main.go
```

👇 がんばって、 Execute メソッドに挿入してほしい  

```go
	// ...略...
	// この下にコマンドを挟んでいく
	// -------------------------
	// ...略...

	// * アルファベット順になる位置に、以下のケース文を挿入
	case "remove_ren": // [O1o1o0g22o5o2o0]
		// Example: "remove_ren B2"
		var point = k.Board.GetPointFromCode(tokens[1])
		var ren = k.GetLiberty(point)
		k.RemoveRen(ren)
		logg.C.Infof("=\n")
		logg.J.Infow("ok")
		return true

	// ...略...
	// この上にコマンドを挟んでいく
	// -------------------------
	// ...略...
```

### Step [O1o1o0g22o5o3o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
set_board file data/board1.txt
remove_ren B2
```

Output > Console:  

```plaintext
[2022-09-17 12:17:02]   # remove_ren B2
[2022-09-17 12:17:02]   =
```

## Step [O1o1o0g22o6o0] 死に石の連の打ち上げ - Captured

石Aを盤上に置いて指を離したばかりの盤面とする。  
石Aに隣接する相手の石の連のうち、呼吸点が０のものは打ち上げる。  

このとき、 `OpponentEye` のルールと相反することがある

### Step [O1o1o0g22o6o1o0] ファイル編集 - play.go

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	└── 📄 board1.txt
	├── 📂 kernel
	│	├── 📂 play_rule
	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 board.go
 	│	├── 📄 check_board.go
 	│	├── 📄 color.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	├── 📄 kernel_facade.go
 	│	├── 📄 kernel.go
 	│	├── 📄 liberty.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
👉 	│	├── 📄 play.go
 	│	├── 📄 point.go
 	│	├── 📄 ren.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
	└── 📄 main.go
```

👇 がんばって挿入してほしい  

```go
// func (k *Kernel) Play(stoneA Stone, pointB Point, logg *Logger,

	// [O1o1o0g22o6o1o0] Captured ルール
	var isExists4rensToRemove = false
	var o4rensToRemove [4]*Ren
	var isChecked4rensToRemove = false

	// ...略...
	// [O1o1o0g22o3o1o0]
	// var renC = k.GetLiberty(pointB)
	// if renC.GetArea() == 1 {
		// if stoneA.GetColor() == renC.AdjacentColor.GetOpponent() {

			// * 以下を追加
			// [O1o1o0g22o6o1o0] 打ちあげる死に石の連を取得
			k.Board.SetStoneAt(pointB, stoneA) // いったん、石を置く
			isExists4rensToRemove, o4rensToRemove = k.GetRenToCapture(pointB)
			isChecked4rensToRemove = true
			k.Board.SetStoneAt(pointB, Space) // 石を取り除く

			if !isExists4rensToRemove {
				// `Captured` ルールと被らなければ
				return onOpponentEye()
			}

			// * 以下を削除
			// onOpponentEye()

		// } else if k.CanNotPutOnMyEye && stoneA.GetColor() == renC.AdjacentColor {
			// ...略...
		// }
	// }

	// ...略...
	// 石を置く
	// k.Board.SetStoneAt(pointB, stoneA)

	// * 以下を追加
	// [O1o1o0g22o6o1o0] 打ちあげる死に石の連を取得
	if !isChecked4rensToRemove {
		isExists4rensToRemove, o4rensToRemove = k.GetRenToCapture(pointB)
	}

	// [O1o1o0g22o6o1o0] 死に石を打ちあげる
	if isExists4rensToRemove {
		k.Remove4Rens(o4rensToRemove)
	}

	// return true
// }

// GetRenToCapture - 現在、着手後の盤面とします。打ち上げられる石の連を返します
//
// Returns
// -------
// isExists : bool
// renToRemove : [4]*Ren
// 隣接する東、北、西、南にある石を含む連
func (k *Kernel) GetRenToCapture(placePlay Point) (bool, [4]*Ren) {
	// [O1o1o0g22o6o1o0]
	var isExists bool
	var rensToRemove [4]*Ren

	var setAdjacentPoint = func(dir int, adjacentP Point) {
		var adjacentR = k.GetLiberty(adjacentP)
		if adjacentR.LibertyArea < 1 {
			isExists = true
			rensToRemove[dir] = adjacentR
		}
	}

	// 隣接する４方向
	k.Board.ForeachNeumannNeighborhood(placePlay, setAdjacentPoint)

	return isExists, rensToRemove
}

// Remove4Rens - 最大４つの連を盤上から打ち上げます
func (k *Kernel) Remove4Rens(o4rensToRemove [4]*Ren) {
	// [O1o1o0g22o6o1o0]
	for dir := 0; dir < 4; dir++ {
		if o4rensToRemove[dir] != nil {
			k.RemoveRen(o4rensToRemove[dir])
		}
	}
}
```

### Step [O1o1o0g22o6o2o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
set_board file data/board2.txt
play black D4
```

Output > Console:  

```plaintext
[2022-09-17 14:35:58]   # play black D4
[2022-09-17 14:35:58]   =
```

## Step [O1o1o0g22o7o0] コウの禁止 - Ko

自分が１手前に置いたところに２手続けて置けない

### Step [O1o1o0g22o7o1o0] ファイル編集 - kernel.go

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	└── 📄 board1.txt
	├── 📂 kernel
	│	├── 📂 play_rule
	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 board.go
 	│	├── 📄 check_board.go
 	│	├── 📄 color.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	├── 📄 kernel_facade.go
👉 	│	├── 📄 kernel.go
 	│	├── 📄 liberty.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
 	│	├── 📄 play.go
 	│	├── 📄 point.go
 	│	├── 📄 ren.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
	└── 📄 main.go
```

```go
type Kernel struct {
	// ...略...

	// Ko - [O1o1o0g22o7o1o0] コウの位置
	Ko Point
}
```

### Step [O1o1o0g22o7o2o0] ファイル編集 - play.go

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	└── 📄 board1.txt
	├── 📂 kernel
	│	├── 📂 play_rule
	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 board.go
 	│	├── 📄 check_board.go
 	│	├── 📄 color.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	├── 📄 kernel_facade.go
 	│	├── 📄 kernel.go
 	│	├── 📄 liberty.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
👉 	│	├── 📄 play.go
 	│	├── 📄 point.go
 	│	├── 📄 ren.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
	└── 📄 main.go
```

```go

```

TODO 東、北、西、南に隣接する連の重複を省く

# 参考にした記事

## Go言語

### ファイル分割

📖 [[Go言語] ファイル分割とローカルパッケージ](https://zenn.dev/fm_radio/articles/ca2ff1dfcf89b5)  

### コマンドライン引数

📖 [Goでflagを使ってコマンドライン引数を扱う](https://qiita.com/Yaruki00/items/7edc04720a24e71abfa2)  

### 列挙型

📖 [How to make Go print enum fields as string?](https://stackoverflow.com/questions/41480543/how-to-make-go-print-enum-fields-as-string)  

### 出力

📖 [Go で値を出力する方法](https://golang.keicode.com/basics/go-print-basics.php)  

### 可変長引数

📖 [Concatenating and Building Strings in Go 1.10+](https://www.calhoun.io/concatenating-and-building-strings-in-go/)  
📖 [Convert interface to string](https://yourbasic.org/golang/interface-to-string/)  

### 文字列

📖 [Go: 1文字ずつアクセスする](https://blog.sarabande.jp/post/61104920128)  

### 配列

📖 [スライス操作(要素の追加・削除, ソート, 他のスライスと結合)](https://www.wakuwakubank.com/posts/782-go-slice/)  

### ファイル入出力

📖 [Read a file in Go](https://gosamples.dev/read-file/#:~:text=The%20simplest%20way%20of%20reading,by%20line%20or%20in%20chunks.)  

.
