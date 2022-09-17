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

	// ko - [O1o1o0g22o7o1o0] コウの位置
	ko []Point
}

// NewRecord - 棋譜の新規作成
func NewRecord(maxMoves int, playFirst Stone) *Record {
	var r = new(Record)
	r.playFirst = playFirst
	r.points = make([]Point, maxMoves)

	// [O1o1o0g22o7o1o0] コウの位置
	r.ko = make([]Point, maxMoves)

	return r
}

// GetCurrent - 現在位置
func (r *Record) GetCurrent() int {
	return r.current
}

// Push - 末尾に追加
func (r *Record) Push(placePlay Point,
	// [O1o1o0g22o7o1o0] コウの位置
	ko Point) {
	r.points[r.current] = placePlay

	// [O1o1o0g22o7o1o0] コウの位置
	r.ko[r.current] = ko

	r.current++
}

// Push - 末尾を削除
func (r *Record) Pop(placePlay Point) Point {
	r.current--
	var tail = r.points[r.current]
	r.points[r.current] = Point(0)
	return tail
}

// Foreach - 各要素
func (r *Record) Foreach(setPoint func(int, Point)) {
	for i := 0; i < r.current; i++ {
		setPoint(i, r.points[i])
	}
}

// IsKo - コウか？
func (r *Record) IsKo(placePlay Point) bool {
	// [O1o1o0g22o7o1o0] コウの判定
	// 2手前に着手して石をぴったり１つ打ち上げたとき、その着手点はコウだ
	var i = r.GetCurrent()
	return 1 <= i && r.ko[i-2] == placePlay
}

// EOF [O1o1o0g12o__11o_2o0]
