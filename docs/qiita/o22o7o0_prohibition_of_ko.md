目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O22o7o0] コウの禁止

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# Step [O22o7o0] コウの禁止 - Ko

永遠に同じ局面が繰り返してしまうような１手は置けない  

## Step [O22o7o1o0] ファイル編集 - record_item.go

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
 	│	├── 📄 kernel_facade.go
 	│	├── 📄 kernel.go
 	│	├── 📄 liberty.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
 	│	├── 📄 play.go
 	│	├── 📄 o12o__10o1o0_point.go
👉	│	├── 📄 record_item.go
	│	├── 📄 record.go
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
// ...略...
// type RecordItem struct {
	// ...略...

	// [O22o7o1o0] コウの位置
	ko Point
// }
// ...略...
```

## Step [O22o7o1o0] ファイル編集 - record.go

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
 	│	├── 📄 kernel_facade.go
 	│	├── 📄 kernel.go
 	│	├── 📄 liberty.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
 	│	├── 📄 play.go
 	│	├── 📄 o12o__10o1o0_point.go
	│	├── 📄 record_item.go
👉	│	├── 📄 record.go
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
// ...略...
// Push - 末尾に追加
// func (r *Record) Push(placePlay Point,
	// [O22o7o1o0] コウの位置
	ko Point//) {

	// var item = r.items[r.posNum]
	// item.placePlay = placePlay

	// [O22o7o1o0] コウの位置
	item.ko = ko

	// r.posNum++
// }
```

## Step [O22o7o2o0] ファイル編集 - play.go

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	├── 📄 board1.txt
 	│	├── 📄 board2.txt
 	│	└── 📄 board3.txt
	├── 📂 kernel
	│	├── 📂 play_rule
	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 o12o__11o1o0_board.go
 	│	├── 📄 check_board.go
 	│	├── 📄 color.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	├── 📄 kernel_facade.go
 	│	├── 📄 kernel.go
 	│	├── 📄 liberty.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
👉 	│	├── 📄 play.go
 	│	├── 📄 o12o__10o1o0_point.go
 	│	├── 📄 record_item.go
 	│	├── 📄 record.go
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
// func (k *Kernel) DoPlay(command string, logg *Logger) {

	// ...略...
	// [O22o4o1o0] 自分の眼に石を置こうとした
	// var onForbiddenMyEye = func() bool {
		// ...略...
	// }

	// [O22o7o2o0] コウに石を置こうとした
	var onKo = func() bool {
		logg.C.Infof("? ko my_stone:%s point:%s\n", stone, k.Position.Board.coordinate.GetGtpMoveFromPoint(point))
		logg.J.Infow("error ko", "my_stone", stone, "point", k.Position.Board.coordinate.GetGtpMoveFromPoint(point))
		return false
	}

	// var isOk = k.Play(stone, point, logg,
		// // [O22o1o2o0] ,onMasonry
		// onMasonry,
		// // [O22o3o1o0] ,onOpponentEye
		// onOpponentEye,
		// // [O22o4o1o0] ,onForbiddenMyEye
		// onForbiddenMyEye,
		// // [O22o7o2o0] ,onKo
		onKo//)
	// ...略...

// }
// ...略...

// func (k *Kernel) Play(stoneA Stone, pointB Point, logg *Logger,
	// // [O22o1o2o0] onMasonry
	// onMasonry func() bool,
	// // [O22o3o1o0] onOpponentEye
	// onOpponentEye func() bool,
	// // [O22o4o1o0] onForbiddenMyEye
	// onForbiddenMyEye func() bool,
	// [O22o7o2o0] onKo
	onKo func() bool//) bool {

	// [O22o1o2o0]
	// if k.Position.Board.IsMasonry(pointB) {
	// 	return onMasonry()
	// }
	// ...略...

	// [O22o7o2o0] コウの判定
	if k.Record.IsKo(pointB) {
		return onKo()
	}
	// ...略...

	// * 以下を追加
	// [O22o7o2o0] コウの判定
	var capturedCount = 0 // アゲハマ

//	// [O22o6o1o0] 死に石を打ちあげる
//	if isExists4rensToRemove {
//		for dir := 0; dir < 4; dir++ {
//			var ren = o4rensToRemove[dir]
//			if ren != nil {
//				k.RemoveRen(ren)

				// * 以下を追加
				// [O22o7o2o0] コウの判定
				capturedCount += ren.GetArea()

//			}
//		}
//	}

	// [O22o7o2o0] コウの判定
	var ko = Point(0)
	if capturedCount == 1 {
		ko = pointB
	}

	// 棋譜に追加
	//k.Record.Push(pointB,
		// * 以下を追加
		// [O22o7o2o0] コウの判定
		ko//)

```

## Step [O22o7o4o0] コマンド編集 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	├── 📄 board1.txt
 	│	├── 📄 board2.txt
 	│	└── 📄 board3.txt
	├── 📂 kernel
	│	├── 📂 play_rule
	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 o12o__11o1o0_board.go
 	│	├── 📄 check_board.go
 	│	├── 📄 color.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
 	│	├── 📄 kernel_facade.go
👉 	│	├── 📄 kernel.go
 	│	├── 📄 liberty.go
 	│	├── 📄 logger.go
 	│	├── 📄 masonry.go
 	│	├── 📄 play.go
 	│	├── 📄 o12o__10o1o0_point.go
 	│	├── 📄 record_item.go
 	│	├── 📄 record.go
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
	// case "record": // [O12o__11o_5o0]
		// ...略...

		// var setPoint = func(i int, item *RecordItem) {
			// var positionNth = i + geta // 基数を序数に変換
			// var coord = k.Position.Board.GetGtpMoveFromPoint(item.placePlay)

			// * 以下を削除
			// sb.WriteString(fmt.Sprintf("[%d]%s ", positionNth, coord))

			// * 以下を追加
			// [O22o7o4o0] コウを追加
			var koStr string
			if item.ko == Point(0) {
				koStr = ""
			} else {
				koStr = fmt.Sprintf("(%s)", k.Position.Board.coordinate.GetGtpMoveFromPoint(item.ko))
			}
			sb.WriteString(fmt.Sprintf("[%d]%s%s ", positionNth, coord, koStr))
		// }
		// ...略...

	// ...略...
	// この上にコマンドを挟んでいく
	// -------------------------
	// ...略...
```

## Step [O22o7o4o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
board_set file data/board3.txt
play black D3
play white C3
play black D3
```

Output > Console:  

```plaintext
[2022-09-17 22:39:55]   # board_set file data/board3.txt
[2022-09-17 22:39:55]   =

[2022-09-17 22:39:55]   # play black D3
[2022-09-17 22:39:55]   =

[2022-09-17 22:39:55]   # play white C3
[2022-09-17 22:39:55]   =

[2022-09-17 22:39:55]   # play black D3
[2022-09-17 22:39:55]   ? ko my_stone:x point:D3
```
