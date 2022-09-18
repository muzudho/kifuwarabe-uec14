// BOF [O1o1o0g12o__11o__10o1o0]

package kernel

// 連データベースの要素
type RenDbItem struct {
	// 何手目か （Position ordinals）
	posOrd int

	// その連
	ren *Ren
}

// NewRenDbItem - 連Dbの要素を新規作成する
func NewRenDbItem(positionOrdinal int, ren *Ren) *RenDbItem {
	var i = new(RenDbItem)
	i.posOrd = positionOrdinal
	i.ren = ren
	return i
}

// GetId - 連のIdを取得
func GetRenDbItemId(boardMemoryArea int, positionOrdinal int, minimumLocation Point) int {
	return positionOrdinal*boardMemoryArea + int(minimumLocation)
}

// EOF [O1o1o0g12o__11o__10o1o0]
