// BOF [O1o1o0g22o2o4o0]

package kernel

// GetLiberty - 呼吸点の数え上げ。連の数え上げ
//
// Parameters
// ----------
// * `arbitraryPoint` - 連に含まれる任意の一点
func (k *Kernel) GetLiberty(arbitraryPoint Point) *Ren {
	// チェックボードの初期化
	k.CheckBoard.Init(k.Board.GetWidth(), k.Board.GetHeight())
	// 連の初期化
	k.Ren = new(Ren)
	// 連の色
	k.Ren.Color = k.Board.GetStoneAt(arbitraryPoint)
	// ４方向（東、北、西、南）の番地への相対インデックス
	k.Direction = [4]int{1, -k.Board.GetMemoryWidth(), -1, k.Board.GetMemoryWidth()}

	k.searchRen(arbitraryPoint)

	return k.Ren
}

// 再帰関数。連の探索
func (k *Kernel) searchRen(here Point) {
	k.CheckBoard.Check(here)
	k.Ren.Area++

	// 東、北、西、南
	for dir := 0; dir < 4; dir++ {
		var adjacent = here + Point(k.Direction[dir]) // 隣接する交点
		// 探索済みならスキップ
		if k.CheckBoard.IsChecked(adjacent) {
			continue
		}
		if k.Board.GetStoneAt(adjacent) == Empty { // 空点
			k.CheckBoard.Check(adjacent)
			k.Ren.LibertyArea++
		} else if k.Board.GetStoneAt(adjacent) == k.Ren.Color { // 同色の石
			k.searchRen(adjacent) // 再帰
		}
	}
}

// EOF [O1o1o0g22o2o4o0]
