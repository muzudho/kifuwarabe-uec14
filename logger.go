// BOF [O1o1o0g11o__10o2o0]

package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// CreateSugaredLogger - ロガーを作成します
func CreateSugaredLogger(logFile *os.File) *zap.SugaredLogger {
	// 設定 > 製品用
	var config = zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder // 日本時間のタイムスタンプ

	// コア
	var core = zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(config), // JSON形式
			zapcore.Lock(os.Stderr),        // 出力先は標準エラー
			zapcore.DebugLevel),            // ログレベル
		zapcore.NewCore(
			zapcore.NewJSONEncoder(config), // JSON形式
			zapcore.AddSync(logFile),       // 出力先はログファイル
			zapcore.DebugLevel),            // ログレベル
	)

	// ロガーのビルド
	var logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	// 糖衣構文のインターフェースを取得
	return logger.Sugar()
}

// EOF [O1o1o0g11o__10o2o0]
