# 目次

📖 [Step [O1o0] 導入](https://qiita.com/muzudho1/items/3a78087f812bab4a511f)  
📖 [Step [O11o___100o0] カーネル作成 ～ Step [O11o__10o_1o0] 思考エンジン設定ファイル](https://qiita.com/muzudho1/items/6c0b16d3b87ac598fe86)  
📖 [Step [O11o__10o0] ロガー設定 ～ Step [O11o__10o3o_2o0] welcome プログラム](https://qiita.com/muzudho1/items/26af2c9f5dcc16175acd)  
📖 [Step [O11o__11o0] デバッグ可能標準入力 作成](https://qiita.com/muzudho1/items/252eef6d00417dbd82a1)  
📖 [Step [O11o_3o_10o0] ゲームルール作成](https://qiita.com/muzudho1/items/a102ec5b03811f34b22c)  
📖 [Step [O11o_3o0] カーネルのインタープリター](https://qiita.com/muzudho1/items/374b040f4e025f42b970)  
📖 [Step [O11o_4o0] 石の色定義](https://qiita.com/muzudho1/items/835e0f7ac5160f62bf80)  
📖 [Step [O11o_4o2o0] 連の定義](https://qiita.com/muzudho1/items/68256c2784fa68c71a57)  
📖 [Step [O11o_5o0] 石定義](https://qiita.com/muzudho1/items/f7bb526760c8f995bcca)  
📖 [Step [O12o__10o0] 点定義、またはその盤座標符号定義](https://qiita.com/muzudho1/items/6137d99ce30d962af338)  
📖 [Step [O12o__11o__100o0] データファイル作成 ～ Step [O12o__11o__10o0] 連データベース定義 ～ Step [O12o__11o__10o5o__10o0] 連データベースのロード](https://qiita.com/muzudho1/items/03454c8efcf8b8927086)  
📖 [Step [O12o__11o_1o0] 棋譜定義](https://qiita.com/muzudho1/items/c00dd77f8cefc5c12dcd)  
📖 [Step [O12o__11o0] 盤定義（土台）](https://qiita.com/muzudho1/items/d6584b3a7a200aa90e1b)  
📖 [Step [O12o_1o0] 盤定義（盤面走査）](https://qiita.com/muzudho1/items/ff6e4dc163ecbc023796)  
📖 [Step [O13o_1o0] 盤表示](https://qiita.com/muzudho1/items/3f582b7db2b8c05adb22)  
📖 [Step [O15o__10o0] 盤サイズの変更](https://qiita.com/muzudho1/items/9907f1878480877661fb)  
📖 [Step [O15o__14o_1o0] データファイル作成](https://qiita.com/muzudho1/items/91c040149eecabd3e78c)  
📖 [Step [O15o__14o0] 初期盤面を設定する](https://qiita.com/muzudho1/items/0cc500bc18b36bb7cbd2)  
📖 [Step [O15o_1o0] 座標の定義](https://qiita.com/muzudho1/items/b443433b89dfae8a8098)  
📖 [Step [O19o_1o0] 石を打つ](https://qiita.com/muzudho1/items/99a3bad83d0f7d2887a6)  
📖 [Step [O22o6o0] 死に石の連の打ち上げ](https://qiita.com/muzudho1/items/c7009fc251bb0a06b090)  
📖 [Step [O22o7o0] コウの禁止](https://qiita.com/muzudho1/items/a98a287015beee4601c9)  

# Step [O23o_1o0] データファイル作成 - data/board4.txt ファイル

あとで使うファイルを先に作成する  

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	├── 📄 board1.txt
 	│	├── 📄 board2.txt
 	│	├── 📄 board3.txt
👉 	│	└── 📄 board4.txt
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

```plaintext
     A B C D E F G H J K L M N O P Q R S T
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
 1 + x . x x x . . . . . . . . . x x x . x +
 2 + . . . . . . . . . . . . . . . . . . . +
 3 + x . x x x x x . o o o . . . . . . . x +
 4 + x . x . . . x . o o o x x x . . . . x +
 5 + x . x . x . x . o o o x o x . . . . x +
 6 + . . x x x . x . . . . x . x . . . . . +
 7 + . . . . . . x . . . . x x x . . . . . +
 8 + . . x x x x x . . . . . . o o o o o . +
 9 + . . . . . . . . . . . . . o . o . o . +
10 + o o o o o o o o o o o o o o o o o o o +
11 + . . . . . . . x . . x . . . . . . . . +
12 + . . . . . o . x x . x . x x x x x . . +
13 + . . . . o x o . x x x . x . . . . . . +
14 + . . . o x . x o . . . . x . x x x . . +
15 + x . o x . . . x o . . . x . x . x . x +
16 + x . . o x . x o . . . . x . . . x . x +
17 + x . . . o x o . . . . . x x x x x . x +
18 + . . . . . o . . . . . . . . . . . . . +
19 + x . x x x . . . . . . . . . x x x . x +
   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
```

# Step [O23o_2o0] 盤上の連のスキャン

`石を打つ` ことを実装できたなら、呼吸点を調べるアルゴリズムも実装されていて、  
着手点に隣接する上下左右にある石の連の認識はできているはずだ  

次は、盤面全体に点在する連を認識できるか試したい  

## Step [O23o_2o1o0] ファイル編集 - kernel_facade.go

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	├── 📄 board1.txt
 	│	├── 📄 board2.txt
 	│	├── 📄 board3.txt
 	│	└── 📄 board4.txt
	├── 📂 kernel
	│	├── 📂 play_rule
	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
  	│	├── 📄 o12o__11o1o0_board.go
 	│	├── 📄 check_board.go
 	│	├── 📄 color.go
	│	├── 📄 go.mod
	│	├── 📄 go.sum
👉 	│	├── 📄 kernel_facade.go
 	│	├── 📄 kernel.go
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

👇 がんばって以下を追加してほしい  

```go
// ...略...

// FindAllRens - [O23o_2o1o0] 盤上の全ての連を見つけます
// * 見つけた連は、連データベースへ入れます
func (k *Kernel) FindAllRens() {
	// チェックボードの初期化
	k.CheckBoard.Init(k.Board.coordinate)

	var maxPosNthFigure = k.Record.GetMaxPosNthFigure()

	var setLocation = func(location Point) {
		var ren, isFound = k.findRen(location)
		if isFound {
			k.renDb.RegisterRen(maxPosNthFigure, k.Record.posNum, ren)
		}
	}
	// 盤上の枠の内側をスキャン。筋、段の順
	k.Board.GetCoordinate().ForeachPayloadLocationOrderByYx(setLocation)
}

// ...略...
```

## Step [O23o_2o2o0] コマンド編集 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	├── 📄 board1.txt
 	│	├── 📄 board2.txt
 	│	├── 📄 board3.txt
 	│	└── 📄 board4.txt
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
	case "find_all_rens": // [O23o_2o2o0]
		// Example: `find_all_rens`
		k.FindAllRens()
		logg.C.Infof("=\n")
		logg.J.Infow("ok")
		return true

	// ...略...
	// この上にコマンドを挟んでいく
	// -------------------------
	// ...略...
```

## Step [O23o_2o3o_1o0] ファイル編集 - board_set.go

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	├── 📄 board1.txt
 	│	├── 📄 board2.txt
 	│	├── 📄 board3.txt
 	│	└── 📄 board4.txt
	├── 📂 kernel
	│	├── 📂 play_rule
	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
👉 	│	├── 📄 board_set.go
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

👇 がんばって以下を追加してほしい  

```go
// ...略...
// func (k *Kernel) DoSetBoard(command string, logg *Logger) {
	// if tokens[1] == "file" {
		// ...略...
		// if int(i) != size {
		// 	logg.C.Infof("? not enough size i:%d size:%d\n", i, size)
		// 	logg.J.Infow("error not enough size", "i", i, "size", size)
		// 	return
		// }

		// * 以下を追加
		// [O23o_2o3o_1o0] 連データベース初期化
		k.renDb.Init(k.Board.GetWidth(), k.Board.GetHeight())
		k.FindAllRens()
	// }
// }
// ...略...
```

## Step [O23o_2o3o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
board_set file data/board4.txt
rendb_dump
find_all_rens
rendb_dump
rendb_save data/ren_db_board4_temp.json
```

Output > Console:  

```plaintext
[2022-09-18 23:58:02]   # board_set file data/board4.txt
[2022-09-18 23:58:02]   =

[2022-09-18 23:58:51]   # find_all_rens
[2022-09-18 23:58:51]   =
```

# Step [O23o_2o4o0] TODO 出力ファイルの内容の盤表示

# Step [O23o_2o4o1o0] ファイル編集 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 data
 	│	├── 📄 board1.txt
 	│	├── 📄 board2.txt
 	│	├── 📄 board3.txt
 	│	└── 📄 board4.txt
	├── 📂 kernel
	│	├── 📂 play_rule
	│	├── 📄 board_area.go
  	│	├── 📄 board_coord.go
 	│	├── 📄 board_set.go
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

```go
```

TODO 📄 `ren_db1_temp.json` ファイルを読み込む  

TODO 連データベースには「何手目」の「どの連」という形で項目がいくつもある。その全項目をコンソールへ盤表示  

# Step [O23o_2o4o2o0]

TODO FIXME rendb_dump の結果が全部空っぽ？

📄 .vscode/launch.json:  

```json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}",
            "args" : ["-d=true"]
        }
    ]
}
```

📄 debug.input.txt:  

```plaintext
rendb_load data/ren_db_board4.json
rendb_dump
```

```plaintext
rendb_load data/ren_db_board4.json
find_all_rens
rendb_dump
```

# Step [O23o_3o0] TODO 石を打った後の連の再スキャン

# Step [O23o0] 打った石のアンドゥ - Undo

打った石をやっぱり止める、一手戻す、ということは、石を打つよりも実装がむずかしい  

### Step [O23o1o0] ファイル作成 - play_undo.go

👇 以下のファイルを新規作成してほしい  

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
👉 	│	├── 📄 play_undo.go
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

```go
// BOF [O23o1o0]

package kernel

// DoUndoPlay - 石を打ったのを戻す
//
// * `command` - Example: `undo`
// ........................----
// ........................0
func (k *Kernel) DoUndoPlay(command string, logg *Logger) {
	k.UndoPlay()
}

// UndoPlay - 石を打ったのを戻す
//
// Returns
// -------
// isOk : bool
//
//	石を置けたら真、置けなかったら偽
func (k *Kernel) UndoPlay() bool {

	// 初期局面から前には戻せない
	if k.Record.GetPositionNumber() < 1 {
		return false
	}

	return false
}

// EOF [O23o1o0]
```

TODO 東、北、西、南に隣接する連を重複して数えないように常に注意すること
TODO 棋譜に取った石を記録する
TODO 取った石を戻す
TODO アンドゥ

# 参考にした記事

## Go言語

### ファイル分割

📖 [[Go言語] ファイル分割とローカルパッケージ](https://zenn.dev/fm_radio/articles/ca2ff1dfcf89b5)  

### コマンドライン引数

📖 [Goでflagを使ってコマンドライン引数を扱う](https://qiita.com/Yaruki00/items/7edc04720a24e71abfa2)  

### 列挙型

📖 [How to make Go print enum fields as string?](https://stackoverflow.com/questions/41480543/how-to-make-go-print-enum-fields-as-string)  

### 出力

📖 [Go で値を出力する方法](https://golang.keicode.com/basics/go-print-basics.php)  

### 可変長引数

📖 [Concatenating and Building Strings in Go 1.10+](https://www.calhoun.io/concatenating-and-building-strings-in-go/)  
📖 [Convert interface to string](https://yourbasic.org/golang/interface-to-string/)  

### 文字列

📖 [Go: 1文字ずつアクセスする](https://blog.sarabande.jp/post/61104920128)  
📖 [Golang String Padding Example](https://golang.cafe/blog/golang-string-padding-example.html)  

### 配列

📖 [スライス操作(要素の追加・削除, ソート, 他のスライスと結合)](https://www.wakuwakubank.com/posts/782-go-slice/)  

### ファイル入出力

📖 [Read a file in Go](https://gosamples.dev/read-file/#:~:text=The%20simplest%20way%20of%20reading,by%20line%20or%20in%20chunks.)  

### JSON

📖 [Goにおけるjsonの扱い方を整理・考察してみた ~ データスキーマを添えて](https://zenn.dev/hsaki/articles/go-convert-json-struct)  

### コレクション

#### ハッシュテーブル

📖 [[Go言語] 初心者必見シリーズ: マップ（Map）](https://qiita.com/wifecooky/items/2ffe41d55c575b2ce5e2)  
📖 [Go言語: マップのキーが存在するかチェックしたい](https://qiita.com/suin/items/4cb1da71237fc55a06ee)  

### ビット演算

📖 [Goのビット演算について](https://www.flyenginer.com/low/go/go%E3%81%AE%E3%83%93%E3%83%83%E3%83%88%E6%BC%94%E7%AE%97%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6.html)  
📖 [How to define bit literals in Go?](https://stackoverflow.com/questions/56605810/how-to-define-bit-literals-in-go)  

### 数学

📖 [[math] 2つの値のうち小さい方を返す (Min)](http://www.openspc2.org/reibun/Go/1.1.1/pkg/math/1043/index.html)  

### デバッギング

📖 [標準入力のあるプログラムを delve でデバッグしたい](https://qiita.com/_natsu_no_yuki_/items/505e74e598d3d6a0cb24)  

.
