// BOF [O1o1o0g12o0]

package kernel

// Init - 盤面初期化
func (b *Board) Init(width int, height int) {
	// 盤面のサイズが異なるなら、盤面を作り直す
	if b.memoryWidth != width || b.memoryHeight != height {
		b.resize(width, height)
	}

	// 枠の上辺、下辺を引く
	{
		var y = 0
		var y2 = b.memoryHeight - 1
		for x := 0; x < b.memoryWidth; x++ {
			var i = b.GetPointFromXy(x, y)
			b.cells[i] = Wall

			i = b.GetPointFromXy(x, y2)
			b.cells[i] = Wall
		}
	}
	// 枠の左辺、右辺を引く
	{
		var x = 0
		var x2 = b.memoryWidth - 1
		for y := 1; y < b.memoryHeight-1; y++ {
			var i = b.GetPointFromXy(x, y)
			b.cells[i] = Wall

			i = b.GetPointFromXy(x2, y)
			b.cells[i] = Wall
		}
	}
	// 枠の内側を空点で埋める
	{
		var height = b.GetHeight()
		var width = b.GetWidth()
		for y := 1; y < height; y++ {
			for x := 1; x < width; x++ {
				var i = b.GetPointFromXy(x, y)
				b.cells[i] = Space
			}
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
			var i = b.GetPointFromXy(x, y)
			var stone = b.cells[i]
			setStone(stone)
		}
	}
}

// ForeachNeumannNeighborhood - [O1o1o0g13o__10o0] 隣接する４方向の定義
func (b *Board) ForeachNeumannNeighborhood() {
	// 東、北、西、南
	for dir := 0; dir < 4; dir++ {
		//var adjacentP = here + Point(k.Board.Direction[dir]) // 隣接する交点
	}
}

// EOF [O1o1o0g12o0]
