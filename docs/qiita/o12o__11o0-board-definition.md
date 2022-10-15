目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O12o__11o0] 盤定義（土台）

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# Step [O12o__11o0] 盤定義（土台）

これから盤を作っていく前に、土台を作る  

## Step [O12o__11o1o0] ファイル作成 - board.go

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
👉  │	├── 📄 o12o__11o1o0_board.go
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

// Board - 盤
type Board struct {
	// ゲームルール
	gameRule GameRule
	// 盤座標
	coordinate BoardCoordinate

	// 交点
	//
	// * 英語で交点は node かも知れないが、表計算でよく使われる cell の方を使う
	cells []Stone
}

// NewBoard - 新規作成
func NewBoard(gameRule GameRule, boardWidht int, boardHeight int) *Board {
	var b = new(Board)

	// 設定ファイルから読込むので動的設定
	b.gameRule = gameRule

	// 枠の分、２つ増える
	var memoryBoardWidth = boardWidht + 2
	var memoryBoardHeight = boardHeight + 2

	b.coordinate = BoardCoordinate{
		memoryBoardWidth,
		memoryBoardHeight,
		// ４方向（東、北、西、南）の番地への相対インデックス
		[4]Point{1, Point(-memoryBoardWidth), -1, Point(memoryBoardWidth)}}

	// 盤のサイズ指定と、盤面の初期化
	b.resize(boardWidht, boardHeight)

	return b
}

// GetGameRule - ゲームルール取得
func (b *Board) GetGameRule() *GameRule {
	return &b.gameRule
}

// SetGameRule - ゲームルール設定
func (b *Board) SetGameRule(gameRule *GameRule) {
	b.gameRule = *gameRule
}

// GetCoordinate - 盤座標取得
func (b *Board) GetCoordinate() *BoardCoordinate {
	return &b.coordinate
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

// IsEmpty - 指定の交点は空点か？
func (b *Board) IsSpaceAt(point Point) bool {
	return b.GetStoneAt(point) == Stone_Space
}

func getXyFromPointOnBoard(memoryWidth int, point Point) (int, int) {
	var p = int(point)
	return p % memoryWidth, p / memoryWidth
}

// サイズ変更
func (b *Board) resize(width int, height int) {
	b.coordinate.memoryWidth = width + bothSidesWallThickness
	b.coordinate.memoryHeight = height + bothSidesWallThickness
	b.cells = make([]Stone, b.coordinate.GetMemoryArea())

	// ４方向（東、北、西、南）の番地への相対インデックス
	b.coordinate.cell4Directions = [4]Point{1, Point(-b.coordinate.GetMemoryWidth()), -1, Point(b.coordinate.GetMemoryWidth())}
}

// EOF [O12o__11o1o0]
```

## Step [O12o__11o2o0] コマンド実装 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
  	│	├── 📄 o12o__11o1o0_board.go
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

		var point = k.Position.Board.coordinate.GetPointFromXy(x, y)
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
