// BOF [O1o1o0g22o1o1o0]

package kernel

import "fmt"

func (k *Kernel) IsMasonryError(stone Stone, point Point) bool {
	var target = k.Board.cells[point]

	switch target {
	case Empty:
		return false
	case Black, White, Wall:
		return true
	default:
		panic(fmt.Sprintf("unexpected target cell:%s", target))
	}
}

// EOF [O1o1o0g22o1o1o0]
