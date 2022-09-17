// BOF [O1o1o0g12o__11o_2o0]

package kernel

// Record - 棋譜
type Record struct {
	// 先行
	playFirst Stone

	// 現在位置
	current int

	// 手毎
	items []*RecordItem
}

// NewRecord - 棋譜の新規作成
func NewRecord(maxMoves int, playFirst Stone) *Record {
	var r = new(Record)
	r.playFirst = playFirst

	// 棋譜の一手分毎
	r.items = make([]*RecordItem, maxMoves)
	for i := 0; i < maxMoves; i++ {
		r.items[i] = NewRecordItem()
	}

	return r
}

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

// GetCurrent - 現在位置
func (r *Record) GetCurrent() int {
	return r.current
}

// Push - 末尾に追加
func (r *Record) Push(placePlay Point,
	// [O1o1o0g22o7o1o0] コウの位置
	ko Point) {

	var item = r.items[r.current]
	item.placePlay = placePlay

	// [O1o1o0g22o7o1o0] コウの位置
	item.ko = ko

	r.current++
}

// Push - 末尾を削除
func (r *Record) Pop(placePlay Point) *RecordItem {
	r.current--
	var tail = r.items[r.current]
	r.items[r.current] = NewRecordItem()
	return tail
}

// ForeachItem - 各要素
func (r *Record) ForeachItem(setItem func(int, *RecordItem)) {
	for i := 0; i < r.current; i++ {
		setItem(i, r.items[i])
	}
}

// IsKo - コウか？
func (r *Record) IsKo(placePlay Point) bool {
	// [O1o1o0g22o7o1o0] コウの判定
	// 2手前に着手して石をぴったり１つ打ち上げたとき、その着手点はコウだ
	var i = r.GetCurrent()
	if 2 <= i {
		var item = r.items[i-2]
		return item.ko == placePlay
	}

	return false
}

// EOF [O1o1o0g12o__11o_2o0]
