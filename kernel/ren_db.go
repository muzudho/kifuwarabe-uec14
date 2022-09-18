// BOF [O1o1o0g12o__11o__10o2o0]

package kernel

type RenDb struct {
	items map[int]*RenDbItem
}

// NewRenDb - 連データベースを新規作成する
func NewRenDb() *RenDb {
	var r = new(RenDb)
	return r
}

// FindRen - 連を探す
func (r *RenDb) FindRen(boardMemoryArea int, positionOrdinal int, minimumLocation Point) *Ren {
	var id = GetRenDbItemId(boardMemoryArea, positionOrdinal, minimumLocation)
	var item = r.items[id]

	if item == nil {
		return nil
	}

	return item.ren
}

// EOF [O1o1o0g12o__11o__10o2o0]
