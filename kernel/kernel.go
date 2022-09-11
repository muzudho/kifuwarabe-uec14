// BOF [O1o1o0g11o_3o0]

package kernel

type Kernel struct {
	// Board - 盤
	Board *Board
}

func NewKernel() *Kernel {
	var k = new(Kernel)
	k.Board = NewBoard()
	return k
}

// Execute - 実行
func (k *Kernel) Execute(command string) bool {
	return false
}

// EOF [O1o1o0g11o_3o0]
