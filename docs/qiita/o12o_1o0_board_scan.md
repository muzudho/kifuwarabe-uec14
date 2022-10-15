目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O12o_1o0] 盤定義（盤面走査）

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# Step [O12o_1o0] 盤定義（盤面走査）

## Step [O12o0] ファイル作成 - board_area.go

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
👉  │	├── 📄 board_area.go
  	│	├── 📄 o12o__11o1o0_board.go
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
	if b.coordinate.memoryWidth != width || b.coordinate.memoryHeight != height {
		b.resize(width, height)
	}

	// 枠の上辺、下辺を引く
	{
		var y = 0
		var y2 = b.coordinate.memoryHeight - 1
		for x := 0; x < b.coordinate.memoryWidth; x++ {
			var i = b.coordinate.GetPointFromXy(x, y)
			b.cells[i] = Stone_Wall

			i = b.coordinate.GetPointFromXy(x, y2)
			b.cells[i] = Stone_Wall
		}
	}
	// 枠の左辺、右辺を引く
	{
		var x = 0
		var x2 = b.coordinate.memoryWidth - 1
		for y := 1; y < b.coordinate.memoryHeight-1; y++ {
			var i = b.coordinate.GetPointFromXy(x, y)
			b.cells[i] = Stone_Wall

			i = b.coordinate.GetPointFromXy(x2, y)
			b.cells[i] = Stone_Wall
		}
	}
	// 枠の内側を空点で埋める
	{
		var height = b.coordinate.GetHeight()
		var width = b.coordinate.GetWidth()
		for y := 1; y < height; y++ {
			for x := 1; x < width; x++ {
				var i = b.coordinate.GetPointFromXy(x, y)
				b.cells[i] = Stone_Space
			}
		}
	}
}

// ForeachCellWithoutWall - 枠や改行を含めない各セルの番地
func (bc *BoardCoordinate) ForeachCellWithoutWall(setLocation func(Point)) {
	var height = bc.memoryHeight - 1
	var width = bc.memoryWidth - 1
	for y := 1; y < height; y++ {
		for x := 1; x < width; x++ {
			var i = bc.GetPointFromXy(x, y)
			setLocation(i)
		}
	}
}

// ForeachPayloadLocationOrderByYx - 枠や改行を含めない各セルの番地。筋、段の順
func (bc *BoardCoordinate) ForeachPayloadLocationOrderByYx(setLocation func(Point)) {
	var height = bc.memoryHeight - 1
	var width = bc.memoryWidth - 1
	for x := 1; x < width; x++ {
		for y := 1; y < height; y++ {
			var i = bc.GetPointFromXy(x, y)
			setLocation(i)
		}
	}
}

// ForeachNeumannNeighborhood - [O13o__10o0] 隣接する４方向の定義
func (b *Board) ForeachNeumannNeighborhood(here Point, setAdjacent func(int, Point)) {
	// 東、北、西、南
	for dir := 0; dir < 4; dir++ {
		var p = here + b.coordinate.cell4Directions[dir] // 隣接する交点

		// 範囲外チェック
		if p < 0 || b.coordinate.GetMemoryArea() <= int(p) {
			continue
		}

		// 枠チェック
		if b.GetStoneAt(p) == Stone_Wall {
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
