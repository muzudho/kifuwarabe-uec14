// BOF [O1o1o0g12o__11o__10o1o0]

package kernel

import (
	"fmt"
)

// RenDbItemId - 連データベースの要素のId
type RenDbItemId int

// GetId - 連のIdを取得
func GetRenDbItemId(boardMemoryArea int, positionNumber int, minimumLocation Point) RenDbItemId {
	return RenDbItemId(positionNumber*boardMemoryArea + int(minimumLocation))
}

// 連データベースの要素
type RenDbItem struct {
	// PosNth - 何手目。序数。外部ファイルと入出力するときのみ使う
	PosNth int

	// この連のインスタンス化
	ren *Ren
}

// NewRenDbItem - 連Dbの要素を新規作成する
func NewRenDbItem(positionNumber int, ren *Ren) *RenDbItem {
	var i = new(RenDbItem)
	i.PosNth = positionNumber + geta
	i.ren = ren
	return i
}

// Dump - ダンプ
func (ri *RenDbItem) Dump() string {
	return fmt.Sprintf("pos:%dth ren:%s", ri.PosNth, ri.ren.Dump())
}

// EOF [O1o1o0g12o__11o__10o1o0]
