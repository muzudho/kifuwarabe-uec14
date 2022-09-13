// BOF [O1o1o0g12o0]

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

// NewBoard - 新規作成
func NewBoard() *Board {
	var b = new(Board)

	// 盤のサイズ指定と、盤面の初期化
	b.Resize(19, 19)

	return b
}

// Resize - サイズ変更
func (b *Board) Resize(width int, height int) {
	b.memoryWidth = width + 2
	b.memoryHeight = height + 2
	b.cells = make([]Stone, b.getMemoryArea())

	// 枠を設定する
	// 上辺、下辺を引く
	{
		var y = 0
		var y2 = b.memoryHeight - 1
		for x := 0; x < b.memoryWidth; x++ {
			var i = (y * b.memoryWidth) + x
			b.cells[i] = Wall

			i = (y2 * b.memoryWidth) + x
			b.cells[i] = Wall
		}
	}
	// 左辺、右辺を引く
	{
		var x = 0
		var x2 = b.memoryWidth - 1
		for y := 1; y < b.memoryHeight-1; y++ {
			var i = (y * b.memoryWidth) + x
			b.cells[i] = Wall

			i = (y * b.memoryWidth) + x2
			b.cells[i] = Wall
		}
	}
}

// ForeachLikeText - 枠を含めた各セル
func (b *Board) ForeachLikeText(setStone func(Stone), doNewline func()) {
	for y := 0; y < b.memoryHeight; y++ {
		if y != 0 {
			doNewline()
		}

		for x := 0; x < b.memoryWidth; x++ {
			var i = (y * b.memoryWidth) + x
			var stone = b.cells[i]
			setStone(stone)
		}
	}
}

// 枠付き盤の面積
func (b *Board) getMemoryArea() int {
	return b.memoryWidth * b.memoryHeight
}

// EOF [O1o1o0g12o0]
