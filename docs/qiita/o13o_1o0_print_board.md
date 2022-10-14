目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O13o_1o0] 盤表示

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# Step [O13o_1o0] 盤表示 - board コマンド

## Step [O13o0] コマンド実装 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
  	│	├── 📄 board_area.go
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

	case "board": // [O13o0]
		// 人間向けの出力
		{
			// 25列まで対応
			const fileSimbols = "ABCDEFGHJKLMNOPQRSTUVWXYZ"
			// 25行まで対応
			var rankSimbols = strings.Split("  , 1, 2, 3, 4, 5, 6, 7, 8, 9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25", ",")

			var filesMax = int(math.Min(25, float64(k.Board.coordinate.GetWidth())))
			var rowsMax = int(math.Min(25, float64(k.Board.coordinate.GetHeight())))
			var filesLabel = fileSimbols[:filesMax]

			var sb strings.Builder
			// 枠の上辺
			sb.WriteString(fmt.Sprintf(`= board:'''
.     %s
.    `, filesLabel))

			var rowNumber = 1
			var setPoint = func(point Point) {
				var stone = k.Board.cells[point]
				sb.WriteString(fmt.Sprintf("%v", stone))
			}
			var doNewline = func() {
				var rankLabel string
				if rowNumber <= rowsMax {
					rankLabel = rankSimbols[rowNumber]
				} else {
					rankLabel = ""
				}

				sb.WriteString(fmt.Sprintf("\n. %2s ", rankLabel))
				rowNumber++
			}
			k.Board.GetCoordinate().ForeachLikeText(setPoint, doNewline)
			sb.WriteString("\n. '''\n")
			logg.C.Info(sb.String())
		}
		// コンピューター向けの出力
		{
			var sb strings.Builder

			var setPoint = func(point Point) {
				var stone = k.Board.cells[point]
				sb.WriteString(fmt.Sprintf("%v", stone))
			}
			var doNewline = func() {
				// pass
			}
			k.Board.GetCoordinate().ForeachLikeText(setPoint, doNewline)
			logg.J.Infow("output", "board", sb.String())
		}
		return true

	// この上にコマンドを挟んでいく
	// -------------------------
	// ...略...
```

## Step [O14o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
board
```

Output > Console:  

```plaintext
[2022-09-19 15:50:11]   # board
[2022-09-19 15:50:11]   = board:'''
.     ABCDEFGHJKLMNOPQRST
.    +++++++++++++++++++++
.  1 +...................+
.  2 +...................+
.  3 +...................+
.  4 +...................+
.  5 +...................+
.  6 +...................+
.  7 +...................+
.  8 +...................+
.  9 +...................+
. 10 +...................+
. 11 +...................+
. 12 +...................+
. 13 +...................+
. 14 +...................+
. 15 +...................+
. 16 +...................+
. 17 +...................+
. 18 +...................+
. 19 +...................+
.    +++++++++++++++++++++
. '''
```

* この出力書式は 私の方法であって、公式大会のものではないことに注意されたい
	* 📖 [思考エンジンの思考ログ仕様（きふわらべ2022年以降）](https://qiita.com/muzudho1/items/ceb6130cf558cd373dd7)

👇 以下のファイルが自動生成された  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
  	│	├── 📄 board_area.go
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
👉 	├── 📄 kifuwarabe-uec14-json.log
👉 	├── 📄 kifuwarabe-uec14.log
  	└── 📄 main.go
```

👇 📄 `kifuwarabe-uec14-json.log`  

```plaintext
{"level":"info","ts":"2022-09-11T19:04:58.192+0900","caller":"kifuwarabe-uec14/main.go:57","msg":"input","command":"board"}
{"level":"info","ts":"2022-09-11T19:04:58.195+0900","caller":"kifuwarabe-uec14/main.go:81","msg":"output","board":"++++++++++++++++++++++...................++...................++...................++...................++...................++...................++...................++...................++...................++...................++...................++...................++...................++...................++...................++...................++...................++...................++...................++++++++++++++++++++++"}
```

👇 📄 `kifuwarabe-uec14.log`  

```plaintext
2022-09-11T19:04:58.157+0900	# board
2022-09-11T19:04:58.192+0900	= board:'''
. +++++++++++++++++++++
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
. +...................+
. +...................+
. +++++++++++++++++++++
. '''
```

`quit` コマンドで 思考エンジンを終了してほしい  
