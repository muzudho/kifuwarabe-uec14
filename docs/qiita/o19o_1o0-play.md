目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O19o_1o0] 石を打つ

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# Step [O19o_1o0] 石を打つ - play コマンド

## Step [O19o0] ファイル作成 - play.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 o12o__11o1o0_board.go
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
👉 	│	├── 📄 play.go
 	│	├── 📄 o12o__10o1o0_point.go
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
	var point = k.Board.coordinate.GetPointFromGtpMove(coord)

	// [O22o1o2o0] 石（または枠）の上に石を置こうとした
	var onMasonry = func() bool {
		logg.C.Infof("? masonry my_stone:%s point:%s\n", stone, k.Board.coordinate.GetGtpMoveFromPoint(point))
		logg.J.Infow("error masonry", "my_stone", stone, "point", k.Board.coordinate.GetGtpMoveFromPoint(point))
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
  	│	├── 📄 o12o__11o1o0_board.go
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
  	│	├── 📄 o12o__11o1o0_board.go
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

## Step [O22o1o0] 空点以外のところ（石または枠の上）に石を置くことの禁止 - IsMasonryError関数作成

とりあえず、 `石または枠の上に石を置く行為` に `Masonry` （メイスンリー）という名前を付ける。  
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
  	│	├── 📄 o12o__11o1o0_board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
👉 	│	├── 📄 masonry.go
 	│	├── 📄 play.go
 	│	├── 📄 o12o__10o1o0_point.go
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
  	│	├── 📄 o12o__11o1o0_board.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
👉 	│	├── 📄 play.go
 	│	├── 📄 o12o__10o1o0_point.go
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
	// var point = k.Board.GetPointFromGtpMove(tokens[2])

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
	Mark_BitStone   Mark = 0b00000001
	Mark_BitLiberty Mark = 0b00000010
)

// CheckBoard - チェック盤
type CheckBoard struct {
	// 盤座標
	coordinate BoardCoordinate

	// 交点
	//
	// * 英語で交点は node かも知れないが、表計算でよく使われる cell の方を使う
	cells []Mark
}

// NewCheckBoard - 新規作成
//
// * このメソッドを呼び出した後に Init 関数を呼び出してほしい
func NewCheckBoard(coordinate BoardCoordinate) *CheckBoard {
	var cb = new(CheckBoard)

	cb.coordinate = coordinate

	return cb
}

// Init - 盤面初期化
func (cb *CheckBoard) Init(newBoardCoordinate BoardCoordinate) {
	// 盤面のサイズが異なるなら、盤面を作り直す
	if cb.coordinate.memoryWidth != newBoardCoordinate.memoryWidth || cb.coordinate.memoryHeight != newBoardCoordinate.memoryHeight {
		cb.coordinate = newBoardCoordinate
		cb.cells = make([]Mark, cb.coordinate.GetMemoryArea())
		return
	}

	// 盤面のクリアー
	for i := 0; i < len(cb.cells); i++ {
		cb.cells[i] = 0
	}
}

// CheckStone - 石をチェックした
func (cb *CheckBoard) CheckStone(p Point) {
	cb.cells[p] |= Mark_BitStone
}

// IsChecked - 石はチェックされているか？
func (cb *CheckBoard) IsStoneChecked(p Point) bool {
	return cb.cells[p]&Mark_BitStone == Mark_BitStone
}

// CheckLiberty - 呼吸点をチェックした
func (cb *CheckBoard) CheckLiberty(p Point) {
	cb.cells[p] |= Mark_BitLiberty
}

// UncheckLiberty - 呼吸点のチェックを外した
func (cb *CheckBoard) UncheckLiberty(p Point) {
	cb.cells[p] &= ^Mark_BitLiberty // ^ はビット反転
}

// IsLibertyChecked - 呼吸点はチェックされているか？
func (cb *CheckBoard) IsLibertyChecked(p Point) bool {
	return cb.cells[p]&Mark_BitLiberty == Mark_BitLiberty
}

// EOF [O22o2o2o0]
```

### ~~Step [O22o2o3o_1o0]~~

Removed  

### Step [O22o2o3o0] ファイル編集 - kernel.go

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

// func NewKernel(boardWidht int, boardHeight int) *Kernel {
//	var k = new(Kernel)
//	k.Board = NewBoard(boardWidht, boardHeight)

	// * 以下を追加
	// [O22o2o3o0]
	k.CheckBoard = NewCheckBoard(k.Board.coordinate)

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
		// 呼吸点と枠のチェック
		var stone = k.Board.GetStoneAt(p)
		switch stone {
		case Space: // 空点
			if !k.CheckBoard.IsLibertyChecked(p) { // まだチェックしていない呼吸点なら
				k.CheckBoard.CheckLiberty(p)
				k.tempRen.libertyLocations = append(k.tempRen.libertyLocations, p) // 呼吸点を追加
			}

			return // あとの処理をスキップ

		case Wall: // 枠
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
