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

// EOF [O11o__10o_4o0]
