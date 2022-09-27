目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O11o__10o0] ロガー設定 ～ Step [O11o__10o3o_2o0] welcome プログラム

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# Step [O11o__10o0] ロガー設定

## Step [O11o__10o1o0] インストール

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go get -u go.uber.org/zap
```

## Step [O11o__10o2o_1o0] 設定 - .gitignore ファイル

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

# [O11o__10o2o_1o0]
*.log

# この上に kifuwarabe-uec14 でリポジトリにコミットしないものを追加する
# ---------------------------------------------------------------
# ...略...
```

## Step [O11o__10o2o0] ファイル作成

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
// BOF [O11o__10o2o0]

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

// EOF [O11o__10o2o0]
```

## Step [O11o__10o3o_1o0] リモートリポジトリにプッシュ

がんばって git などを使い、 `github.com/muzudho/kifuwarabe-uec14/kernel` モジュールの各パッケージのソースを  
リモートリポジトリにプッシュしてほしい  

# Step [O11o__10o3o_2o0] welcome プログラム

## Step [O11o__10o3o0] ファイル編集

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

	// [O11o__10o3o0] ログファイル
	var plainTextLogFile, _ = os.OpenFile(engineConfig.GetPlainTextLog(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer plainTextLogFile.Close() // ログファイル使用済み時にファイルを閉じる
	// ログファイル
	var jsonLogFile, _ = os.OpenFile(engineConfig.GetJsonLog(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer jsonLogFile.Close() // ログファイル使用済み時にファイルを閉じる
	// カスタマイズしたロガーを使うなら
	var logg = kernel.NewSugaredLoggerForGame(plainTextLogFile, jsonLogFile) // customized LOGGer

	// この上に初期設定を追加していく
	// ---------------------------

	if name == "hello" { // [O9o0]
		// ...略...
		// この下に分岐を挟んでいく
		// ---------------------
		// ...略...


	} else if name == "welcome" { // [O11o__10o0]
		logg.C.Infof("Welcome! name:'%s' weight:%.1f x:%d", "nihon taro", 92.6, 3)
		logg.J.Infow("Welcome!",
			"name", "nihon taro", "weight", 92.6, "x", 3)


		// ...略...
		// この上に分岐を挟んでいく
		// ---------------------
		// ...略...
	}
```

## Step [O11o__10o4o0] tidy

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

## Step [O11o__10o5o0] 実行

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
