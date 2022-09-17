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
	k.Ren.Color = k.Board.GetColorAt(arbitraryPoint)
	// 隣接する連の色
	k.Ren.AdjacentColor = Color_None

	k.searchRen(arbitraryPoint)

	return k.Ren
}

// 再帰関数。連の探索
func (k *Kernel) searchRen(here Point) {
	k.CheckBoard.Check(here)
	k.Ren.Elements = append(k.Ren.Elements, here)

	// 東、北、西、南
	for dir := 0; dir < 4; dir++ {
		var adjacentP = here + Point(k.Board.Direction[dir]) // 隣接する交点
		// 探索済みならスキップ
		if k.CheckBoard.IsChecked(adjacentP) {
			continue
		}

		var adjacentS = k.Board.GetStoneAt(adjacentP)
		if adjacentS == Space { // 空点
			k.CheckBoard.Check(adjacentP)
			k.Ren.LibertyArea++
			continue
		} else if adjacentS == Wall { // 壁
			continue
		}

		var adjacentC = adjacentS.GetColor()
		// 隣接する色、追加
		k.Ren.AdjacentColor = k.Ren.AdjacentColor.GetAdded(adjacentC)

		if adjacentC == k.Ren.Color { // 同色の石
			k.searchRen(adjacentP) // 再帰
		}
	}
}

// EOF [O1o1o0g22o2o4o0]
