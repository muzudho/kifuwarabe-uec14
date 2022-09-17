// BOF[O1o1o0g12o__11o_2o_1o0]

package kernel

// RecordItem - 棋譜の一手分
type RecordItem struct {
	// 着手点
	placePlay Point

	// [O1o1o0g22o7o1o0] コウの位置
	ko Point
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

// EOF[O1o1o0g12o__11o_2o_1o0]
