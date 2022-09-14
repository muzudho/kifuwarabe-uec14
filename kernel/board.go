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

// GetMemoryWidth - 枠の厚みを含んだ横幅
func (b *Board) GetMemoryWidth() int {
	return b.memoryWidth
}

// GetMemoryHeight - 枠の厚みを含んだ縦幅
func (b *Board) GetMemoryHeight() int {
	return b.memoryHeight
}

// GetWidth - 枠の厚みを含まない横幅
func (b *Board) GetWidth() int {
	return b.memoryWidth - 2
}

// GetHeight - 枠の厚みを含まない縦幅
func (b *Board) GetHeight() int {
	return b.memoryHeight - 2
}

// GetStoneAt - 指定座標の石を取得
func (b *Board) GetStoneAt(i Point) Stone {
	return b.cells[i]
}

// GetPointFromXy - 座標変換 (x,y) → Point
//
// Parameters
// ----------
// x : int
//	筋番号。 Example: 19路盤なら0～18
// y : int
//	段番号。 Example: 19路盤なら0～18
//
// Returns
// -------
// point : Point
//  配列インデックス。 Example: 2,3 なら 65
func (b *Board) GetPointFromXy(x int, y int) Point {
	return Point(y*b.memoryWidth + x)
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
