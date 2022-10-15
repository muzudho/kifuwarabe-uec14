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
			b.cells[i] = Stone_Wall

			i = b.coordinate.GetPointFromXy(x, y2)
			b.cells[i] = Stone_Wall
		}
	}
	// 枠の左辺、右辺を引く
	{
		var x = 0
		var x2 = b.coordinate.memoryWidth - 1
		for y := 1; y < b.coordinate.memoryHeight-1; y++ {
			var i = b.coordinate.GetPointFromXy(x, y)
			b.cells[i] = Stone_Wall

			i = b.coordinate.GetPointFromXy(x2, y)
			b.cells[i] = Stone_Wall
		}
	}
	// 枠の内側を空点で埋める
	{
		var height = b.coordinate.GetHeight()
		var width = b.coordinate.GetWidth()
		for y := 1; y < height; y++ {
			for x := 1; x < width; x++ {
				var i = b.coordinate.GetPointFromXy(x, y)
				b.cells[i] = Stone_Space
			}
		}
	}
}

// ForeachNeumannNeighborhood - [O13o__10o0] 隣接する４方向の定義
func (b *Board) ForeachNeumannNeighborhood(here Point, setAdjacent func(Cell_4Directions, Point)) {
	// 東、北、西、南
	for dir := Cell_4Directions(0); dir < 4; dir++ {
		var p = here + b.coordinate.cell4Directions[dir] // 隣接する交点

		// 範囲外チェック
		if p < 0 || b.coordinate.GetMemoryArea() <= int(p) {
			continue
		}

		// 枠チェック
		if b.GetStoneAt(p) == Stone_Wall {
			continue
		}

		setAdjacent(dir, p)
	}
}

// EOF [O12o0]
