目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O11o__10o0] ロガー設定

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
