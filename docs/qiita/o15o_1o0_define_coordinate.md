目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O15o_1o0] 座標の定義

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# Step [O15o_1o0] 座標の定義

## ~~Step [O15o0]~~

Moved to `[O12o__10o1o0]`  

## ~~Step [O16o0]~~

Remove  

### Step [O16o1o0] コマンド実装 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 board_coord.go
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
	// ...略...

	// * アルファベット順になる位置に、以下のケース文を挿入
	case "test_get_point_from_code": // [O16o1o0]
		// Example: "test_get_point_from_code A1"
		var point = k.Position.Board.coordinate.GetPointFromGtpMove(tokens[1])
		var code = k.Position.Board.coordinate.GetGtpMoveFromPoint(point)
		logg.C.Infof("= %d %s", point, code)
		logg.J.Infow("ok", "point", point, "code", code)
		return true

	// ...略...
	// この上にコマンドを挟んでいく
	// -------------------------
	// ...略...
```

## Step [O16o2o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
test_get_point_from_code A1
```

Output > Console:  

```plaintext
[2022-09-17 17:18:46]   # test_get_point_from_code A1
[2022-09-17 17:18:46]   = 22 A1
```

## ~~Step [O17o0]~~

Moved to `[O12o__10o2o0]`  

## ~~Step [O18o0]~~

Moved to `[O12o__10o3o0]`  
