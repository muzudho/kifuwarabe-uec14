// BOF [O1o1o0g12o__11o_2o0]

package kernel

// Record - 棋譜
type Record struct {
	// 先行
	playFirst Stone

	// 着手点
	points []Point
}

// NewRecord - 棋譜の新規作成
func NewRecord(maxMoves int, playFirst Stone) *Record {
	var r = new(Record)
	r.playFirst = playFirst
	r.points = make([]Point, maxMoves)
	return r
}

// Push - 末尾に追加
func (r *Record) Push(placePlay Point) {
	r.points = append(r.points, placePlay)
}

// Push - 末尾を削除
func (r *Record) Pop(placePlay Point) Point {
	var lastIndex = len(r.points) - 1
	var tail = r.points[lastIndex]
	r.points = r.points[:lastIndex]
	return tail
}

// EOF [O1o1o0g12o__11o_2o0]
