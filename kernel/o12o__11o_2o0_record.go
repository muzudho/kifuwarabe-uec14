// BOF [O12o__11o_2o0]

package kernel

import (
	"math"
	"strconv"
)

// Record - 棋譜
type Record struct {
	// 先行
	playFirst Stone

	// 何手目。基数（Position number）
	posNum int

	// 手毎
	items []*RecordItem
}

// NewRecord - 新規作成
//
// * maxPositionNumber - 手数上限。配列サイズ決定のための判断材料
// * memoryBoardArea - メモリー盤サイズ。配列サイズ決定のための判断材料
func NewRecord(maxPositionNumber PositionNumberInt, memoryBoardArea int, playFirst Stone) *Record {
	var r = new(Record)
	r.playFirst = playFirst

	// 動的に長さがきまる配列を生成、その内容をインスタンスで埋めます
	// 例えば、0手目が初期局面として、 400 手目まであるとすると、要素数は401要る。だから 1 足す
	// しかし、プレイアウトでは終局まで打ちたいので、多めにとっておきたいのでは。盤サイズより適当に18倍（>2πe）取る
	var positionLength = int(math.Max(float64(maxPositionNumber+1), float64(memoryBoardArea*18)))
	r.items = make([]*RecordItem, positionLength)

	for i := PositionNumberInt(0); i < PositionNumberInt(positionLength); i++ {
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
	// [O22o7o1o0] コウの位置
	ko Point) {

	var item = r.items[r.posNum]
	item.placePlay = placePlay

	// [O22o7o1o0] コウの位置
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
	// [O22o7o1o0] コウの判定
	// 2手前に着手して石をぴったり１つ打ち上げたとき、その着手点はコウだ
	var posNum = r.GetPositionNumber()
	if 2 <= posNum {
		var item = r.items[posNum-2]
		return item.ko == placePlay
	}

	return false
}

// EOF [O12o__11o_2o0]
