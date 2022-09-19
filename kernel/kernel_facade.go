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
// * 見つけた連は、連データベースへ入れます
func (k *Kernel) FindAllRens() {
	// チェックボードの初期化
	k.CheckBoard.Init(k.Board.GetWidth(), k.Board.GetHeight())

	var setLocation = func(location Point) {
		var ren, isFound = k.findRen(location)
		if isFound {
			k.renDb.RegisterRen(k.Record.posNum, ren)
		}
	}
	// 盤上の枠の内側をスキャン
	k.Board.ForeachPayloadLocation(setLocation)
}

// EOF [O1o1o0g22o5o1o0]
