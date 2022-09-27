目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O12o__11o_1o0] 棋譜定義

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# Step [O12o__11o_1o0] 棋譜定義

## Step [O12o__11o_2o_1o0] ファイル作成 - record_item.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
	│	├── 📄 point.go
👉	│	├── 📄 record_item.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
  	└── 📄 main.go
```

```go
// BOF[O12o__11o_2o_1o0]

package kernel

// RecordItem - 棋譜の一手分
type RecordItem struct {
	// 着手点
	placePlay Point
}

// NewRecordItem - 棋譜の一手分
func NewRecordItem() *RecordItem {
	var ri = new(RecordItem)
	return ri
}

// Clear - 空っぽにします
func (ri *RecordItem) Clear() {
	ri.placePlay = Point(0)
	ri.ko = Point(0)
}

// EOF[O12o__11o_2o_1o0]
```

## Step [O12o__11o_2o0] ファイル作成 - record.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
	│	├── 📄 point.go
	│	├── 📄 record_item.go
👉	│	├── 📄 record.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
  	└── 📄 main.go
```

```go
// BOF [O12o__11o_2o0]

package kernel

// Record - 棋譜
type Record struct {
	// 先行
	playFirst Stone

	// 何手目。基数（Position number）
	posNum int

	// 手毎
	items []*RecordItem
}

// NewRecord - 棋譜の新規作成
func NewRecord(maxMoves int, playFirst Stone) *Record {
	var r = new(Record)
	r.playFirst = playFirst

	// 棋譜の一手分毎
	r.items = make([]*RecordItem, maxMoves)
	for i := 0; i < maxMoves; i++ {
		r.items[i] = NewRecordItem()
	}

	return r
}

// GetMaxPosNthFigure - 手数（序数）の最大値の桁数
func (r *Record) GetMaxPosNthFigure() int {
	var nth = r.GetMaxPosNth()
	var nthText = strconv.Itoa(nth)
	return len(nthText)
}

// GetMaxPosNth - 手数（序数）の最大値
func (r *Record) GetMaxPosNth() int {
	return len(r.items) + geta
}

// GetPositionNumber - 何手目。基数
func (r *Record) GetPositionNumber() int {
	return r.posNum
}

// Push - 末尾に追加
func (r *Record) Push(placePlay Point) {

	var item = r.items[r.posNum]
	item.placePlay = placePlay

	r.posNum++
}

// RemoveTail - 末尾を削除
func (r *Record) RemoveTail(placePlay Point) {
	r.posNum--
	r.items[r.posNum].Clear()
}

// ForeachItem - 各要素
func (r *Record) ForeachItem(setItem func(int, *RecordItem)) {
	for i := 0; i < r.posNum; i++ {
		setItem(i, r.items[i])
	}
}

// IsKo - コウか？
func (r *Record) IsKo(placePlay Point) bool {
	// [O22o7o1o0] コウの判定
	// 2手前に着手して石をぴったり１つ打ち上げたとき、その着手点はコウだ
	var posNum = r.GetPositionNumber()
	if 2 <= posNum {
		var item = r.items[posNum-2]
		return item.ko == placePlay
	}

	return false
}

// EOF [O12o__11o_2o0]
```

## Step [O12o__11o_3o0] ファイル編集 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
👉 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
	│	├── 📄 point.go
	│	├── 📄 record.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
  	└── 📄 main.go
```

```go
// type Kernel struct {
	// ...略...

	// Record - [O12o__11o_3o0] 棋譜
	Record Record

// }

// NewKernel - カーネルの新規作成
// func NewKernel(boardWidht int, boardHeight int,
	// [O12o__11o_2o0] ,棋譜の初期化
	maxMoves int, playFirst Stone//) *Kernel {
	// ...略...

	// * 以下を追加
	// [O12o__11o_2o0] 棋譜の初期化
	k.Record = *NewRecord(maxMoves, playFirst)

	// ...略...
	// return k
// }
```

## Step [O12o__11o_4o0] ファイル編集 - main.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
	│	├── 📄 point.go
	│	├── 📄 record.go
 	│	└── 📄 stone.go
    ├── 📄 .gitignore
 	├── 📄 engine_config.go
  	├── 📄 engine.toml
    ├── 📄 go.mod
  	├── 📄 go.work
👉 	└── 📄 main.go
```

```go
		// [O11o_3o0]
		//var kernel1 = kernel.NewKernel(engineConfig.GetBoardSize(), engineConfig.GetBoardSize(),
			// [O12o__11o_4o0] 棋譜の初期化
			engineConfig.GetMaxMovesNum(),
			kernel.GetStoneOrDefaultFromTurn(engineConfig.GetPlayFirst(), onUnknownTurn)//)
```

## Step [O12o__11o_5o0] コマンド実装 - kernel.go ファイル

👇 以下の既存ファイルを編集してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	├── 📂 kernel
	│	├── 📄 go.mod
👉 	│	├── 📄 kernel.go
 	│	├── 📄 logger.go
	│	├── 📄 point.go
	│	├── 📄 record.go
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
	case "record": // [O12o__11o_5o0]
		// Example: "record"
		var sb strings.Builder

		var setPoint = func(positionNumber int, item *RecordItem) {
			var positionNth = positionNumber + geta // 基数を序数に変換
			var coord = k.Board.GetCodeFromPoint(item.placePlay)
			sb.WriteString(fmt.Sprintf("[%d]%s ", positionNth, coord))
		}

		k.Record.ForeachItem(setPoint)

		var text = sb.String()
		if 0 < len(text) {
			text = text[:len(text)-1]
		}
		logg.C.Infof("= record:'%s'\n", text)
		logg.J.Infow("ok", "record", text)
		return true

	// ...略...
	// この上にコマンドを挟んでいく
	// -------------------------
	// ...略...
```

## Step [O12o__11o_6o0] 動作確認

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい

Input:  

```shell
go run .
```

これで、思考エンジン内の入力待機ループに入った  

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
play black A1
play black B2
play black C3
play black D4
play black E5
record
```

Output > Console:  

```plaintext
[2022-09-17 19:05:16]   # record
[2022-09-17 19:05:16]   = record:'[1]A1 [2]B2 [3]C3 [4]D4 [5]E5'
```
