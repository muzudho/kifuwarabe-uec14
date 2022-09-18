// BOF [O1o1o0g12o__11o__10o2o0]

package kernel

type RenDb struct {
	// 盤サイズ
	boardMemoryArea int

	// 要素
	items map[RenDbItemId]*RenDbItem
}

// NewRenDb - 連データベースを新規作成
func NewRenDb(boardMemoryArea int) *RenDb {
	var r = new(RenDb)
	r.boardMemoryArea = boardMemoryArea
	return r
}

// FindRen - 連を取得
func (r *RenDb) GetRen(renDbItemId RenDbItemId) (*Ren, bool) {
	var item, isOk = r.items[renDbItemId]

	if isOk {
		return item.ren, true
	}

	return nil, false
}

// RegisterRen - 連を登録
func (r *RenDb) RegisterRen(positionNumber int, ren *Ren) {
	var renDbItemId = GetRenDbItemId(r.boardMemoryArea, positionNumber, ren.minimumLocation)
	r.items[renDbItemId] = NewRenDbItem(positionNumber, ren)
}

// EOF [O1o1o0g12o__11o__10o2o0]
