目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O11o_3o0] カーネルのインタープリター ～ Step [O11o_4o0] 石の色定義 ～ Step [O11o_4o2o0] 連の定義

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
