目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O22o1o0] 空点以外のところ（石または枠の上）に石を置くことの禁止

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# Step [O22o0] 囲碁の石を打つルールの実装

# Step [O22o1o0] 空点以外のところ（石または枠の上）に石を置くことの禁止 - IsMasonryError関数作成

とりあえず、 `石または枠の上に石を置く行為` に `Masonry` （メイスンリー）という名前を付ける。  
従って この主のエラーは `Masonry error` と呼ぶことにする。  
そのようなエラーであるかどうか判定する関数の名前は `IsMasonryError` と呼ぶことにする  

## Step [O22o1o1o0] ファイル作成 - masonry.go

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

## Step [O22o1o2o0] 呼出し

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
