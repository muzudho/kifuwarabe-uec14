// BOF [O1o1o0g12o__11o_2o0]

package kernel

import "strconv"

// Record - 棋譜
type Record struct {
	// 先行
	playFirst Stone

	// 何手目。基数（Position number）
	posNum int

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

// GetMaxPosNthFigure - 手数（序数）の最大値の桁数
func (r *Record) GetMaxPosNthFigure() int {
	var nth = r.GetMaxPosNth()
	var nthText = strconv.Itoa(nth)
	return len(nthText)
}

// GetMaxPosNth - 手数（序数）の最大値
func (r *Record) GetMaxPosNth() int {
	return len(r.items) + geta
}

// GetPositionNumber - 何手目。基数
func (r *Record) GetPositionNumber() int {
	return r.posNum
}

// Push - 末尾に追加
func (r *Record) Push(placePlay Point,
	// [O1o1o0g22o7o1o0] コウの位置
	ko Point) {

	var item = r.items[r.posNum]
	item.placePlay = placePlay

	// [O1o1o0g22o7o1o0] コウの位置
	item.ko = ko

	r.posNum++
}

// RemoveTail - 末尾を削除
func (r *Record) RemoveTail(placePlay Point) {
	r.posNum--
	r.items[r.posNum].Clear()
}

// ForeachItem - 各要素
func (r *Record) ForeachItem(setItem func(int, *RecordItem)) {
	for i := 0; i < r.posNum; i++ {
		setItem(i, r.items[i])
	}
}

// IsKo - コウか？
func (r *Record) IsKo(placePlay Point) bool {
	// [O1o1o0g22o7o1o0] コウの判定
	// 2手前に着手して石をぴったり１つ打ち上げたとき、その着手点はコウだ
	var posNum = r.GetPositionNumber()
	if 2 <= posNum {
		var item = r.items[posNum-2]
		return item.ko == placePlay
	}

	return false
}

// EOF [O1o1o0g12o__11o_2o0]
