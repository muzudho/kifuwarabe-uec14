// BOF [O12o0]

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

// ForeachLikeText - 枠を含めた各セルの石
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

// ForeachPayloadLocation - 枠や改行を含めない各セルの番地
func (b *Board) ForeachPayloadLocation(setLocation func(Point)) {
	var height = b.memoryHeight - 1
	var width = b.memoryWidth - 1
	for y := 1; y < height; y++ {
		for x := 1; x < width; x++ {
			var i = b.GetPointFromXy(x, y)
			setLocation(i)
		}
	}
}

// ForeachPayloadLocation - 枠や改行を含めない各セルの番地。筋、段の順
func (b *Board) ForeachPayloadLocationOrderByYx(setLocation func(Point)) {
	var height = b.memoryHeight - 1
	var width = b.memoryWidth - 1
	for x := 1; x < width; x++ {
		for y := 1; y < height; y++ {
			var i = b.GetPointFromXy(x, y)
			setLocation(i)
		}
	}
}

// ForeachNeumannNeighborhood - [O13o__10o0] 隣接する４方向の定義
func (b *Board) ForeachNeumannNeighborhood(here Point, setAdjacent func(int, Point)) {
	// 東、北、西、南
	for dir := 0; dir < 4; dir++ {
		var p = here + Point(b.Direction[dir]) // 隣接する交点

		// 範囲外チェック
		if p < 0 || b.getMemoryArea() <= int(p) {
			continue
		}

		// 壁チェック
		if b.GetStoneAt(p) == Wall {
			continue
		}

		setAdjacent(dir, p)
	}
}

// EOF [O12o0]
