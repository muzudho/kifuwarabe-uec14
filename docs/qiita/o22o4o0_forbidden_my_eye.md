目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O22o4o0] 自分の眼に石を置くことの任意の禁止

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

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
		logg.C.Infof("? my_eye my_stone:%s point:%s\n", stone, k.Position.Board.coordinate.GetGtpMoveFromPoint(point))
		logg.J.Infow("error my_eye", "my_stone", stone, "point", k.Position.Board.coordinate.GetGtpMoveFromPoint(point))
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

		} else if k.Position.CanNotPutOnMyEye && stoneA.GetColor() == renC.adjacentColor {
			// [O22o4o1o0]
			// かつ、連Cに隣接する連の色が、石Aの色であったなら、
			// 自分の眼に石を置こうとしたとみなす
			return onForbiddenMyEye()

	// }

	// ...略...
	// k.Position.Board.cells[point] = stone
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
  	│	├── 📄 o12o__11o1o0_board.go
 	│	├── 📄 check_board.go
 	│	├── 📄 color.go
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
			var value = k.Position.CanNotPutOnMyEye
			logg.C.Infof("= %t\n", value)
			logg.J.Infow("ok", "value", value)
			return true

		case "set":
			var value = tokens[2]
			switch value {
			case "true":
				k.Position.CanNotPutOnMyEye = true
				return true
			case "false":
				k.Position.CanNotPutOnMyEye = false
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
