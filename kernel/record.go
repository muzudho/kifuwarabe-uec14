// BOF [O1o1o0g12o__11o_2o0]

package kernel

// Record - 棋譜
type Record struct {
	// 先行
	playFirst []Stone

	// 着手点
	points []Point
}

// NewRecord - 棋譜の新規作成
func NewRecord(maxMoves int) *Record {
	var r = new(Record)
	r.playFirst = make([]Stone, maxMoves)
	r.points = make([]Point, maxMoves)
	return r
}

// EOF [O1o1o0g12o__11o_2o0]
