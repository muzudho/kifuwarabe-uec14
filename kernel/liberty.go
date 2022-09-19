// BOF [O1o1o0g22o2o4o0]

package kernel

// GetLiberty - 呼吸点の数え上げ。連の数え上げ。
// `GetOneRen` とでもいう名前の方がふさわしいが、慣習に合わせた関数名にした
//
// Parameters
// ----------
// * `arbitraryPoint` - 連に含まれる任意の一点
//
// Returns
// -------
// - *Ren is ren or nil
// - bool is found
func (k *Kernel) GetLiberty(arbitraryPoint Point) (*Ren, bool) {
	// チェックボードの初期化
	k.CheckBoard.Init(k.Board.GetWidth(), k.Board.GetHeight())
	return k.findRen(arbitraryPoint)
}

// 連の検索
//
// Returns
// -------
// - *Ren is ren or nil
// - bool is found
func (k *Kernel) findRen(arbitraryPoint Point) (*Ren, bool) {
	// 連の初期化
	k.tempRen = NewRen(k.Board.GetColorAt(arbitraryPoint))
	var stone = k.Board.GetStoneAt(arbitraryPoint)

	// 探索済みならスキップ
	if k.CheckBoard.IsChecked(arbitraryPoint) {
		return nil, false
	}

	if stone == Space {
		k.searchSpaceRen(arbitraryPoint)
	} else {
		k.searchStoneRen(arbitraryPoint)
	}

	return k.tempRen, true
}

// 石の連の探索
// - 再帰関数
func (k *Kernel) searchStoneRen(here Point) {
	k.CheckBoard.Check(here)
	k.tempRen.AddLocation(here)

	var setAdjacentPoint = func(dir int, adjacentP Point) {
		// 探索済みならスキップ
		if k.CheckBoard.IsChecked(adjacentP) {
			return
		}

		var adjacentS = k.Board.GetStoneAt(adjacentP)
		switch adjacentS {
		case Space: // 空点
			// k.CheckBoard.Check(adjacentP)
			k.tempRen.LibertyArea++ //呼吸点を数え上げる
			return                  // スキップ
		case Wall: // 壁
			return
		}

		var adjacentC = adjacentS.GetColor()
		// 隣接する色、追加
		k.tempRen.adjacentColor = k.tempRen.adjacentColor.GetAdded(adjacentC)

		if adjacentC == k.tempRen.color { // 同色の石
			k.searchStoneRen(adjacentP) // 再帰
		}
	}

	// 隣接する４方向
	k.Board.ForeachNeumannNeighborhood(here, setAdjacentPoint)
}

// 空点の連の探索
// - 再帰関数
func (k *Kernel) searchSpaceRen(here Point) {
	k.CheckBoard.Check(here)
	k.tempRen.AddLocation(here)

	var setAdjacentPoint = func(dir int, adjacentP Point) {
		// 探索済みならスキップ
		if k.CheckBoard.IsChecked(adjacentP) {
			return
		}

		var adjacentS = k.Board.GetStoneAt(adjacentP)
		if adjacentS != Space { // 空点でなければスキップ
			return
		}

		var adjacentC = adjacentS.GetColor()
		// 隣接する色、追加
		k.tempRen.adjacentColor = k.tempRen.adjacentColor.GetAdded(adjacentC)
		k.searchSpaceRen(adjacentP) // 再帰
	}

	// 隣接する４方向
	k.Board.ForeachNeumannNeighborhood(here, setAdjacentPoint)
}

// EOF [O1o1o0g22o2o4o0]
