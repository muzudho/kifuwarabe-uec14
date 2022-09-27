# 目次

📖 [Step [O1o0] 導入](https://qiita.com/muzudho1/items/3a78087f812bab4a511f)  
📖 [Step [O11o___100o0] カーネル作成 ～ Step [O11o__10o_1o0] 思考エンジン設定ファイル](https://qiita.com/muzudho1/items/6c0b16d3b87ac598fe86)  
📖 [Step [O11o__10o0] ロガー設定 ～ Step [O11o__10o3o_2o0] welcome プログラム](https://qiita.com/muzudho1/items/26af2c9f5dcc16175acd)  
📖 [Step [O11o__11o0] デバッグ可能標準入力 作成](https://qiita.com/muzudho1/items/252eef6d00417dbd82a1)  
📖 [Step [O11o_3o0] カーネルのインタープリター ～ Step [O11o_4o0] 石の色定義 ～ Step [O11o_4o2o0] 連の定義 ～ Step [O11o_5o0] 石定義 ～ Step [O12o__10o0] 点定義、またはその盤座標符号定義](https://qiita.com/muzudho1/items/374b040f4e025f42b970)  
📖 [Step [O12o__11o__100o0] データファイル作成 ～ Step [O12o__11o__10o0] 連データベース定義 ～ Step [O12o__11o__10o5o__10o0] 連データベースのロード](https://qiita.com/muzudho1/items/03454c8efcf8b8927086)  
📖 [Step [O12o__11o_1o0] 棋譜定義](https://qiita.com/muzudho1/items/c00dd77f8cefc5c12dcd)  

# Step [O12o__11o0] 盤定義（土台）

これから盤を作っていく前に、土台を作る  

## Step [O12o__11o1o0] ファイル作成 - board.go

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
// BOF [O12o__11o1o0]

package kernel

// 片方の枠の厚み。東、北、西、南
const oneSideWallThickness = 1

// 両側の枠の厚み。南北、または東西
const bothSidesWallThickness = 2

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
func NewBoard(boardWidht int, boardHeight int) *Board {
	var b = new(Board)

	// 盤のサイズ指定と、盤面の初期化
	b.resize(boardWidht, boardHeight)

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
	return b.memoryWidth - bothSidesWallThickness
}

// GetHeight - 枠の厚みを含まない縦幅
func (b *Board) GetHeight() int {
	return b.memoryHeight - bothSidesWallThickness
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
	return getXyFromPointOnBoard(b.memoryWidth, point)
}

func getXyFromPointOnBoard(memoryWidth int, point Point) (int, int) {
	var p = int(point)
	return p % memoryWidth, p / memoryWidth
}

// サイズ変更
func (b *Board) resize(width int, height int) {
	b.memoryWidth = width + bothSidesWallThickness
	b.memoryHeight = height + bothSidesWallThickness
	b.cells = make([]Stone, b.getMemoryArea())
}

// 枠付き盤の面積
func (b *Board) getMemoryArea() int {
	return b.memoryWidth * b.memoryHeight
}

// EOF [O12o__11o1o0]
```

## Step [O12o__11o2o0] コマンド実装 - kernel.go ファイル

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
	case "test_get_point_from_xy": // [O12o__11o2o0]
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

## Step [O12o__11o3o0] 動作確認

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

# Step [O12o_1o0] 盤定義（盤面）

## Step [O12o0] ファイル作成 - board_area.go

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
// BOF [O12o0]

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

// ForeachPayloadLocation - 枠や改行を含めない各セルの番地
func (b *Board) ForeachPayloadLocation(setLocation func(Point)) {
	var height = b.memoryHeight - 1
	var width = b.memoryWidth - 1
	for y := 1; y < height; y++ {
		for x := 1; x < width; x++ {
			var i = b.GetPointFromXy(x, y)
			setLocation(i)
		}
	}
}

// ForeachPayloadLocation - 枠や改行を含めない各セルの番地。筋、段の順
func (b *Board) ForeachPayloadLocationOrderByYx(setLocation func(Point)) {
	var height = b.memoryHeight - 1
	var width = b.memoryWidth - 1
	for x := 1; x < width; x++ {
		for y := 1; y < height; y++ {
			var i = b.GetPointFromXy(x, y)
			setLocation(i)
		}
	}
}

// ForeachNeumannNeighborhood - [O13o__10o0] 隣接する４方向の定義
func (b *Board) ForeachNeumannNeighborhood(here Point, setAdjacent func(int, Point)) {
	// 東、北、西、南
	for dir := 0; dir < 4; dir++ {
		var p = here + Point(b.Direction[dir]) // 隣接する交点

		// 範囲外チェック
		if p < 1 || b.getMemoryArea() <= int(p) {
			continue
		}

		// 壁チェック
		if b.GetStoneAt(p) == Wall {
			continue
		}

		setAdjacent(dir, p)
	}
}

// EOF [O12o0]
```

## Step [O12o1o0] カレントディレクトリーを移動

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd kernel
```

## Step [O12o2o0] tidy

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go get -u
go mod tidy
```

## Step [O12o3o0] カレントディレクトリーを戻す

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd ..
```

## Step [O12o4o0] リモートリポジトリにプッシュ

がんばって git などを使い、 `github.com/muzudho/kifuwarabe-uec14/kernel` モジュールの各パッケージのソースを  
リモートリポジトリにプッシュしてほしい  

## Step [O12o5o0] tidy

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

# Step [O13o_1o0] 盤表示 - board コマンド

## Step [O13o0] コマンド実装 - kernel.go ファイル

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

	case "board": // [O13o0]
		// 人間向けの出力
		{
			// 25列まで対応
			const fileSimbols = "ABCDEFGHJKLMNOPQRSTUVWXYZ"
			// 25行まで対応
			var rankSimbols = strings.Split("  , 1, 2, 3, 4, 5, 6, 7, 8, 9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25", ",")

			var filesMax = int(math.Min(25, float64(k.Board.GetWidth())))
			var rowsMax = int(math.Min(25, float64(k.Board.GetHeight())))
			var filesLabel = fileSimbols[:filesMax]

			var sb strings.Builder
			// 枠の上辺
			sb.WriteString(fmt.Sprintf(`= board:'''
.     %s
.    `, filesLabel))

			var rowNumber = 1
			var setStone = func(s Stone) {
				sb.WriteString(fmt.Sprintf("%v", s))
			}
			var doNewline = func() {
				var rankLabel string
				if rowNumber <= rowsMax {
					rankLabel = rankSimbols[rowNumber]
				} else {
					rankLabel = ""
				}

				sb.WriteString(fmt.Sprintf("\n. %2s ", rankLabel))
				rowNumber++
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

## Step [O14o0] 動作確認

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

Output > Console:  

```plaintext
[2022-09-19 15:50:11]   # board
[2022-09-19 15:50:11]   = board:'''
.     ABCDEFGHJKLMNOPQRST
.    +++++++++++++++++++++
.  1 +...................+
.  2 +...................+
.  3 +...................+
.  4 +...................+
.  5 +...................+
.  6 +...................+
.  7 +...................+
.  8 +...................+
.  9 +...................+
. 10 +...................+
. 11 +...................+
. 12 +...................+
. 13 +...................+
. 14 +...................+
. 15 +...................+
. 16 +...................+
. 17 +...................+
. 18 +...................+
. 19 +...................+
.    +++++++++++++++++++++
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

# Step [O15o__10o0] 盤サイズの変更 - resize コマンド

## Step [O15o__11o0] コマンド実装 - kernel.go ファイル

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

	case "boardsize": // [O15o__11o0]
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

## Step [O15o__12o0] 動作確認

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

# ~~Step [O15o__13o0]~~

Removed  

## ~~Step [O15o__13o1o0]~~

Moved to `O11o__10o_2o0`  

## ~~Step [O15o__13o2o_1o0]~~

Moved to `[O11o__10o_3o0]`  

## ~~Step [O15o__13o2o_2o0]~~

Moved to `[O11o__10o_4o0]`  

## ~~Step [O15o__13o2o_3o0]~~

Merged to `[O11o_3o0]`  

## ~~Step [O15o__13o2o_4o0]~~

Moved to `[O11o__10o_6o0]`  

# Step [O15o__14o_1o0] データファイル作成 - data/board1.txt ファイル

あとで使うファイルを先に作成する  

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
👉 	│	└── 📄 board1.txt
	├── 📂 kernel
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
 	│	├── 📄 board_set.go
  	│	├── 📄 board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
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

# Step [O15o__14o0] 初期盤面を設定する - board_set コマンド

## Step [O15o__14o1o0] ファイル作成 - board_set.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
👉 	│	├── 📄 board_set.go
  	│	├── 📄 board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
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
// BOF [O15o__14o1o0]

package kernel

import (
	"os"
	"strings"
)

// DoSetBoard - 盤面を設定する
//
// コマンドラインの複数行入力は難しいので、ファイルから取ることにする
// * `command` - Example: `board_set file data/board1.txt`
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

// EOF [O15o__14o1o0]
```

## Step [O15o__14o2o0] コマンド実装 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
 	│	├── 📄 board_set.go
  	│	├── 📄 board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
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

	case "board_set": // [O15o__14o2o0]
		// Example: `board_set file data/board1.txt`
		k.DoSetBoard(command, logg)
		logg.C.Infof("=\n")
		logg.J.Infow("ok")
		return true

	// ...略...
	// この上にコマンドを挟んでいく
	// -------------------------
	// ...略...
```

## Step [O15o__14o3o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
board_set file data/board1.txt
board
```

出力は略  

# Step [O15o_1o0] 座標の定義

## ~~Step [O15o0]~~

Moved to `[O12o__10o1o0]`  

## Step [O16o0] ファイル作成 - board_coord.go ファイル

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
// BOF [O16o0]

package kernel

import "fmt"

// GetPointFromCode - "A7" や "J13" といった符号を Point へ変換します
//
// * `code` - 座標の符号。 Example: "A7" や "J13"
func (b *Board) GetPointFromCode(code string) Point {
	return b.GetPointFromXy(
		GetXFromFile(GetFileFromCode(code))+oneSideWallThickness,
		GetYFromRank(GetRankFromCode(code))+oneSideWallThickness)
}

// GetCodeFromPoint - `GetPointFromCode` の逆関数
func (b *Board) GetCodeFromPoint(point Point) string {
	return getCodeFromPointOnBoard(b.memoryWidth, point)
}

// 例えば "A1" のように、行番号をゼロサプレスして返す
func getCodeFromPointOnBoard(memoryWidth int, point Point) string {
	var file, rank = getFileRankFromPointOnBoard(memoryWidth, point)
	return fmt.Sprintf("%s%d", file, rank)
}

// 例えば "01A" のように、符号を行、列の順とし、かつ、行番号を一律２桁のゼロ埋めにします
func getRenIdFromPointOnBoard(memoryWidth int, point Point) string {
	var file, rank = getFileRankFromPointOnBoard(memoryWidth, point)
	return fmt.Sprintf("%02d%s", rank, file)
}

func getFileRankFromPointOnBoard(memoryWidth int, point Point) (string, int) {
	var x, y = getXyFromPointOnBoard(memoryWidth, point)
	var file = GetFileFromX(x - oneSideWallThickness)
	var rank = getRankFromY(y) - oneSideWallThickness
	return file, rank
}

// EOF [O16o0]
```

### Step [O16o1o0] コマンド実装 - kernel.go ファイル

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
	case "test_get_point_from_code": // [O16o1o0]
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

## Step [O16o2o0] 動作確認

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

## ~~Step [O17o0]~~

Moved to `[O12o__10o2o0]`  

## ~~Step [O18o0]~~

Moved to `[O12o__10o3o0]`  

# Step [O19o_1o0] 石を打つ - play コマンド

## Step [O19o0] ファイル作成 - play.go ファイル

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
// BOF [O19o0]

package kernel

import "strings"

// DoPlay - 打つ
//
// * `command` - Example: `play black A19`
// ........................---- ----- ---
// ........................0    1     2
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

	// [O22o1o2o0] 石（または壁）の上に石を置こうとした
	var onMasonry = func() bool {
		logg.C.Infof("? masonry my_stone:%s point:%s\n", stone, k.Board.GetCodeFromPoint(point))
		logg.J.Infow("error", "my_stone", stone, "point", k.Board.GetCodeFromPoint(point))
		return false
	}

	var isOk = k.Play(stone, point, logg,
		// [O22o1o2o0] ,onMasonry
		onMasonry)

	if isOk {
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
	// [O22o1o2o0] onMasonry
	onMasonry func() bool) bool {

	// [O22o1o2o0]
	if k.IsMasonryError(stoneA, pointB) {
		return onMasonry()
	}

	// 石を置く
	k.Board.SetStoneAt(pointB, stoneA)

	k.Record.Push(pointB) // 棋譜に追加

	return true
}

// EOF [O19o0]
```

## Step [O20o0] 実装 - kernel.go ファイル

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
	case "play": // [O20o0]
		// Example: `play black A19`
		k.DoPlay(command, logg)
		return true

	// ...略...
	// この上にコマンドを挟んでいく
	// -------------------------
	// ...略...
```

## Step [O21o0] 実行

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

# Step [O22o_1o0] データファイル作成 - data/board2.txt ファイル

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
 	│	├── 📄 board_set.go
  	│	├── 📄 board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
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

# Step [O22o_2o0] データファイル作成 - data/board3.txt ファイル

あとで使うファイルを先に作成する  

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	├── 📄 board1.txt
 	│	├── 📄 board2.txt
👉 	│	└── 📄 board3.txt
	├── 📂 kernel
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
 	│	├── 📄 board_set.go
  	│	├── 📄 board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
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
    ABCDEFGHJKLMNOPQRST
   +++++++++++++++++++++
 1 +...................+
 2 +..xo...............+
 3 +.xo.o..............+
 4 +..xo...............+
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

# Step [O22o0] 囲碁の石を打つルールの実装

## Step [O22o1o0] 空点以外のところ（石または壁の上）に石を置くことの禁止 - IsMasonryError関数作成

とりあえず、 `石または壁の上に石を置く行為` に `Masonry` （メイスンリー）という名前を付ける。  
従って この主のエラーは `Masonry error` と呼ぶことにする。  
そのようなエラーであるかどうか判定する関数の名前は `IsMasonryError` と呼ぶことにする  

### Step [O22o1o1o0] ファイル作成 - masonry.go

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
// BOF [O22o1o1o0]

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

// EOF [O22o1o1o0]
```

### Step [O22o1o2o0] 呼出し

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
	// [O22o1o2o0]
	var onMasonry = func() bool {
		logg.C.Infof("? masonry my_stone:%s point:%s\n", stone, point)
		logg.J.Infow("error", "my_stone", stone, "point", point)
		return false
	}

	// var isOk = k.Play(stone, point, logg,
		// * 以下を追加
		// [O22o1o2o0] ,onMasonry
		onMasonry//)
	// ...略...
// }

// func (k *Kernel) Play(stone Stone, point Point, logg *Logger,
	// * 以下を追加
	// [O22o1o2o0] onMasonry
	onMasonry func() bool//) bool {

	// * 以下を追加
	// [O22o1o2o0]
	if k.IsMasonryError(stone, point) {
		return onMasonry()
	}

//	k.Board.cells[point] = stone
//	return true
// }
```

## Step [O22o2o0] 連の認識と、呼吸点のカウント - GetLiberty 関数作成

盤上の座標を指定することで、そこにある `連` の `呼吸点` の数を算出したい  

* `連` - Ren、れん。コンピューター囲碁用語。説明は省略
* `呼吸点` - Liberty、こきゅうてん。コンピューター囲碁用語。説明は省略

👇 呼吸点を数えるために探索すると、一緒に以下のことも行える  

* 連の認識
* 隣接する連の色の取得

このような探索を行う関数に名前を付ける。  
`GetRen` がふさわしいが、慣習を優先して `GetLiberty` と名付けることにする  

### ~~Step [O22o2o1o0]~~

Moved to `[O11o_4o2o1o0]`  

### Step [O22o2o2o0] ファイル作成 - check_board.go

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
// BOF [O22o2o2o0]

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
	cells []uint8
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
		b.cells[i] = 0
	}
}

// Resize - サイズ変更
func (b *CheckBoard) Resize(width int, height int) {
	b.memoryWidth = width + 2
	b.memoryHeight = height + 2
	b.cells = make([]uint8, b.getMemoryArea())
}

// CheckStone - 石をチェックした
func (b *CheckBoard) CheckStone(p Point) {
	b.cells[p] |= 0b00000001
}

// IsChecked - 石はチェックされているか？
func (b *CheckBoard) IsStoneChecked(p Point) bool {
	return b.cells[p]&0b00000001 == 0b00000001
}

// CheckLiberty - 呼吸点をチェックした
func (b *CheckBoard) CheckLiberty(p Point) {
	b.cells[p] |= 0b00000010
}

// UncheckLiberty - 呼吸点のチェックを外した
func (b *CheckBoard) UncheckLiberty(p Point) {
	b.cells[p] &= 0b11111101
}

// IsLibertyChecked - 呼吸点はチェックされているか？
func (b *CheckBoard) IsLibertyChecked(p Point) bool {
	return b.cells[p]&0b00000010 == 0b00000010
}

// 枠付き盤の面積
func (b *CheckBoard) getMemoryArea() int {
	return b.memoryWidth * b.memoryHeight
}

// EOF [O22o2o2o0]
```

### Step [O22o2o3o_1o0] ファイル編集 - board.go

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

### Step [O22o2o3o0] ファイル編集 - kernel.go

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
	// [O22o2o3o0]
	// CheckBoard - 呼吸点の探索時に使います
	CheckBoard *CheckBoard
	// tempRen - 呼吸点の探索時に使います
	tempRen *Ren
//}

// func NewKernel(boardWidht int, boardHeight int) *Kernel {
//	var k = new(Kernel)
//	k.Board = NewBoard(boardWidht, boardHeight)

	// * 以下を追加
	// [O22o2o3o0]
	k.CheckBoard = NewCheckBoard()

//	return k
// }
```

### Step [O22o2o4o0] ファイル作成 - liberty.go

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
// BOF [O22o2o4o0]

package kernel

// GetLiberty - 呼吸点の数え上げ。連の数え上げ。
// `GetOneRen` とでもいう名前の方がふさわしいが、慣習に合わせた関数名にした
//
// Parameters
// ----------
// * `arbitraryPoint` - 連に含まれる任意の一点
//
// Returns
// -------
// - *Ren is ren or nil
// - bool is found
func (k *Kernel) GetLiberty(arbitraryPoint Point) (*Ren, bool) {
	// チェックボードの初期化
	k.CheckBoard.Init(k.Board.GetWidth(), k.Board.GetHeight())
	return k.findRen(arbitraryPoint)
}

// 連の検索
//
// Returns
// -------
// - *Ren is ren or nil
// - bool is found
func (k *Kernel) findRen(arbitraryPoint Point) (*Ren, bool) {
	// 探索済みならスキップ
	if k.CheckBoard.IsStoneChecked(arbitraryPoint) {
		return nil, false
	}

	// 連の初期化
	k.tempRen = NewRen(k.Board.GetStoneAt(arbitraryPoint))

	if k.tempRen.stone == Space {
		k.searchSpaceRen(arbitraryPoint)
	} else {
		k.searchStoneRen(arbitraryPoint)

		// チェックボードの呼吸点のチェックをクリアー
		for _, p := range k.tempRen.libertyLocations {
			k.CheckBoard.UncheckLiberty(p)
		}
	}

	return k.tempRen, true
}

// 石の連の探索
// - 再帰関数
func (k *Kernel) searchStoneRen(here Point) {
	k.CheckBoard.CheckStone(here)
	k.tempRen.AddLocation(here)

	var setAdjacent = func(dir int, p Point) {
		// 呼吸点と壁のチェック
		var stone = k.Board.GetStoneAt(p)
		switch stone {
		case Space: // 空点
			if !k.CheckBoard.IsLibertyChecked(p) { // まだチェックしていない呼吸点なら
				k.CheckBoard.CheckLiberty(p)
				k.tempRen.libertyLocations = append(k.tempRen.libertyLocations, p) // 呼吸点を追加
			}

			return // あとの処理をスキップ

		case Wall: // 壁
			return // あとの処理をスキップ
		}

		// 探索済みの石ならスキップ
		if k.CheckBoard.IsStoneChecked(p) {
			return
		}

		var color = stone.GetColor()
		// 隣接する色、追加
		k.tempRen.adjacentColor = k.tempRen.adjacentColor.GetAdded(color)

		if stone == k.tempRen.stone { // 同じ石
			k.searchStoneRen(p) // 再帰
		}
	}

	// 隣接する４方向
	k.Board.ForeachNeumannNeighborhood(here, setAdjacent)
}

// 空点の連の探索
// - 再帰関数
func (k *Kernel) searchSpaceRen(here Point) {
	k.CheckBoard.CheckStone(here)
	k.tempRen.AddLocation(here)

	var setAdjacent = func(dir int, p Point) {
		// 探索済みならスキップ
		if k.CheckBoard.IsStoneChecked(p) {
			return
		}

		var stone = k.Board.GetStoneAt(p)
		if stone != Space { // 空点でなければスキップ
			return
		}

		var color = stone.GetColor()
		// 隣接する色、追加
		k.tempRen.adjacentColor = k.tempRen.adjacentColor.GetAdded(color)
		k.searchSpaceRen(p) // 再帰
	}

	// 隣接する４方向
	k.Board.ForeachNeumannNeighborhood(here, setAdjacent)
}

// EOF [O22o2o4o0]
```

### Step [O22o2o5o0] コマンド実装 - kernel.go ファイル

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
	case "test_get_liberty": // [O22o2o5o0]
		// Example: "test_get_liberty B2"
		var coord = tokens[1]
		var point = k.Board.GetPointFromCode(coord)
		var ren, isFound = k.GetLiberty(point)
		if isFound {
			logg.C.Infof("= ren stone:%s area:%d libertyArea:%d adjacentColor:%s\n", ren.stone, ren.GetArea(), ren.GetLibertyArea(), ren.adjacentColor)
			logg.J.Infow("output ren", "color", ren.stone, "area", ren.GetArea(), "libertyArea", ren.GetLibertyArea(), "adjacentColor", ren.adjacentColor)
			return true
		}

		logg.C.Infof("? not found ren coord:%s%\n", coord)
		logg.J.Infow("error not found ren", "coord", coord)
		return false

	// この上にコマンドを挟んでいく
	// -------------------------


// ...略...
```

### Step [O22o2o6o0] 動作確認

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

## Step [O22o3o0] 相手の眼に石を置くことの禁止 - OpponentEye

囲碁のルールでは、相手の眼へは石を置けない。これを判定する  

とりあえず、このルールへ直訳で短い名前を付ける。 仮に `OpponentEye` とでも呼ぶことにする  

このルールは、あとで出てくる `Captured` のルールよりは優先度が低いとする  

### Step [O22o3o1o0] ファイル編集 - kernel.go

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
	// [O22o3o1o0] 相手の眼に石を置こうとした
	var onOpponentEye = func() bool {
		logg.C.Infof("? opponent_eye my_stone:%s point:%s\n", stone, k.Board.GetCodeFromPoint(point))
		logg.J.Infow("error opponent_eye", "my_stone", stone, "point", k.Board.GetCodeFromPoint(point))
		return false
	}

//	var isOk = k.Play(stone, point, logg,
//		// [O22o1o2o0] ,onMasonry
//		onMasonry,
		// [O22o3o1o0] ,onOpponentEye
		onOpponentEye//)
//
//	if isOk {
//		logg.C.Info("=\n")
//		logg.J.Infow("ok")
//	}
// }

// func (k *Kernel) Play(stoneA Stone, pointB Point, logg *Logger,
	// // [O22o1o2o0] onMasonry
	// onMasonry func() bool,
	// [O22o3o1o0] onOpponentEye
	onOpponentEye func() bool//) bool {

	// ...略...
	// // [O22o1o2o0]
	// if k.IsMasonryError(stone, point) {
	//	return onMasonry()
	// }

	// [O22o3o1o0]
	var renC, isFound = k.GetLiberty(pointB)
	if isFound && renC.GetArea() == 1 { // 石Aを置いた交点を含む連Cについて、連Cの面積が1である（眼）
		if stoneA.GetColor() == renC.adjacentColor.GetOpponent() {
			// かつ、連Cに隣接する連の色が、石Aのちょうど反対側の色であったなら、
			// 相手の眼に石を置こうとしたとみなす
			return onOpponentEye()
		}
	}

	// ...略...
	// k.Board.cells[point] = stone
	// return true
```

### Step [O22o3o2o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
board_set file data/board1.txt
play white C3
```

Output > Console:  

```plaintext
[2022-09-17 00:41:29]   # play white C3
[2022-09-17 00:41:29]   ? opponent_eye my_stone:o point:C3
```

## Step [O22o4o0] 自分の眼に石を置くことの任意の禁止

囲碁のルール上可能だが、明らかに損な手は、プレイアウトから除外したい。  
対局時には許可し、プレイアウト時には禁止するよう、選択できるようにする  

### Step [O22o4o1o0] ファイル編集 - play.go

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

	// CanNotPutOnMyEye - [O22o4o1o0] 自分の眼に石を置くことはできません
	CanNotPutOnMyEye bool
// }
// ...略...

// func (k *Kernel) DoPlay(command string, logg *Logger) {

	// ...略...
	// [O22o4o1o0] 自分の眼に石を置こうとした
	var onForbiddenMyEye = func() bool {
		logg.C.Infof("? my_eye my_stone:%s point:%s\n", stone, k.Board.GetCodeFromPoint(point))
		logg.J.Infow("error my_eye", "my_stone", stone, "point", k.Board.GetCodeFromPoint(point))
		return false
	}

//	var isOk = k.Play(stone, point, logg,
//		// [O22o1o2o0] ,onMasonry
//		onMasonry,
//		// [O22o3o1o0] ,onOpponentEye
//		onOpponentEye,
		// [O22o4o1o0] ,onForbiddenMyEye
		onForbiddenMyEye//)
//
//	if isOk {
//		logg.C.Info("=\n")
//		logg.J.Infow("ok")
//	}
// }

// func (k *Kernel) Play(stoneA Stone, pointB Point, logg *Logger,
	// // [O22o1o2o0] onMasonry
	// onMasonry func() bool,
	// [O22o3o1o0] onOpponentEye
	onOpponentEye func() bool,
	// [O22o4o1o0]
	onForbiddenMyEye func() bool//) bool {

	// ...略...
	// // [O22o3o1o0]
	// var renC, isFound = k.GetLiberty(pointB)
	// if isFound && renC.GetArea() == 1 { // 石Aを置いた交点を含む連Cについて、連Cの面積が1である（眼）
	// 	if stoneA.GetColor() == renC.adjacentColor.GetOpponent() {
			// かつ、連Cに隣接する連の色が、石Aのちょうど反対側の色であったなら、
			// 相手の眼に石を置こうとしたとみなす
	// 		return onOpponentEye()

		} else if k.CanNotPutOnMyEye && stoneA.GetColor() == renC.adjacentColor {
			// [O22o4o1o0]
			// かつ、連Cに隣接する連の色が、石Aの色であったなら、
			// 自分の眼に石を置こうとしたとみなす
			return onForbiddenMyEye()

	// }

	// ...略...
	// k.Board.cells[point] = stone
	// return true
```

### Step [O22o4o2o_1o0] ファイル編集 - kernel.go

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
	case "can_not_put_on_my_eye": // [O22o4o2o_1o0]
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

### Step [O22o4o2o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
board_set file data/board1.txt
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
board_set file data/board1.txt
can_not_put_on_my_eye set true
play black C3
```

Output > Console:  

```plaintext
[2022-09-17 09:11:48]   # play black C3
[2022-09-17 09:11:48]   ? my_eye my_stone:x point:C3
```

## Step [O22o5o0] 任意の連の打ち上げ - RemoveRen 関数作成

### Step [O22o5o1o0] ファイル編集 - kernel_facade.go

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
// RemoveRen - 石の連を打ち上げます
func (k *Kernel) RemoveRen(ren *Ren) {
	// 空点をセット
	var setLocation = func(i int, location Point) {
		k.Board.SetStoneAt(location, Space)
	}

	// 場所毎に
	ren.ForeachLocation(setLocation)
}
```

### Step [O22o5o2o0] コマンド実装 - kernel.go ファイル

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
	case "remove_ren": // [O22o5o2o0]
		// Example: `remove_ren B2`
		var coord = tokens[1]
		var point = k.Board.GetPointFromCode(coord)
		var ren, isFound = k.GetLiberty(point)
		if isFound {
			k.RemoveRen(ren)
			logg.C.Infof("=\n")
			logg.J.Infow("ok")
			return true
		}

		logg.C.Infof("? not found ren coord:%s%\n", coord)
		logg.J.Infow("error not found ren", "coord", coord)
		return false

	// ...略...
	// この上にコマンドを挟んでいく
	// -------------------------
	// ...略...
```

### Step [O22o5o3o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
board_set file data/board1.txt
remove_ren B2
```

Output > Console:  

```plaintext
[2022-09-17 12:17:02]   # remove_ren B2
[2022-09-17 12:17:02]   =
```

## Step [O22o6o0] 死に石の連の打ち上げ - Captured

石Aを盤上に置いて指を離したばかりの盤面とする。  
石Aに隣接する相手の石の連のうち、呼吸点が０のものは打ち上げる。  

このとき、 `OpponentEye` のルールと相反することがある

### Step [O22o6o1o0] ファイル編集 - play.go

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

	// [O22o6o1o0] Captured ルール
	var isExists4rensToRemove = false
	var o4rensToRemove [4]*Ren
	var isChecked4rensToRemove = false

	// ...略...
	// [O22o3o1o0]
	// var renC, isFound = k.GetLiberty(pointB)
	// if isFound && renC.GetArea() == 1 { // 石Aを置いた交点を含む連Cについて、連Cの面積が1である（眼）
		// if stoneA.GetColor() == renC.adjacentColor.GetOpponent() {

			// * 以下を追加
			// [O22o6o1o0] 打ちあげる死に石の連を取得
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

		// } else if k.CanNotPutOnMyEye && stoneA.GetColor() == renC.adjacentColor {
			// ...略...
		// }
	// }

	// ...略...
	// 石を置く
	// k.Board.SetStoneAt(pointB, stoneA)

	// * 以下を追加
	// [O22o6o1o0] 打ちあげる死に石の連を取得
	if !isChecked4rensToRemove {
		isExists4rensToRemove, o4rensToRemove = k.GetRenToCapture(pointB)
	}

	// [O22o6o1o0] 死に石を打ちあげる
	if isExists4rensToRemove {
		for dir := 0; dir < 4; dir++ {
			var ren = o4rensToRemove[dir]
			if ren != nil {
				k.RemoveRen(ren)
			}
		}
	}

	// return true
// }

// GetRenToCapture - 現在、着手後の盤面とする。打ち上げられる石の連を返却
//
// Returns
// -------
// isExists : bool
// renToRemove : [4]*Ren
// 隣接する東、北、西、南にある石を含む連
func (k *Kernel) GetRenToCapture(placePlay Point) (bool, [4]*Ren) {
	// [O22o6o1o0]
	var isExists bool
	var rensToRemove [4]*Ren
	var renIds = [4]Point{math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt}

	var setAdjacentPoint = func(dir int, adjacentP Point) {
		var adjacentR, isFound = k.GetLiberty(adjacentP)
		if isFound {
			// 同じ連を数え上げるのを防止する
			var renId = adjacentR.GetMinimumLocation()
			for i := 0; i < dir; i++ {
				if renIds[i] == renId { // Idが既存
					return
				}
			}

			// 取れる石を見つけた
			if adjacentR.GetLibertyArea() < 1 {
				isExists = true
				rensToRemove[dir] = adjacentR
			}
		}
	}

	// 隣接する４方向
	k.Board.ForeachNeumannNeighborhood(placePlay, setAdjacentPoint)

	return isExists, rensToRemove
}
```

### Step [O22o6o2o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
board_set file data/board2.txt
play black D4
```

Output > Console:  

```plaintext
[2022-09-17 14:35:58]   # play black D4
[2022-09-17 14:35:58]   =
```

## Step [O22o7o0] コウの禁止 - Ko

自分が１手前に置いたところに２手続けて置けない

### Step [O22o7o1o0] ファイル編集 - record_item.go

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
 	│	├── 📄 play.go
 	│	├── 📄 point.go
👉	│	├── 📄 record_item.go
	│	├── 📄 record.go
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
// ...略...
// type RecordItem struct {
	// ...略...

	// [O22o7o1o0] コウの位置
	ko Point
// }
// ...略...
```

### Step [O22o7o1o0] ファイル編集 - record.go

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
 	│	├── 📄 play.go
 	│	├── 📄 point.go
	│	├── 📄 record_item.go
👉	│	├── 📄 record.go
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
// ...略...
// Push - 末尾に追加
// func (r *Record) Push(placePlay Point,
	// [O22o7o1o0] コウの位置
	ko Point//) {

	// var item = r.items[r.posNum]
	// item.placePlay = placePlay

	// [O22o7o1o0] コウの位置
	item.ko = ko

	// r.posNum++
// }
```

### Step [O22o7o2o0] ファイル編集 - play.go

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	├── 📄 board1.txt
 	│	├── 📄 board2.txt
 	│	└── 📄 board3.txt
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
 	│	├── 📄 record_item.go
 	│	├── 📄 record.go
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
// func (k *Kernel) DoPlay(command string, logg *Logger) {

	// ...略...
	// [O22o4o1o0] 自分の眼に石を置こうとした
	// var onForbiddenMyEye = func() bool {
		// ...略...
	// }

	// [O22o7o2o0] コウに石を置こうとした
	var onKo = func() bool {
		logg.C.Infof("? ko my_stone:%s point:%s\n", stone, k.Board.GetCodeFromPoint(point))
		logg.J.Infow("error ko", "my_stone", stone, "point", k.Board.GetCodeFromPoint(point))
		return false
	}

	// var isOk = k.Play(stone, point, logg,
		// // [O22o1o2o0] ,onMasonry
		// onMasonry,
		// // [O22o3o1o0] ,onOpponentEye
		// onOpponentEye,
		// // [O22o4o1o0] ,onForbiddenMyEye
		// onForbiddenMyEye,
		// // [O22o7o2o0] ,onKo
		onKo//)
	// ...略...

// }
// ...略...

// func (k *Kernel) Play(stoneA Stone, pointB Point, logg *Logger,
	// // [O22o1o2o0] onMasonry
	// onMasonry func() bool,
	// // [O22o3o1o0] onOpponentEye
	// onOpponentEye func() bool,
	// // [O22o4o1o0] onForbiddenMyEye
	// onForbiddenMyEye func() bool,
	// [O22o7o2o0] onKo
	onKo func() bool//) bool {

	// [O22o1o2o0]
	// if k.IsMasonryError(stoneA, pointB) {
	// 	return onMasonry()
	// }
	// ...略...

	// [O22o7o2o0] コウの判定
	if k.Record.IsKo(pointB) {
		return onKo()
	}
	// ...略...

	// * 以下を追加
	// [O22o7o2o0] コウの判定
	var capturedCount = 0 // アゲハマ

//	// [O22o6o1o0] 死に石を打ちあげる
//	if isExists4rensToRemove {
//		for dir := 0; dir < 4; dir++ {
//			var ren = o4rensToRemove[dir]
//			if ren != nil {
//				k.RemoveRen(ren)

				// * 以下を追加
				// [O22o7o2o0] コウの判定
				capturedCount += ren.GetArea()

//			}
//		}
//	}

	// [O22o7o2o0] コウの判定
	var ko = Point(0)
	if capturedCount == 1 {
		ko = pointB
	}

	// 棋譜に追加
	//k.Record.Push(pointB,
		// * 以下を追加
		// [O22o7o2o0] コウの判定
		ko//)

```

### Step [O22o7o4o0] コマンド編集 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	├── 📄 board1.txt
 	│	├── 📄 board2.txt
 	│	└── 📄 board3.txt
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
 	│	├── 📄 record_item.go
 	│	├── 📄 record.go
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
	// case "record": // [O12o__11o_5o0]
		// ...略...

		// var setPoint = func(i int, item *RecordItem) {
			// var positionNth = i + geta // 基数を序数に変換
			// var coord = k.Board.GetCodeFromPoint(item.placePlay)

			// * 以下を削除
			// sb.WriteString(fmt.Sprintf("[%d]%s ", positionNth, coord))

			// * 以下を追加
			// [O22o7o4o0] コウを追加
			var koStr string
			if item.ko == Point(0) {
				koStr = ""
			} else {
				koStr = fmt.Sprintf("(%s)", k.Board.GetCodeFromPoint(item.ko))
			}
			sb.WriteString(fmt.Sprintf("[%d]%s%s ", positionNth, coord, koStr))
		// }
		// ...略...

	// ...略...
	// この上にコマンドを挟んでいく
	// -------------------------
	// ...略...
```

### Step [O22o7o4o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
board_set file data/board3.txt
play black D3
play white C3
play black D3
```

Output > Console:  

```plaintext
[2022-09-17 22:39:55]   # board_set file data/board3.txt
[2022-09-17 22:39:55]   =

[2022-09-17 22:39:55]   # play black D3
[2022-09-17 22:39:55]   =

[2022-09-17 22:39:55]   # play white C3
[2022-09-17 22:39:55]   =

[2022-09-17 22:39:55]   # play black D3
[2022-09-17 22:39:55]   ? ko my_stone:x point:D3
```

# Step [O23o_1o0] データファイル作成 - data/board4.txt ファイル

あとで使うファイルを先に作成する  

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	├── 📄 board1.txt
 	│	├── 📄 board2.txt
 	│	├── 📄 board3.txt
👉 	│	└── 📄 board4.txt
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
 	│	├── 📄 play.go
 	│	├── 📄 point.go
 	│	├── 📄 record_item.go
 	│	├── 📄 record.go
 	│	├── 📄 ren.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
	└── 📄 main.go
```

```plaintext
     A B C D E F G H J K L M N O P Q R S T
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
 1 + x . x x x . . . . . . . . . x x x . x +
 2 + . . . . . . . . . . . . . . . . . . . +
 3 + x . x x x x x . o o o . . . . . . . x +
 4 + x . x . . . x . o o o x x x . . . . x +
 5 + x . x . x . x . o o o x o x . . . . x +
 6 + . . x x x . x . . . . x . x . . . . . +
 7 + . . . . . . x . . . . x x x . . . . . +
 8 + . . x x x x x . . . . . . o o o o o . +
 9 + . . . . . . . . . . . . . o . o . o . +
10 + o o o o o o o o o o o o o o o o o o o +
11 + . . . . . . . x . . x . . . . . . . . +
12 + . . . . . o . x x . x . x x x x x . . +
13 + . . . . o x o . x x x . x . . . . . . +
14 + . . . o x . x o . . . . x . x x x . . +
15 + x . o x . . . x o . . . x . x . x . x +
16 + x . . o x . x o . . . . x . . . x . x +
17 + x . . . o x o . . . . . x x x x x . x +
18 + . . . . . o . . . . . . . . . . . . . +
19 + x . x x x . . . . . . . . . x x x . x +
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
```

# Step [O23o_2o0] 盤上の連のスキャン

`石を打つ` ことを実装できたなら、呼吸点を調べるアルゴリズムも実装されていて、  
着手点に隣接する上下左右にある石の連の認識はできているはずだ  

次は、盤面全体に点在する連を認識できるか試したい  

## Step [O23o_2o1o0] ファイル編集 - kernel_facade.go

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	├── 📄 board1.txt
 	│	├── 📄 board2.txt
 	│	├── 📄 board3.txt
 	│	└── 📄 board4.txt
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
 	│	├── 📄 record_item.go
 	│	├── 📄 record.go
 	│	├── 📄 ren.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
	└── 📄 main.go
```

👇 がんばって以下を追加してほしい  

```go
// ...略...
// FindAllRens - [O23o_2o1o0] 盤上の全ての連を見つけます
// * 見つけた連は、連データベースへ入れます
func (k *Kernel) FindAllRens() {
	// チェックボードの初期化
	k.CheckBoard.Init(k.Board.GetWidth(), k.Board.GetHeight())

	var maxPosNthFigure = k.Record.GetMaxPosNthFigure()

	var setLocation = func(location Point) {
		var ren, isFound = k.findRen(location)
		if isFound {
			k.renDb.RegisterRen(maxPosNthFigure, k.Record.posNum, ren)
		}
	}
	// 盤上の枠の内側をスキャン。筋、段の順
	k.Board.ForeachPayloadLocationOrderByYx(setLocation)
}
// ...略...
```

## Step [O23o_2o2o0] コマンド編集 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	├── 📄 board1.txt
 	│	├── 📄 board2.txt
 	│	├── 📄 board3.txt
 	│	└── 📄 board4.txt
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
 	│	├── 📄 record_item.go
 	│	├── 📄 record.go
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
	case "find_all_rens": // [O23o_2o2o0]
		// Example: `find_all_rens`
		k.FindAllRens()
		logg.C.Infof("=\n")
		logg.J.Infow("ok")
		return true

	// ...略...
	// この上にコマンドを挟んでいく
	// -------------------------
	// ...略...
```

## Step [O23o_2o3o_1o0] ファイル編集 - board_set.go

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	├── 📄 board1.txt
 	│	├── 📄 board2.txt
 	│	├── 📄 board3.txt
 	│	└── 📄 board4.txt
	├── 📂 kernel
	│	├── 📂 play_rule
	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
👉 	│	├── 📄 board_set.go
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
 	│	├── 📄 play.go
 	│	├── 📄 point.go
 	│	├── 📄 record_item.go
 	│	├── 📄 record.go
 	│	├── 📄 ren.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
	└── 📄 main.go
```

👇 がんばって以下を追加してほしい  

```go
// ...略...
// func (k *Kernel) DoSetBoard(command string, logg *Logger) {
	// if tokens[1] == "file" {
		// ...略...
		// if int(i) != size {
		// 	logg.C.Infof("? not enough size i:%d size:%d\n", i, size)
		// 	logg.J.Infow("error not enough size", "i", i, "size", size)
		// 	return
		// }

		// * 以下を追加
		// [O23o_2o3o_1o0] 連データベース初期化
		k.renDb.Init(k.Board.GetWidth(), k.Board.GetHeight())
		k.FindAllRens()
	// }
// }
// ...略...
```

## Step [O23o_2o3o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
board_set file data/board4.txt
rendb_dump
find_all_rens
rendb_dump
rendb_save data/ren_db_board4_temp.json
```

Output > Console:  

```plaintext
[2022-09-18 23:58:02]   # board_set file data/board4.txt
[2022-09-18 23:58:02]   =

[2022-09-18 23:58:51]   # find_all_rens
[2022-09-18 23:58:51]   =
```

# Step [O23o_2o4o0] TODO 出力ファイルの内容の盤表示

# Step [O23o_2o4o1o0] ファイル編集 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	├── 📄 board1.txt
 	│	├── 📄 board2.txt
 	│	├── 📄 board3.txt
 	│	└── 📄 board4.txt
	├── 📂 kernel
	│	├── 📂 play_rule
	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
 	│	├── 📄 board_set.go
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
 	│	├── 📄 record_item.go
 	│	├── 📄 record.go
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

TODO 📄 `ren_db1_temp.json` ファイルを読み込む  

TODO 連データベースには「何手目」の「どの連」という形で項目がいくつもある。その全項目をコンソールへ盤表示  

# Step [O23o_2o4o2o0]

TODO FIXME rendb_dump の結果が全部空っぽ？

📄 .vscode/launch.json:  

```json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}",
            "args" : ["-d=true"]
        }
    ]
}
```

📄 debug.input.txt:  

```plaintext
rendb_load data/ren_db_board4.json
rendb_dump
```

```plaintext
rendb_load data/ren_db_board4.json
find_all_rens
rendb_dump
```

# Step [O23o_3o0] TODO 石を打った後の連の再スキャン

# Step [O23o0] 打った石のアンドゥ - Undo

打った石をやっぱり止める、一手戻す、ということは、石を打つよりも実装がむずかしい  

### Step [O23o1o0] ファイル作成 - play_undo.go

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	├── 📄 board1.txt
 	│	├── 📄 board2.txt
 	│	└── 📄 board3.txt
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
👉 	│	├── 📄 play_undo.go
 	│	├── 📄 play.go
 	│	├── 📄 point.go
 	│	├── 📄 record_item.go
 	│	├── 📄 record.go
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
// BOF [O23o1o0]

package kernel

// DoUndoPlay - 石を打ったのを戻す
//
// * `command` - Example: `undo`
// ........................----
// ........................0
func (k *Kernel) DoUndoPlay(command string, logg *Logger) {
	k.UndoPlay()
}

// UndoPlay - 石を打ったのを戻す
//
// Returns
// -------
// isOk : bool
//
//	石を置けたら真、置けなかったら偽
func (k *Kernel) UndoPlay() bool {

	// 初期局面から前には戻せない
	if k.Record.GetPositionNumber() < 1 {
		return false
	}

	return false
}

// EOF [O23o1o0]
```

TODO 東、北、西、南に隣接する連を重複して数えないように常に注意すること
TODO 棋譜に取った石を記録する
TODO 取った石を戻す
TODO アンドゥ

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
📖 [Golang String Padding Example](https://golang.cafe/blog/golang-string-padding-example.html)  

### 配列

📖 [スライス操作(要素の追加・削除, ソート, 他のスライスと結合)](https://www.wakuwakubank.com/posts/782-go-slice/)  

### ファイル入出力

📖 [Read a file in Go](https://gosamples.dev/read-file/#:~:text=The%20simplest%20way%20of%20reading,by%20line%20or%20in%20chunks.)  

### JSON

📖 [Goにおけるjsonの扱い方を整理・考察してみた ~ データスキーマを添えて](https://zenn.dev/hsaki/articles/go-convert-json-struct)  

### コレクション

#### ハッシュテーブル

📖 [[Go言語] 初心者必見シリーズ: マップ（Map）](https://qiita.com/wifecooky/items/2ffe41d55c575b2ce5e2)  
📖 [Go言語: マップのキーが存在するかチェックしたい](https://qiita.com/suin/items/4cb1da71237fc55a06ee)  

### ビット演算

📖 [Goのビット演算について](https://www.flyenginer.com/low/go/go%E3%81%AE%E3%83%93%E3%83%83%E3%83%88%E6%BC%94%E7%AE%97%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6.html)  
📖 [How to define bit literals in Go?](https://stackoverflow.com/questions/56605810/how-to-define-bit-literals-in-go)  

### 数学

📖 [[math] 2つの値のうち小さい方を返す (Min)](http://www.openspc2.org/reibun/Go/1.1.1/pkg/math/1043/index.html)  

### デバッギング

📖 [標準入力のあるプログラムを delve でデバッグしたい](https://qiita.com/_natsu_no_yuki_/items/505e74e598d3d6a0cb24)  

.
