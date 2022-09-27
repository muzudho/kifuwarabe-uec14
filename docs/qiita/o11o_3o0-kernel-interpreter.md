目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O11o_3o0] カーネルのインタープリター ～ Step [O11o_4o0] 石の色定義 ～ Step [O11o_4o2o0] 連の定義 ～ Step [O11o_5o0] 石定義 ～ Step [O12o__10o0] 点定義、またはその盤座標符号定義

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# Step [O11o_3o0] カーネルのインタープリター

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
👉 	│	├── 📄 kernel.go
 	│	└── 📄 logger.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```go
// BOF [O11o_3o0]

package kernel

import "strings"

const geta = 1 // Japanese wooden clogs. Used to convert bases and ordinals.

// Kernel - カーネル
type Kernel struct {
	// Board - 盤
	Board *Board
}

// NewKernel - カーネルの新規作成
func NewKernel(boardWidht int, boardHeight int) *Kernel {
	var k = new(Kernel)
	k.Board = NewBoard(boardWidht, boardHeight)
	return k
}

// Execute - 実行
//
// Returns
// -------
// isHandled : bool
// 処理済なら真
func (k *Kernel) Execute(command string, logg *Logger) bool {

	var tokens = strings.Split(command, " ")
	switch tokens[0] {

	// この下にコマンドを挟んでいく
	// -------------------------

	// この上にコマンドを挟んでいく
	// -------------------------

	default:
	}

	return false
}

// EOF [O11o_3o0]
```

## Step [O11o_3o1o0] 外側のインタープリターから呼び出す

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	└── 📄 logger.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
👉 	└── 📄 main.go
```

👇 がんばって挿入してほしい  

```go
// ...略...

import (
	// ...略...

	"github.com/muzudho/kifuwarabe-uec14/kernel"
)
// ...略...

		// [O12o__11o_4o0] 棋譜の初期化に利用
		var onUnknownTurn = func() kernel.Stone {
			var errMsg = fmt.Sprintf("? unexpected play_first:%s", engineConfig.GetPlayFirst())
			logg.C.Info(errMsg)
			logg.J.Infow("error", "play_first", engineConfig.GetPlayFirst())
			panic(errMsg)
		}

		// [O11o_3o0]
		var kernel1 = kernel.NewKernel(engineConfig.GetBoardSize(), engineConfig.GetBoardSize())
		// 設定ファイルの内容をカーネルへ反映
		kernel1.Board.Init(engineConfig.GetBoardSize(), engineConfig.GetBoardSize())

		/*
		...以下略...
		// [O11o_1o0] コンソール等からの文字列入力
		var scanner = bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			var command = scanner.Text()
			logg.C.Infof("# %s", command)             // 人間向けの出力
			logg.J.Infow("input", "command", command) // コンピューター向けの出力
		...以上略...
		*/

			// * これを追加する
			// [O11o_3o0]
			var isHandled = kernel1.Execute(command, logg)
			if isHandled {
				continue
			}

			/*
			...以下略...
			// [O11o_1o0]
			var tokens = strings.Split(command, " ")
			switch tokens[0] {
			*/
```

* カーネルと、自作の部分で コマンドが被ったなら、カーネルの方を優先する
  * これにより カーネルのアップデートにより 自作のコマンドが避けられることから、アップデート時は動作テストしてほしい

# Step [O11o_4o0] 石の色定義

`石` を定義していないが、先に `石の色` を定義する  

石の色の組み合わせを定義する。  
石の色の組み合わせは以下の４通りある。これらの集合を `Color` と名付けることにする   

* `Color_None` - 隣接する連は１つもない
* `Color_Black` - 黒石の連とだけ隣接する
* `Color_White` - 白石の連とだけ隣接する
* `Color_Mixed` - 黒石と白石の連の両方に隣接する

## Step [O11o_4o1o0] ファイル作成 - color.go

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
👉	│	├── 📄 color.mod
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	└── 📄 logger.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```go
// BOF [O11o_4o1o0]

package kernel

import "fmt"

type Color uint

const (
	Color_None Color = iota
	Color_Black
	Color_White
	Color_Mixed
)

// String - 文字列化
func (c Color) String() string {
	switch c {
	case Color_None:
		return ""
	case Color_Black:
		return "x"
	case Color_White:
		return "o"
	case Color_Mixed:
		return "xo"
	default:
		panic(fmt.Sprintf("unexpected color:%d", int(c)))
	}
}

// GetAdded - 色の加算
func (c1 Color) GetAdded(c2 Color) Color {
	switch c1 {
	case Color_None:
		return c2
	case Color_Black:
		switch c2 {
		case Color_None:
			return Color_Black
		case Color_Black:
			return Color_Black
		case Color_White:
			return Color_Mixed
		case Color_Mixed:
			return Color_Mixed
		default:
			panic(fmt.Sprintf("unexpected my_color:%s adds_color:%s", c1, c2))
		}
	case Color_White:
		switch c2 {
		case Color_None:
			return Color_White
		case Color_Black:
			return Color_Mixed
		case Color_White:
			return Color_White
		case Color_Mixed:
			return Color_Mixed
		default:
			panic(fmt.Sprintf("unexpected my_color:%s adds_color:%s", c1, c2))
		}
	case Color_Mixed:
		return Color_Mixed
	default:
		panic(fmt.Sprintf("unexpected my_color:%s adds_color:%s", c1, c2))
	}
}

// GetOpponent - 色の反転
func (c Color) GetOpponent() Color {
	switch c {
	case Color_None:
		return Color_Mixed
	case Color_Black:
		return Color_White
	case Color_White:
		return Color_Black
	case Color_Mixed:
		return Color_None
	default:
		panic(fmt.Sprintf("unexpected color:%d", int(c)))
	}
}

// EOF [O11o_4o1o0]
```

# Step [O11o_4o2o0] 連の定義

`石` を定義していないが、先に `連` （れん）を定義する。  
`連` とは何かの説明は、ここでは省く  

### Step [O11o_4o2o1o0] ファイル作成 - ren.go

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 color.mod
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
👉 	│	└── 📄 ren.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```go
// BOF [O11o_4o2o1o0]

package kernel

import "math"

// Ren - 連，れん
type Ren struct {
	// Sto - （外部ファイル向け）石
	Sto string `json:"stone"`
	// Loc - （外部ファイル向け）石の盤上の座標符号の空白区切りのリスト
	Loc string `json:"locate"`
	// LibLoc - （外部ファイル向け）呼吸点の盤上の座標符号の空白区切りのリスト
	LibLoc string `json:"liberty"`

	// 石
	stone Stone
	// 隣接する石の色
	adjacentColor Color
	// 要素の石の位置
	locations []Point
	// 呼吸点の位置
	libertyLocations []Point
	// 最小の場所。Idとして利用することを想定
	minimumLocation Point
}

// NewRen - 連を新規作成
//
// Parameters
// ----------
// color - 色
func NewRen(stone Stone) *Ren {
	var r = new(Ren)
	r.stone = stone
	r.adjacentColor = Color_None
	r.minimumLocation = math.MaxInt
	return r
}

// GetArea - 面積。アゲハマの数
func (r *Ren) GetArea() int {
	return len(r.locations)
}

// GetLibertyArea - 呼吸点の面積
func (r *Ren) GetLibertyArea() int {
	return len(r.libertyLocations)
}

// GetStone - 石
func (r *Ren) GetStone() Stone {
	return r.stone
}

// GetAdjacentColor - 隣接する石の色
func (r *Ren) GetAdjacentColor() Color {
	return r.adjacentColor
}

// GetMinimumLocation - 最小の場所。Idとして利用することを想定
func (r *Ren) GetMinimumLocation() Point {
	return r.minimumLocation
}

// AddLocation - 場所の追加
func (r *Ren) AddLocation(location Point) {
	r.locations = append(r.locations, location)

	// 最小の数を更新
	r.minimumLocation = Point(math.Min(float64(r.minimumLocation), float64(location)))
}

// ForeachLocation - 場所毎に
func (r *Ren) ForeachLocation(setLocation func(int, Point)) {
	for i, point := range r.locations {
		setLocation(i, point)
	}
}

// Dump - ダンプ
//
// Example: `22 23 24 25`
func (r *Ren) Dump() string {
	var convertLocation = func(location Point) string {
		return fmt.Sprintf("%d", location)
	}
	var tokens = r.createCoordBelt(r.locations, convertLocation)
	return strings.Join(tokens, " ")
}

// 文字列の配列を作成
// Example: {`22`, `23` `24`, `25`}
func (r *Ren) createCoordBelt(locations []Point, convertLocation func(Point) string) []string {
	var tokens []string

	// 全ての要素
	for _, location := range locations {
		var token = convertLocation(location)
		tokens = append(tokens, token)
	}

	return tokens
}

// RefreshToExternalFile - 外部ファイルに出力されてもいいように内部状態を整形します
func (r *Ren) RefreshToExternalFile(convertLocation func(Point) string) {
	// Examples: `.`, `x`, `o`, `+`
	r.Sto = r.stone.String()

	{
		// Example: `A1 B2 C3 D4`
		var tokens = r.createCoordBelt(r.locations, convertLocation)
		// sort.Strings(tokens) // 辞書順ソート - 走査方向が変わってしまうので止めた
		r.Loc = strings.Join(tokens, " ")
	}
	{
		var tokens = r.createCoordBelt(r.libertyLocations, convertLocation)
		r.LibLoc = strings.Join(tokens, " ")
	}
}

// EOF [O11o_4o2o1o0]
```

# Step [O11o_5o0] 石定義

## Step [O11o0] ファイル作成 - stone.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
👉 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
 	└── 📄 main.go
```

```go
// BOF [O11o0]

package kernel

import "fmt"

type Stone uint

const (
	Space Stone = iota
	Black
	White
	Wall
)

// GetStoneFromName - 文字列の名前を与えると、Stone値を返します
//
// Returns
// -------
// isOk : bool
// stone : Stone
func GetStoneFromName(stoneName string, getDefaultStone func() (bool, Stone)) (bool, Stone) {
	switch stoneName {
	case "space":
		return true, Space
	case "black":
		return true, Black
	case "white":
		return true, White
	case "wall":
		return true, Wall
	default:
		return getDefaultStone()
	}
}

// GetStoneOrDefaultFromTurn - black または white を与えると、Stone値を返します
//
// Returns
// -------
// stone : Stone
func GetStoneOrDefaultFromTurn(stoneName string, getDefaultStone func() Stone) Stone {
	switch stoneName {
	case "black":
		return Black
	case "white":
		return White
	default:
		return getDefaultStone()
	}
}

// GetStoneFromChar - １文字与えると、Stone値を返します
//
// Returns
// -------
// isOk : bool
// stone : Stone
func GetStoneFromChar(stoneChar string, getDefaultStone func() (bool, Stone)) (bool, Stone) {
	switch stoneChar {
	case ".":
		return true, Space
	case "x":
		return true, Black
	case "o":
		return true, White
	case "+":
		return true, Wall
	default:
		return getDefaultStone()
	}
}

// String - 文字列化
func (s Stone) String() string {
	switch s {
	case Space:
		return "."
	case Black:
		return "x"
	case White:
		return "o"
	case Wall:
		return "+"
	default:
		panic(fmt.Sprintf("%d", int(s)))
	}
}

// GetColor - 色の取得
func (s Stone) GetColor() Color {
	switch s {
	case Space:
		return Color_None
	case Black:
		return Color_Black
	case White:
		return Color_White
	case Wall:
		return Color_None
	default:
		panic(fmt.Sprintf("%d", int(s)))
	}
}

// EOF [O11o0]
```

## Step [O11o1o0] リモートリポジトリにプッシュ

がんばって git などを使い、 `github.com/muzudho/kifuwarabe-uec14/kernel` モジュールの各パッケージのソースを  
リモートリポジトリにプッシュしてほしい  

## Step [O11o2o0] カレントディレクトリーを移動

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd kernel
```

## Step [O11o3o0] tidy

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go get -u
go mod tidy
```

## Step [O11o4o0] カレントディレクトリーを戻す

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

```shell
cd ..
```

## Step [O11o5o0] tidy

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go mod tidy
```

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
👉  │	├── 📄 point.go
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

import (
	"fmt"
	"strconv"
)

// Point - 交点の座標。いわゆる配列のインデックス。壁を含む盤の左上を 0 とします
type Point int

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
	return strconv.Itoa(y + 1)
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

// EOF [O12o__10o1o0]
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
