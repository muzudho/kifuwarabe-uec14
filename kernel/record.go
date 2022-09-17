// BOF [O1o1o0g12o__11o_2o0]

package kernel

// Record - 棋譜
type Record struct {
	// 先行
	playFirst Stone

	// 着手点
	points []Point

	// 現在位置
	current int
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
	r.points[r.current] = placePlay
	r.current++
}

// Push - 末尾を削除
func (r *Record) Pop(placePlay Point) Point {
	r.current--
	var tail = r.points[r.current]
	r.points[r.current] = Point(0)
	return tail
}

// EOF [O1o1o0g12o__11o_2o0]
