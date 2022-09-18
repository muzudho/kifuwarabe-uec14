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
	// 何手目。基数（Position number）
	posNum int

	// その連
	ren *Ren
}

// NewRenDbItem - 連Dbの要素を新規作成する
func NewRenDbItem(positionNumber int, ren *Ren) *RenDbItem {
	var i = new(RenDbItem)
	i.posNum = positionNumber
	i.ren = ren
	return i
}

// Dump - ダンプ
func (ri *RenDbItem) Dump() string {
	return fmt.Sprintf("n:%d ren:%s", ri.posNum, ri.ren.Dump())
}

// EOF [O1o1o0g12o__11o__10o1o0]
