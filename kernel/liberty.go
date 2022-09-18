// BOF [O1o1o0g22o2o4o0]

package kernel

// GetLiberty - 呼吸点の数え上げ。連の数え上げ。
// `GetOneRen` とでもいう名前の方がふさわしいが、慣習に合わせた関数名にした
//
// Parameters
// ----------
// * `arbitraryPoint` - 連に含まれる任意の一点
func (k *Kernel) GetLiberty(arbitraryPoint Point) *Ren {
	// チェックボードの初期化
	k.CheckBoard.Init(k.Board.GetWidth(), k.Board.GetHeight())

	return k.getRen(arbitraryPoint)
}

// 連の取得
func (k *Kernel) getRen(arbitraryPoint Point) *Ren {
	// 連の初期化
	k.tempRen = NewRen()
	// 連の色
	k.tempRen.Color = k.Board.GetColorAt(arbitraryPoint)
	// 隣接する連の色
	k.tempRen.AdjacentColor = Color_None

	k.searchRen(arbitraryPoint)

	return k.tempRen
}

// 再帰関数。連の探索
func (k *Kernel) searchRen(here Point) {
	k.CheckBoard.Check(here)
	k.tempRen.AddLocation(here)

	var setAdjacentPoint = func(dir int, adjacentP Point) {
		// 探索済みならスキップ
		if k.CheckBoard.IsChecked(adjacentP) {
			return
		}

		var adjacentS = k.Board.GetStoneAt(adjacentP)
		if adjacentS == Space { // 空点
			k.CheckBoard.Check(adjacentP)
			k.tempRen.LibertyArea++
			return
		} else if adjacentS == Wall { // 壁
			return
		}

		var adjacentC = adjacentS.GetColor()
		// 隣接する色、追加
		k.tempRen.AdjacentColor = k.tempRen.AdjacentColor.GetAdded(adjacentC)

		if adjacentC == k.tempRen.Color { // 同色の石
			k.searchRen(adjacentP) // 再帰
		}
	}

	// 隣接する４方向
	k.Board.ForeachNeumannNeighborhood(here, setAdjacentPoint)
}

// EOF [O1o1o0g22o2o4o0]
