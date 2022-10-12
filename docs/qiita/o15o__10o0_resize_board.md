目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O15o__10o0] 盤サイズの変更

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# Step [O15o__10o0] 盤サイズの変更 - resize コマンド

## Step [O15o__11o0] コマンド実装 - kernel.go ファイル

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

	case "boardsize": // [O15o__11o0]
		// Example: `boardsize 19`
		var sideLength, err = strconv.Atoi(tokens[1])

		if err != nil {
			logg.C.Infof("? unexpected sideLength:%s\n", tokens[1])
			logg.J.Infow("error", "sideLength", tokens[1])
			return true
		}

		k.Board.Init(sideLength, sideLength)
		logg.C.Info("=\n")
		logg.J.Infow("ok")

		return true

	// ...略...

	// この上にコマンドを挟んでいく
	// -------------------------

// ...略...
```

## Step [O15o__12o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
boardsize 9
```

Output > Console:  

```plaintext
[2022-09-12 20:44:12]   # boardsize 9
[2022-09-12 20:44:12]   =
```

Output > Log > Plain:  

```plaintext
2022-09-12T20:44:12.860+0900	# boardsize 9
2022-09-12T20:44:12.896+0900	=
```

Output > Log > Json:  

```json
{"level":"info","ts":"2022-09-12T20:44:12.894+0900","caller":"kifuwarabe-uec14/main.go:61","msg":"input","command":"boardsize 9"}
{"level":"info","ts":"2022-09-12T20:44:12.896+0900","caller":"kernel/kernel.go:76","msg":"ok"}
```

Input:  

```shell
board
```

Output:  

```plaintext
[2022-09-12 20:56:20]   # board
[2022-09-12 20:56:20]   = board:'''
. +++++++++++
. +.........+
. +.........+
. +.........+
. +.........+
. +.........+
. +.........+
. +.........+
. +.........+
. +.........+
. +++++++++++
. '''
```

# ~~Step [O15o__13o0]~~

Removed  

## ~~Step [O15o__13o1o0]~~

Moved to `O11o__10o_2o0`  

## ~~Step [O15o__13o2o_1o0]~~

Moved to `[O11o__10o_3o0]`  

## ~~Step [O15o__13o2o_2o0]~~

Moved to `[O11o__10o_4o0]`  

## ~~Step [O15o__13o2o_3o0]~~

Merged to `[O11o_3o0]`  

## ~~Step [O15o__13o2o_4o0]~~

Moved to `[O11o__10o_6o0]`  

