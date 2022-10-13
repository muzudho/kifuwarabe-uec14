// BOF [O12o0]

package kernel

// Init - 盤面初期化
func (b *Board) Init(width int, height int) {
	// 盤面のサイズが異なるなら、盤面を作り直す
	if b.coordinate.memoryWidth != width || b.coordinate.memoryHeight != height {
		b.resize(width, height)
	}

	// 枠の上辺、下辺を引く
	{
		var y = 0
		var y2 = b.coordinate.memoryHeight - 1
		for x := 0; x < b.coordinate.memoryWidth; x++ {
			var i = b.coordinate.GetPointFromXy(x, y)
			b.cells[i] = Wall

			i = b.coordinate.GetPointFromXy(x, y2)
			b.cells[i] = Wall
		}
	}
	// 枠の左辺、右辺を引く
	{
		var x = 0
		var x2 = b.coordinate.memoryWidth - 1
		for y := 1; y < b.coordinate.memoryHeight-1; y++ {
			var i = b.coordinate.GetPointFromXy(x, y)
			b.cells[i] = Wall

			i = b.coordinate.GetPointFromXy(x2, y)
			b.cells[i] = Wall
		}
	}
	// 枠の内側を空点で埋める
	{
		var height = b.coordinate.GetBoardHeight()
		var width = b.coordinate.GetWidth()
		for y := 1; y < height; y++ {
			for x := 1; x < width; x++ {
				var i = b.coordinate.GetPointFromXy(x, y)
				b.cells[i] = Space
			}
		}
	}
}

// ForeachLikeText - 枠を含めた各セルの石
func (b *Board) ForeachLikeText(setStone func(Stone), doNewline func()) {
	for y := 0; y < b.coordinate.memoryHeight; y++ {
		if y != 0 {
			doNewline()
		}

		for x := 0; x < b.coordinate.memoryWidth; x++ {
			var i = b.coordinate.GetPointFromXy(x, y)
			var stone = b.cells[i]
			setStone(stone)
		}
	}
}

// ForeachPayloadLocation - 枠や改行を含めない各セルの番地
func (b *Board) ForeachPayloadLocation(setLocation func(Point)) {
	var height = b.coordinate.memoryHeight - 1
	var width = b.coordinate.memoryWidth - 1
	for y := 1; y < height; y++ {
		for x := 1; x < width; x++ {
			var i = b.coordinate.GetPointFromXy(x, y)
			setLocation(i)
		}
	}
}

// ForeachPayloadLocation - 枠や改行を含めない各セルの番地。筋、段の順
func (b *Board) ForeachPayloadLocationOrderByYx(setLocation func(Point)) {
	var height = b.coordinate.memoryHeight - 1
	var width = b.coordinate.memoryWidth - 1
	for x := 1; x < width; x++ {
		for y := 1; y < height; y++ {
			var i = b.coordinate.GetPointFromXy(x, y)
			setLocation(i)
		}
	}
}

// ForeachNeumannNeighborhood - [O13o__10o0] 隣接する４方向の定義
func (b *Board) ForeachNeumannNeighborhood(here Point, setAdjacent func(int, Point)) {
	// 東、北、西、南
	for dir := 0; dir < 4; dir++ {
		var p = here + b.coordinate.cell4Directions[dir] // 隣接する交点

		// 範囲外チェック
		if p < 0 || b.coordinate.GetMemoryArea() <= int(p) {
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
