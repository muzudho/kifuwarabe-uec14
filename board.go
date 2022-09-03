// BOF [O1o1o0g12o0]

package main

// MemoryWidth - 枠付きの横幅
const MemoryWidth int = 19 + 2

// MemoryHeight - 枠付きの縦幅
const MemoryHeight int = 19 + 2

// MemoryArea - 枠付きの面積
const MemoryArea int = MemoryWidth * MemoryHeight

// Board - 盤
type Board struct {
	// 交点
	nodes []Stone
}

// NewBoard - 新規作成
func NewBoard() *Board {
	var b = new(Board)

	b.nodes = make([]Stone, MemoryArea)

	// 枠を設定する
	// 上辺、下辺を引く
	{
		var y = 0
		var y2 = MemoryHeight - 1
		for x := 0; x < MemoryWidth; x++ {
			var i = (y * MemoryWidth) + x
			b.nodes[i] = Wall

			i = (y2 * MemoryWidth) + x
			b.nodes[i] = Wall
		}
	}
	// 左辺、右辺を引く
	{
		var x = 0
		var x2 = MemoryWidth - 1
		for y := 1; y < MemoryHeight-1; y++ {
			var i = (y * MemoryWidth) + x
			b.nodes[i] = Wall

			i = (y * MemoryWidth) + x2
			b.nodes[i] = Wall
		}
	}

	return b
}

// ForeachLikeText - 枠を含めた各セル
func (b *Board) ForeachLikeText(setStone func(Stone), doNewline func()) {
	for y := 0; y < MemoryHeight; y++ {
		if y != 0 {
			doNewline()
		}

		for x := 0; x < MemoryWidth; x++ {
			var i = (y * MemoryWidth) + x
			var stone = b.nodes[i]
			setStone(stone)
		}
	}
}

// EOF [O1o1o0g12o0]
