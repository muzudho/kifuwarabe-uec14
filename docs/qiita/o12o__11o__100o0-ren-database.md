目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O12o__11o__100o0] データファイル作成 ～ Step [O12o__11o__10o0] 連データベース定義 ～ Step [O12o__11o__10o5o__10o0] 連データベースのロード

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# Step [O12o__11o__100o0] データファイル作成 - data/ren_db1.json ファイル

あとで使うファイルを先に作成する  

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
👉 	│	└── 📄 ren_db1.json
	├── 📂 kernel
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

```plaintext
{
    "header": {
        "boardWidth": 19,
        "boardHeight": 19
    },
    "rens": {
        "001,01A": {
            "stone": "x",
            "locate": "A1",
            "liberty": "B1 A2"
        },
        "001,06N": {
            "stone": ".",
            "locate": "N6",
            "liberty": ""
        },
        "001,13E": {
            "stone": "o",
            "locate": "E13",
            "liberty": "E12 D13"
        }
    }
}
```

# ~~Step [O12o__11o__101o0]~~

Removed

# Step [O12o__11o__10o0] 連データベース定義

* 取った石の場所を記憶しておく構造を作成する
* 入出力ファイルの構造でもある

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会☆（＾ｑ＾）＜その４＞](http://grayscale2.dou-jin.com/go/%E7%9B%AE%E6%8C%87%E3%81%9B%EF%BC%81%E7%AC%AC%EF%BC%91%EF%BC%94%E5%9B%9E%EF%BC%B5%EF%BC%A5%EF%BC%A3%E6%9D%AF%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF%E3%83%BC%E5%9B%B2%E7%A2%81%E5%A4%A7%E4%BC%9A%E2%98%86%EF%BC%88%EF%BC%BE_19)  

## ~~Step [O12o__11o__10o1o0]~~

Removed  

## Step [O12o__11o__10o2o0] ファイル作成 - ren_db.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
	│	├── 📄 o12o__10o1o0_point.go
	│	├── 📄 ren_db_item.go
👉	│	├── 📄 ren_db.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
  	└── 📄 main.go
```

```go
// BOF [O12o__11o__10o2o0]

package kernel

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// RenId - 連データベースに格納される連のId
// * 外部ファイルの可読性を優先して数値型ではなく文字列
type RenId string

// GetRenId - 連のIdを取得
func GetRenId(boardMemoryWidth int, positionNthFigure int, positionNumber int, minimumLocation Point) RenId {
	var posNth = positionNumber + geta
	var coord = getRenIdFromPointOnBoard(boardMemoryWidth, minimumLocation)

	return RenId(fmt.Sprintf("%0*d,%s", positionNthFigure, posNth, coord))
}

// RenDb - 連データベース
type RenDb struct {
	// Header - ヘッダー
	Header RenDbDocHeader `json:"header"`

	// 要素
	Rens map[RenId]*Ren `json:"rens"`
}

// NewRenDb - 連データベースを新規作成
func NewRenDb(boardWidth int, boardHeight int) *RenDb {
	var db = new(RenDb)
	db.Header.Init(boardWidth, boardHeight)
	db.Rens = make(map[RenId]*Ren)
	return db
}

// Init - 初期化
func (db *RenDb) Init(boardWidth int, boardHeight int) {
	db.Header.Init(boardWidth, boardHeight)

	// Clear
	for ri := range db.Rens {
		delete(db.Rens, ri)
	}
}

// Save - 連データベースの外部ファイル書込
func (db *RenDb) Save(path string, convertLocation func(Point) string, onError func(error) bool) bool {

	// 外部ファイルに出力するための、内部状態の整形
	db.RefreshToExternalFile(convertLocation)

	// Marshal関数でjsonエンコード
	// ->返り値jsonDataにはエンコード結果が[]byteの形で格納される
	jsonBinary, errA := json.Marshal(db)
	if errA != nil {
		return onError(errA)
	}

	// ファイル読込
	var errB = os.WriteFile(path, jsonBinary, 0664)
	if errB != nil {
		return onError(errB)
	}

	return true
}

// FindRen - 連を取得
func (db *RenDb) GetRen(renId RenId) (*Ren, bool) {
	var ren, isOk = db.Rens[renId]

	if isOk {
		return ren, true
	}

	return nil, false
}

// RegisterRen - 連を登録
// * すでに Id が登録されているなら、上書きしない
func (db *RenDb) RegisterRen(positionNthFigure int, positionNumber int, ren *Ren) {
	var renId = GetRenId(db.Header.GetBoardMemoryWidth(), positionNthFigure, positionNumber, ren.minimumLocation)

	var _, isExists = db.Rens[renId]
	if !isExists {
		db.Rens[renId] = ren
	}
}

// Dump - ダンプ
func (db *RenDb) Dump() string {
	var sb strings.Builder

	// 全ての要素
	for idStr, item := range db.Rens {
		sb.WriteString(fmt.Sprintf("[%s]%s ", idStr, item.Dump()))
	}

	var text = sb.String()
	if 0 < len(text) {
		text = text[:len(text)-1]
	}
	return text
}

// RefreshToExternalFile - 外部ファイルに出力されてもいいように内部状態を整形します
func (db *RenDb) RefreshToExternalFile(convertLocation func(Point) string) {
	for _, ren := range db.Rens {
		ren.RefreshToExternalFile(convertLocation)
	}
}

// RenDbDocHeader - ヘッダー
type RenDbDocHeader struct {
	// BoardWidth - 盤の横幅
	BoardWidth int `json:"boardWidth"`
	// BoardHeight - 盤の縦幅
	BoardHeight int `json:"boardHeight"`
}

// Init - 初期化
func (h *RenDbDocHeader) Init(boardWidth int, boardHeight int) {
	h.BoardWidth = boardWidth
	h.BoardHeight = boardHeight
}

// GetBoardMemoryArea - 枠付き盤の面積
func (h *RenDbDocHeader) GetBoardMemoryArea() int {
	return h.GetBoardMemoryWidth() * h.GetBoardMemoryHeight()
}

// GetBoardMemoryWidth - 枠付き盤の横幅
func (h *RenDbDocHeader) GetBoardMemoryWidth() int {
	return h.BoardWidth + bothSidesWallThickness
}

// GetBoardMemoryHeight - 枠付き盤の縦幅
func (h *RenDbDocHeader) GetBoardMemoryHeight() int {
	return h.BoardHeight + bothSidesWallThickness
}

// EOF [O12o__11o__10o2o0]
```

## Step [O12o__11o__10o3o0] ファイル編集 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
👉 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
	│	├── 📄 o12o__10o1o0_point.go
	│	├── 📄 ren_db_item.go
	│	├── 📄 ren_db.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
  	└── 📄 main.go
```

```go
// ...略...
// type Kernel struct {
	// ...略...

	// * 以下を追加
	// RenDb - [O12o__11o__10o3o0] 連データベース
	renDb *RenDb
// }
// ...略...

// func NewKernel(boardWidht int, boardHeight int,
	// ...略...

	// * 以下を追加
	// RenDb - [O12o__11o__10o3o0] 連データベース
	k.renDb = NewRenDb(k.Board.coordinate.GetWidth(), k.Board.coordinate.GetHeight())

//	return k
// }
// ...略...
```

## Step [O12o__11o__10o4o0] コマンド編集 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
👉 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
	│	├── 📄 o12o__10o1o0_point.go
	│	├── 📄 ren_db_item.go
	│	├── 📄 ren_db.go
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
	case "rendb_dump": // [O12o__11o__10o4o0]
		var text = k.renDb.Dump()
		logg.C.Info("= dump'''%s\n'''\n", text)
		logg.J.Infow("ok", "dump", text)
		return true

	case "rendb_save": // [O12o__11o__10o4o0]
		// Example: `rendb_save data/ren_db1_temp.json`
		// * ファイルパスにスペースがはいっていてはいけない
		var path = tokens[1]

		var convertLocation = func(location Point) string {
			return k.Board.coordinate.GetGtpMoveFromPoint(location)
		}

		var onError = func(err error) bool {
			logg.C.Infof("? error:%s\n", err)
			logg.J.Infow("error", "err", err)
			return false
		}

		var isOk = k.renDb.Save(path, convertLocation, onError)
		if isOk {
			logg.C.Infof("=\n")
			logg.J.Infow("ok")
			return true
		}

		return false

	// ...略...
	// この上にコマンドを挟んでいく
	// -------------------------
	// ...略...
```


# Step [O12o__11o__10o5o__10o0] 連データベースのロード

連データベースをロードするには、盤のサイズも、連も既知でないといけない  
盤のサイズ、連の定義を終えた段階で、連データベースのロードを実装する  

## Step [O12o__11o__10o5o__10o_10o0] ファイル作成 - kernel/kernel_facade.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
👉 	│	├── 📄 kernel_facade.go
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
	│	├── 📄 o12o__10o1o0_point.go
	│	├── 📄 ren_db_item.go
	│	├── 📄 ren_db.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
  	└── 📄 main.go
```

```go
// BOF [O12o__11o__10o5o__10o0]

package kernel

import (
	"encoding/json"
	"os"
)

// LoadRenDb - [O12o__11o__10o5o__10o_10o0] 連データベースの外部ファイル読取
func (k *Kernel) LoadRenDb(path string, onError func(error) bool) bool {
	// ファイル読込
	var binary, errA = os.ReadFile(path)
	if errA != nil {
		return onError(errA)
	}

	var db = new(RenDb)
	var errB = json.Unmarshal(binary, db)
	if errB != nil {
		return onError(errB)
	}

	// 外部ファイルからの入力を、内部状態へ適用
	for _, ren := range db.Rens {
		var isOk = k.RefreshRenToInternal(ren)
		if !isOk {
			return false
		}
	}

	// 差し替え
	k.renDb = db
	return true
}

// RefreshRenToInternal - TODO 外部ファイルから入力された内容を内部状態に適用します
func (k *Kernel) RefreshRenToInternal(r *Ren) bool {
	{
		var getDefaultStone = func() (bool, Stone) {
			panic(fmt.Sprintf("unexpected stone:%s", r.Sto))
		}

		// TODO stone from r.Sto
		// Example: "x" --> black
		var isOk, stone = GetStoneFromChar(r.Sto, getDefaultStone)
		if !isOk {
			return false
		}
		r.stone = stone
	}
	{
		// TODO locations from r.Loc
		// Example: "C1 D1 E1"
		if 0 < len(r.Loc) {
			var codes = strings.Split(r.Loc, " ")

			var numbers = []Point{}
			for _, code := range codes {
				var location = k.Board.coordinate.GetPointFromGtpMove(code)
				numbers = append(numbers, location)
			}

			r.locations = numbers
		}
	}
	{
		// TODO libertyLocations from r.LibLoc
		// Example: "F1 E2 D2 B1 C2"
		if 0 < len(r.LibLoc) {
			var codes = strings.Split(r.LibLoc, " ")

			var numbers = []Point{}
			for _, code := range codes {
				var location = k.Board.coordinate.GetPointFromGtpMove(code)
				numbers = append(numbers, location)
			}

			r.libertyLocations = numbers
		}
	}

	return true
}

// EOF [O12o__11o__10o5o__10o0]
```

## Step [O12o__11o__10o5o__10o1o0] コマンド編集 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
👉 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
	│	├── 📄 o12o__10o1o0_point.go
	│	├── 📄 ren_db_item.go
	│	├── 📄 ren_db.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
  	└── 📄 main.go
```

```go
	// ...略...

	// この下にコマンドを挟んでいく
	// -------------------------
	// ...略...

	case "rendb_load": // [O12o__11o__10o5o__10o1o0]
		// Example: `rendb_load data/ren_db1_temp.json`
		// * ファイルパスにスペースがはいっていてはいけない
		var path = tokens[1]
		var onError = func(err error) bool {
			logg.C.Infof("? error:%s\n", err)
			logg.J.Infow("error", "err", err)
			return false
		}

		var isOk = k.LoadRenDb(path, onError)
		if isOk {
			logg.C.Infof("=\n")
			logg.J.Infow("ok")
			return true
		}
		return false

	// ...略...

	// この上にコマンドを挟んでいく
	// -------------------------
	// ...略...
```

## Step [O12o__11o__10o5o_1o0] 設定 - .gitignore ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
👉  └── 📄 .gitignore
```

👇 例えば冒頭に追加  

```plaintext
# この下に kifuwarabe-uec14 でリポジトリにコミットしないものを追加する
# ---------------------------------------------------------------
# ...略...

# [O12o__11o__10o5o_1o0]
*_temp.json

# この上に kifuwarabe-uec14 でリポジトリにコミットしないものを追加する
# ---------------------------------------------------------------
# ...略...
```

## Step [O12o__11o__10o5o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
rendb_load data/ren_db1.json
```

Output > Console:  

```plaintext
[2022-09-18 21:21:15]   # rendb_load data/ren_db1.json
[2022-09-18 21:21:15]   =
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
rendb_dump
```

Output > Console:  

```plaintext
[2022-09-25 22:22:44]   # rendb_dump
[2022-09-25 22:22:44]   = dump'''[001,01A]22
[001,06N]139
[001,13E]278
'''
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
rendb_save data/ren_db1_temp.json
```

Output > Console:  

```plaintext
[2022-09-18 22:15:43]   # rendb_save data/ren_db1_temp.json
[2022-09-18 22:15:43]   =
```

📄 ren_db1_temp.json:  

```json
{"header":{"boardWidth":19,"boardHeight":19},"rens":{"001,01A":{"stone":".","locate":"","liberty":""},"001,06N":{"stone":".","locate":"","liberty":""},"001,13E":{"stone":".","locate":"","liberty":""}}}
```
