// BOF [O1o1o0g22o1o1o0]

package kernel

import "fmt"

func (k *Kernel) IsMasonryError(stone Stone, point Point) bool {
	var target = k.Board.cells[point]
	switch target {
	case Black:
	case White:
	case Wall:
		return true
	case Empty:
		return false
	default:
		panic(fmt.Sprintf("unexpected target cell:%s", target))
	}
	return false
}

// EOF [O1o1o0g22o1o1o0]
