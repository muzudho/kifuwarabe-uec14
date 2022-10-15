目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O11o_3o_11o0] ポジション作成

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# Step [O11o_3o_11o0] ポジション作成

あとで使うファイルを先に作成する  

## Step [O11o_3o_11o1o0] ファイル作成 - kernel/o11o_3o_11o0_position.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	└── 📂 kernel
👉   	└── 📄 o11o_3o_11o0_position.go
```

```go
// BOF [O11o_3o_11o1o0]

package kernel

type Position struct {
	// Board - 盤
	Board *Board
}

// NewDirtyKernel - カーネルの新規作成
// - 一部のメンバーは、初期化されていないので、別途初期化処理が要る
func NewDirtyPosition(gameRule GameRule, boardWidht int, boardHeight int) *Position {
	var p = new(Position)

	p.Board = NewBoard(gameRule, boardWidht, boardHeight)

	return p
}

// EOF [O11o_3o_11o1o0]
```
