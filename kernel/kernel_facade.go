// BOF [O1o1o0g22o5o1o0]

package kernel

// RemoveRen - 石の連を打ち上げます
func (k *Kernel) RemoveRen(ren *Ren) {
	// 空点をセット
	var setLocation = func(i int, location Point) {
		k.Board.SetStoneAt(location, Space)
	}

	// 場所毎に
	ren.ForeachLocation(setLocation)
}

// FindAllRens - [O1o1o0g23o_2o1o0] 盤上の全ての連を見つけます
func (k *Kernel) FindAllRens() {
	// チェックボードの初期化
	k.CheckBoard.Init(k.Board.GetWidth(), k.Board.GetHeight())

	//
}

// EOF [O1o1o0g22o5o1o0]
