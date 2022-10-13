目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O12o__10o0] 点定義、またはその盤座標符号定義

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# Step [O12o__10o0] 点定義、またはその盤座標符号定義

盤を作る前に、これから盤座標符号を定義する  

## Step [O12o__10o1o0] ファイル作成 - point.go ファイル

👇 以下の既存ファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
👉  │	├── 📄 o12o__10o1o0_point.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
  	└── 📄 main.go
```

```go
// BOF [O12o__10o1o0]

package kernel

// Point - 交点の座標。枠を含む盤の左上を 0 とします
type Point int

// Cell_Pass - パス
const Cell_Pass Point = 0

// IllegalZ - 石が置けない番地の目印として使用。例：UCT計算中に石が置けなかった
const Cell_Illegal Point = -1

// EOF [O12o__10o1o0]
```

## Step [O12o__10o2o_1o0] 盤座標定義

## Step [O12o__10o2o_1o1o0] ファイル作成 - board_coordinate.go ファイル

👇 以下の既存ファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
👉	│	├── 📄 o12o__10o2o_1o1o0_board_coordinate.go
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
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
// BOF [O12o__10o2o_1o1o0]

package kernel

import (
	"fmt"
	"strconv"
)

// 片方の枠の厚み。東、北、西、南
const oneSideWallThickness = 1

// 両側の枠の厚み。南北、または東西
const bothSidesWallThickness = 2

// Cell_4Directions - 東、北、西、南を指す配列のインデックスに対応
type Cell_4Directions int

// 東、北、西、南を指す配列のインデックスに対応
const (
	Cell_East Cell_4Directions = iota
	Cell_North
	Cell_West
	Cell_South
)

// BoardCoordinate - 盤座標
type BoardCoordinate struct {
	// 枠付きの盤の水平一辺の交点の要素数
	memoryWidth int
	// 枠付きの盤の垂直一辺の交点の要素数
	memoryHeight int

	// ４方向（東、北、西、南）への相対番地。2015年講習会サンプル、GoGo とは順序が違います
	cell4Directions [4]Point
}

// GetMemoryWidth - 枠付きの盤の水平一辺の交点数
func (bc *BoardCoordinate) GetMemoryWidth() int {
	return bc.memoryWidth
}

// GetMemoryHeight - 枠付きの盤の垂直一辺の交点数
func (bc *BoardCoordinate) GetMemoryHeight() int {
	return bc.memoryHeight
}

// GetMemoryArea - 枠付き盤の面積
func (bc *BoardCoordinate) GetMemoryArea() int {
	return bc.GetMemoryWidth() * bc.GetMemoryHeight()
}

// GetWidth - 枠無し盤の横幅
func (bc *BoardCoordinate) GetWidth() int {
	// 枠の分、２つ減らす
	return bc.memoryWidth - bothSidesWallThickness
}

// GetHeight - 枠無し盤の縦幅
func (bc *BoardCoordinate) GetHeight() int {
	// 枠の分、２つ減らす
	return bc.memoryHeight - bothSidesWallThickness
}

// GetBoardArea - 枠無し盤の面積
func (bc *BoardCoordinate) GetBoardArea() int {
	return bc.GetWidth() * bc.GetHeight()
}

// GetRelativePointOf - ４方向（東、北、西、南）の先の番地
func (bc *BoardCoordinate) GetRelativePointOf(dir4 Cell_4Directions) Point {
	return bc.cell4Directions[dir4]
}

// GetEastOf - 東
func (bc *BoardCoordinate) GetEastOf(point Point) Point {
	return point + bc.cell4Directions[Cell_East]
}

// GetNorthEastOf - 北東
func (bc *BoardCoordinate) GetNorthEastOf(point Point) Point {
	return point + bc.cell4Directions[Cell_North] + bc.cell4Directions[Cell_East]
}

// GetNorthOf - 北
func (bc *BoardCoordinate) GetNorthOf(point Point) Point {
	return point + bc.cell4Directions[Cell_North]
}

// GetNorthWestOf - 北西
func (bc *BoardCoordinate) GetNorthWestOf(point Point) Point {
	return point + bc.cell4Directions[Cell_North] + bc.cell4Directions[Cell_West]
}

// GetWestOf - 西
func (bc *BoardCoordinate) GetWestOf(point Point) Point {
	return point + bc.cell4Directions[Cell_West]
}

// GetSouthWestOf - 南西
func (bc *BoardCoordinate) GetSouthWestOf(point Point) Point {
	return point + bc.cell4Directions[Cell_South] + bc.cell4Directions[Cell_West]
}

// GetSouthOf - 南
func (bc *BoardCoordinate) GetSouthOf(point Point) Point {
	return point + bc.cell4Directions[Cell_South]
}

// GetSouthEastOf - 南東
func (bc *BoardCoordinate) GetSouthEastOf(point Point) Point {
	return point + bc.cell4Directions[Cell_South] + bc.cell4Directions[Cell_East]
}

// GetPointFromXy - x,y 形式の座標を、 point （配列のインデックス）へ変換します。
// point は枠を含む盤上での座標です
//
// Parameters
// ----------
// x : int
// 枠を含む盤での筋番号。 Example: 19路盤なら0～20
// y : int
// 枠を含む盤での段番号。 Example: 19路盤なら0～20
//
// Returns
// -------
// point : Point
// 配列インデックス
func (bc *BoardCoordinate) GetPointFromXy(x int, y int) Point {
	return Point(y*bc.memoryWidth + x)
}

// GetXyFromPoint - `GetPointFromXy` の逆関数
func (bc *BoardCoordinate) GetXyFromPoint(point Point) (int, int) {
	return getXyFromPointOnBoard(bc.memoryWidth, point)
}

// GetXFromFile - `A` ～ `Z` を 0 ～ 24 へ変換します。 国際囲碁連盟のルールに倣い、筋の符号に `I` は使いません
func GetXFromFile(file string) int {
	// 筋
	var x = file[0] - 'A'
	if file[0] >= 'I' {
		x--
	}
	return int(x)
}

// GetFileFromX - GetXFromFile の逆関数
func GetFileFromX(x int) string {
	// ABCDEFGHI
	// 012345678
	if 7 < x {
		// 'I' を飛ばす
		x++
	}
	// 筋
	return fmt.Sprintf("%c", 'A'+x)
}

// GetYFromRank - '1' ～ '99' を 0 ～ 98 へ変換します
func GetYFromRank(rank string) int {
	// 段
	var y = int(rank[0] - '0')
	if 1 < len(rank) {
		y *= 10
		y += int(rank[1] - '0')
	}
	return y - 1
}

// GetRankFromY - GetYFromRank の逆関数
//
// Parameters
// ----------
// y : int
//
//	0 .. 98
//
// Returns
// -------
// rank : string
//
//	"1" .. "99"
func GetRankFromY(y int) string {
	return strconv.Itoa(getRankFromY(y))GetBoardHeight
}

func getRankFromY(y int) int {
	return y + 1
}

// GetFileFromCode - 座標の符号の筋の部分を抜き出します
//
// * `code` - 座標の符号。 Example: "A7" や "J13"
func GetFileFromCode(code string) string {
	return code[0:1]
}

// GetRankFromCode - 座標の符号の段の部分を抜き出します
//
// * `code` - 座標の符号。 Example: "A7" や "J13"
func GetRankFromCode(code string) string {
	if 2 < len(code) {
		return code[1:3]
	}

	return code[1:2]
}

// EOF [O12o__10o2o_1o1o0]
```

## Step [O12o__10o2o0] コマンド実装 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
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

	case "test_coord": // [O12o__10o2o0]
		// Example: "test_coord A13"
		var point = k.Board.coordinate.GetPointFromGtpMove(tokens[1])
		logg.C.Infof("= %d\n", point)
		logg.J.Infow("output", "point", point)
		return true

	// * アルファベット順になる位置に、以下のケース文を挿入
	case "test_file": // [O12o__10o2o0]
		// Example: "test_file A"
		var file = GetFileFromCode(tokens[1])
		logg.C.Infof("= %s\n", file)
		logg.J.Infow("output", "file", file)
		return true

	case "test_rank": // [O12o__10o2o0]
		// Example: "test_rank 13"
		var rank = GetRankFromCode(tokens[1])
		logg.C.Infof("= %s\n", rank)
		logg.J.Infow("output", "rank", rank)
		return true

	case "test_x": // [O12o__10o2o0]
		// Example: "test_x 18"
		var x, err = strconv.Atoi(tokens[1])
		if err != nil {
			logg.C.Infof("? unexpected x:%s\n", tokens[1])
			logg.J.Infow("error", "x", tokens[1])
			return true
		}
		var file = GetFileFromX(x)
		logg.C.Infof("= %s\n", file)
		logg.J.Infow("output", "file", file)
		return true

	case "test_y": // [O12o__10o2o0]
		// Example: "test_y 18"
		var y, err = strconv.Atoi(tokens[1])
		if err != nil {
			logg.C.Infof("? unexpected y:%s\n", tokens[1])
			logg.J.Infow("error", "y", tokens[1])
			return true
		}
		var rank = GetRankFromY(y)
		logg.C.Infof("= %s\n", rank)
		logg.J.Infow("output", "rank", rank)
		return true

	// この上にコマンドを挟んでいく
	// -------------------------


// ...略...
```

## Step [O12o__10o3o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
test_file A1
```

Output > Console:  

```plaintext
[2022-09-11 23:27:52]   # test_file A1
[2022-09-11 23:27:52]   = A
```

Output > Log > PlainText:  

```plaintext
2022-09-11T23:27:52.547+0900	# test_file A1
2022-09-11T23:27:52.583+0900	= A
```

Output > Log > JSON:  

```json
{"level":"info","ts":"2022-09-11T23:27:52.583+0900","caller":"kifuwarabe-uec14/main.go:61","msg":"input","command":"test_file A1"}
{"level":"info","ts":"2022-09-11T23:27:52.584+0900","caller":"kernel/kernel.go:73","msg":"output","file":"A"}
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
test_rank A1
```

Output > Console:  

```plaintext
[2022-09-11 23:31:11]   # test_rank A1
[2022-09-11 23:31:11]   = 1
```

Output > Log > PlainText:  

```plaintext
2022-09-11T23:31:11.020+0900	# test_rank A1
2022-09-11T23:31:11.020+0900	= 1
```

Output > Log > JSON:  

```json
{"level":"info","ts":"2022-09-11T23:31:11.020+0900","caller":"kifuwarabe-uec14/main.go:61","msg":"input","command":"test_rank A1"}
{"level":"info","ts":"2022-09-11T23:31:11.021+0900","caller":"kernel/kernel.go:80","msg":"output","rank":"1"}
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
test_coord A1
```

Output > Console:  

```plaintext
[2022-09-11 23:32:42]   # test_coord A1
[2022-09-11 23:32:42]   = 22
```

Output > Log > PlainText:  

```plaintext
2022-09-11T23:32:42.228+0900	# test_coord A1
2022-09-11T23:32:42.229+0900	= 22
```

Output > Log > JSON:  

```json
{"level":"info","ts":"2022-09-11T23:32:42.229+0900","caller":"kifuwarabe-uec14/main.go:61","msg":"input","command":"test_coord A1"}
{"level":"info","ts":"2022-09-11T23:32:42.229+0900","caller":"kernel/kernel.go:66","msg":"output","point":22}
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
test_x 18
```

* x は 0 から始まるので、 19列目は 18

Output > Console:  

```plaintext
[2022-09-13 23:53:40]   # test_x 18
[2022-09-13 23:53:40]   = T
```

Output > Log > PlainText:  

```plaintext
2022-09-13T23:53:40.906+0900	# test_x 18
2022-09-13T23:53:40.943+0900	= T
```

Output > Log > JSON:  

```json
{"level":"info","ts":"2022-09-13T23:53:40.943+0900","caller":"kifuwarabe-uec14/main.go:76","msg":"input","command":"test_x 18"}
{"level":"info","ts":"2022-09-13T23:53:40.943+0900","caller":"kernel/kernel.go:115","msg":"output","file":"T"}
```

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
test_y 18
```

* y は 0 から始まるので、 19列目は 18

Output > Console:  

```plaintext
[2022-09-13 23:58:42]   # test_y 18
[2022-09-13 23:58:42]   = 19
```

Output > Log > PlainText:  

```plaintext
2022-09-13T23:58:42.739+0900	# test_y 18
2022-09-13T23:58:42.781+0900	= 19
```

Output > Log > JSON:  

```json
{"level":"info","ts":"2022-09-13T23:58:42.781+0900","caller":"kifuwarabe-uec14/main.go:76","msg":"input","command":"test_y 18"}
{"level":"info","ts":"2022-09-13T23:58:42.782+0900","caller":"kernel/kernel.go:128","msg":"output","rank":"19"}
```
