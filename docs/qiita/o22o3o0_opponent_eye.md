目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O22o3o0] 相手の眼に石を置くことの禁止

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# Step [O22o3o0] 相手の眼に石を置くことの禁止 - OpponentEye

囲碁のルールでは、相手の眼へは石を置けない。これを判定する  

とりあえず、このルールへ直訳で短い名前を付ける。 仮に `OpponentEye` とでも呼ぶことにする  

このルールは、あとで出てくる `Captured` のルールよりは優先度が低いとする  

## Step [O22o3o1o0] ファイル編集 - kernel.go

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	└── 📄 board1.txt
	├── 📂 kernel
	│	├── 📂 play_rule
  	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 o12o__11o1o0_board.go
 	│	├── 📄 check_board.go
 	│	├── 📄 color.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	├── 📄 kernel.go
 	│	├── 📄 liberty.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
👉 	│	├── 📄 play.go
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
// func (k *Kernel) DoPlay(command string, logg *Logger) {

	// ...略...
	// [O22o3o1o0] 相手の眼に石を置こうとした
	var onOpponentEye = func() bool {
		logg.C.Infof("? opponent_eye my_stone:%s point:%s\n", stone, k.Position.Board.coordinate.GetGtpMoveFromPoint(point))
		logg.J.Infow("error opponent_eye", "my_stone", stone, "point", k.Position.Board.coordinate.GetGtpMoveFromPoint(point))
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
	// if k.Position.Board.IsMasonry(point) {
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
	// k.Position.Board.cells[point] = stone
	// return true
```

## Step [O22o3o2o0] 動作確認

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
