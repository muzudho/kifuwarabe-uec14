// BOF [O1o1o0g12o__11o1o0]

package kernel

// Board - 盤
type Board struct {
	// 枠付きの横幅
	memoryWidth int

	// 枠付きの縦幅
	memoryHeight int

	// 交点
	//
	// * 英語で交点は node かも知れないが、表計算でよく使われる cell の方を使う
	cells []Stone
}

// GetWidth - 枠の厚みを含まない横幅
func (b *Board) GetWidth() int {
	return b.memoryWidth - 2
}

// GetHeight - 枠の厚みを含まない縦幅
func (b *Board) GetHeight() int {
	return b.memoryHeight - 2
}

// GetPointFromXy - 座標変換 (x,y) → Point
func (b *Board) GetPointFromXy(x int, y int) Point {
	// 枠の厚み 1 を考慮
	return Point((y+1)*b.memoryWidth + x + 1)
}

// サイズ変更
func (b *Board) resize(width int, height int) {
	b.memoryWidth = width + 2
	b.memoryHeight = height + 2
	b.cells = make([]Stone, b.getMemoryArea())
}

// 枠付き盤の面積
func (b *Board) getMemoryArea() int {
	return b.memoryWidth * b.memoryHeight
}

// EOF [O1o1o0g12o__11o1o0]
