目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞ Step [O11o_3o_10o0] ゲームルール作成

# 連載の目次

📖 [目指せ！第１４回ＵＥＣ杯コンピューター囲碁大会＜本編＞](https://qiita.com/muzudho1/items/744f6051954525878b84)  

# Step [O11o_3o_10o0] ゲームルール作成

あとで使うファイルを先に作成する  

## Step [O11o_3o_10o1o0] ファイル作成 - kernel/o11o_3o_10o0_game_rule.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
	└── 📂 kernel
👉   	└── 📄 o11o_3o_10o0_game_rule.go
```

```go
// BOF [O11o_3o_10o1o0]

package kernel

// KomiFloat - コミ。 6.5 といった数字を入れるだけ。実行速度優先で 64bitに
type KomiFloat float64

// PositionNumberInt - 何手目
type PositionNumberInt int

// GameRule - 対局ルール
type GameRule struct {
	// コミ。 6.5 といった数字を入れるだけ。実行速度優先で 64bitに
	komi KomiFloat

	// 上限手数
	maxPositionNumber PositionNumberInt
}

// NewGameRule - 新規作成
func NewGameRule(komi KomiFloat, maxPositionNumber PositionNumberInt) *GameRule {
	var gr = new(GameRule)

	gr.komi = komi
	gr.maxPositionNumber = maxPositionNumber

	return gr
}

// GetKomi - コミ取得
func (gr *GameRule) GetKomi() KomiFloat {
	return gr.komi
}

// GetMaxPositionNumber - 上限手数
func (gr *GameRule) GetMaxPositionNumber() PositionNumberInt {
	return gr.maxPositionNumber
}

// EOF [O11o_3o_10o1o0]
```
