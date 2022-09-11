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
	var configF = zap.NewProductionEncoderConfig()
	configF.EncodeTime = zapcore.ISO8601TimeEncoder // 日本時間のタイムスタンプ

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
