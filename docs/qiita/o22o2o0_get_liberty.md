目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O22o2o0] 連の認識と、呼吸点のカウント

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# Step [O22o2o0] 連の認識と、呼吸点のカウント - GetLiberty 関数作成

盤上の座標を指定することで、そこにある `連` の `呼吸点` の数を算出したい  

* `連` - Ren、れん。コンピューター囲碁用語。説明は省略
* `呼吸点` - Liberty、こきゅうてん。コンピューター囲碁用語。説明は省略

👇 呼吸点を数えるために探索すると、一緒に以下のことも行える  

* 連の認識
* 隣接する連の色の取得

このような探索を行う関数に名前を付ける。  
`GetRen` がふさわしいが、慣習を優先して `GetLiberty` と名付けることにする  

## ~~Step [O22o2o1o0]~~

Moved to `[O11o_4o2o1o0]`  

## Step [O22o2o2o0] ファイル作成 - check_board.go

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📂 play_rule
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 o12o__11o1o0_board.go
👉 	│	├── 📄 check_board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
 	│	├── 📄 play.go
 	│	├── 📄 o12o__10o1o0_point.go
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

// Mark - 目印
type Mark uint8

const (
	Mark_BitAllZeros Mark = 0b00000000
	Mark_BitStone    Mark = 0b00000001
	Mark_BitLiberty  Mark = 0b00000010
)

// CheckBoard - チェック盤
type CheckBoard struct {
	// 盤座標
	coordinate BoardCoordinate

	// 長さが可変な盤
	//
	// * 英語で交点は node かも知れないが、表計算でよく使われる cell の方を使う
	cells []Mark
}

// NewDirtyCheckBoard - 新規作成するが、初期化されていない
//
// * このメソッドを呼び出した後に Init 関数を呼び出してほしい
func NewDirtyCheckBoard() *CheckBoard {
	var cb = new(CheckBoard)

	cb.coordinate = BoardCoordinate{}

	return cb
}

// Init - 初期化
func (cb *CheckBoard) Init(newBoardCoordinate BoardCoordinate) {
	// 盤面のサイズが異なるなら、盤面を作り直す
	if cb.coordinate.memoryWidth != newBoardCoordinate.memoryWidth || cb.coordinate.memoryHeight != newBoardCoordinate.memoryHeight {
		cb.coordinate = newBoardCoordinate
		cb.cells = make([]Mark, cb.coordinate.GetMemoryArea())
		return
	}

	// 盤面のクリアー
	for p := Point(0); p < Point(len(cb.cells)); p++ {
		cb.cells[p] = Mark_BitAllZeros
	}
}

// GetAllBitsAt - 指定した交点の目印を取得
func (cb *CheckBoard) GetAllBitsAt(point Point) Mark {
	return cb.cells[point]
}

// SetAllBitsAt - 指定した交点に目印を設定
func (cb *CheckBoard) SetAllBitsAt(point Point, mark Mark) {
	cb.cells[point] = mark
}

// ClearAllBitsAt - フラグを消す
func (cb *CheckBoard) ClearAllBitsAt(point Point) {
	cb.cells[point] = Mark(0)
}

// IsZeroAt - 指定した交点に目印は付いていないか？
func (cb *CheckBoard) IsZeroAt(point Point) bool {
	return cb.cells[point] == Mark_BitAllZeros
}

// Overwrite - 上書き
func (cb *CheckBoard) Overwrite(point Point, mark Mark) {
	cb.cells[point] |= mark
}

// Erase - 消す
func (cb *CheckBoard) Erase(point Point, mark Mark) {
	cb.cells[point] &= ^mark // ^ はビット反転
}

// Contains - 含む
func (cb *CheckBoard) Contains(point Point, mark Mark) bool {
	return cb.cells[point]&mark == mark
}

// EOF [O22o2o2o0]
```

## ~~Step [O22o2o3o_1o0]~~

Removed  

## Step [O22o2o3o0] ファイル編集 - kernel.go

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📂 play_rule
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 o12o__11o1o0_board.go
 	│	├── 📄 check_board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
👉 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
 	│	├── 📄 play.go
 	│	├── 📄 o12o__10o1o0_point.go
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

// func NewDirtyKernel(boardWidht int, boardHeight int) *Kernel {
//	var k = new(Kernel)
//	k.Board = NewBoard(boardWidht, boardHeight)

	// * 以下を追加
	// [O22o2o3o0]
	k.CheckBoard = NewDirtyCheckBoard()

//	return k
// }
```

## Step [O22o2o4o0] ファイル作成 - liberty.go

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📂 play_rule
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 o12o__11o1o0_board.go
 	│	├── 📄 check_board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	├── 📄 kernel.go
👉 	│	├── 📄 liberty.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
 	│	├── 📄 play.go
 	│	├── 📄 o12o__10o1o0_point.go
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
	k.CheckBoard.Init(k.Board.coordinate)

	var libertySearchAlgorithm = NewLibertySearchAlgorithm(k.Board, k.CheckBoard, k.tempRen)

	return libertySearchAlgorithm.findRen(arbitraryPoint)
}

// LibertySearchAlgorithm - 呼吸点探索アルゴリズム
type LibertySearchAlgorithm struct {
	// 盤
	board *Board
	// チェック盤
	checkBoard *CheckBoard
	// tempRen - 呼吸点の探索時に使います
	tempRen *Ren
}

// NewLibertySearchAlgorithm - 新規作成
func NewLibertySearchAlgorithm(board *Board, checkBoard *CheckBoard, tempRen *Ren) *LibertySearchAlgorithm {
	var ls = new(LibertySearchAlgorithm)

	ls.board = board
	ls.checkBoard = checkBoard
	ls.tempRen = tempRen

	return ls
}

// 連の検索
//
// Returns
// -------
// - *Ren is ren or nil
// - bool is found
func (ls *LibertySearchAlgorithm) findRen(arbitraryPoint Point) (*Ren, bool) {
	// 探索済みならスキップ
	if ls.checkBoard.Contains(arbitraryPoint, Mark_BitStone) {
		return nil, false
	}

	// 連の初期化
	ls.tempRen = NewRen(ls.board.GetStoneAt(arbitraryPoint))

	if ls.tempRen.stone == Stone_Space {
		ls.searchSpaceRen(arbitraryPoint)
	} else {
		ls.searchStoneRenRecursive(arbitraryPoint)

		// チェックボードの「呼吸点」チェックのみクリアー
		var eachPoint = func(point Point) {
			ls.checkBoard.Erase(point, Mark_BitLiberty)
		}
		ls.board.coordinate.ForeachCellWithoutWall(eachPoint)
	}

	return ls.tempRen, true
}

// 石の連の探索
//
// * 再帰関数
func (ls *LibertySearchAlgorithm) searchStoneRenRecursive(here Point) {

	// 石のチェック
	ls.checkBoard.Overwrite(here, Mark_BitStone)

	ls.tempRen.AddLocation(here)

	var eachAdjacent = func(dir Cell_4Directions, p Point) {
		// 呼吸点と枠のチェック
		var stone = ls.board.GetStoneAt(p)
		switch stone {
		case Stone_Space: // 空点
			if !ls.checkBoard.Contains(p, Mark_BitLiberty) { // まだチェックしていない呼吸点なら
				ls.checkBoard.Overwrite(p, Mark_BitLiberty)
				ls.tempRen.libertyLocations = append(ls.tempRen.libertyLocations, p) // 呼吸点を追加
			}

			return // あとの処理をスキップ

		case Stone_Wall: // 枠
			return // あとの処理をスキップ
		}

		// 探索済みの石ならスキップ
		if ls.checkBoard.Contains(p, Mark_BitStone) {
			return
		}

		var color = stone.GetColor()
		// 隣接する色、追加
		ls.tempRen.adjacentColor = ls.tempRen.adjacentColor.GetAdded(color)

		if stone == ls.tempRen.stone { // 同じ石
			ls.searchStoneRenRecursive(p) // 再帰
		}
	}

	// 隣接する４方向
	ls.board.ForeachNeumannNeighborhood(here, eachAdjacent)
}

// 空点の連の探索
// - 再帰関数
func (ls *LibertySearchAlgorithm) searchSpaceRen(here Point) {
	ls.checkBoard.Overwrite(here, Mark_BitStone)
	ls.tempRen.AddLocation(here)

	var eachAdjacent = func(dir Cell_4Directions, p Point) {
		// 探索済みならスキップ
		if ls.checkBoard.Contains(p, Mark_BitStone) {
			return
		}

		var stone = ls.board.GetStoneAt(p)
		if stone != Stone_Space { // 空点でなければスキップ
			return
		}

		var color = stone.GetColor()
		// 隣接する色、追加
		ls.tempRen.adjacentColor = ls.tempRen.adjacentColor.GetAdded(color)
		ls.searchSpaceRen(p) // 再帰
	}

	// 隣接する４方向
	ls.board.ForeachNeumannNeighborhood(here, eachAdjacent)
}

// EOF [O22o2o4o0]
```

## Step [O22o2o5o0] コマンド実装 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📂 play_rule
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 o12o__11o1o0_board.go
 	│	├── 📄 check_board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
👉 	│	├── 📄 kernel.go
 	│	├── 📄 liberty.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
 	│	├── 📄 play.go
 	│	├── 📄 o12o__10o1o0_point.go
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
		var point = k.Board.coordinate.GetPointFromGtpMove(coord)
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

## Step [O22o2o6o0] 動作確認

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
