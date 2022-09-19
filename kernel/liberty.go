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
	// 探索済みならスキップ
	if k.CheckBoard.IsStoneChecked(arbitraryPoint) {
		return nil, false
	}

	// 連の初期化
	k.tempRen = NewRen(k.Board.GetStoneAt(arbitraryPoint))

	if k.tempRen.stone == Space {
		k.searchSpaceRen(arbitraryPoint)
	} else {
		k.searchStoneRen(arbitraryPoint)

		// チェックボードの呼吸点のチェックをクリアー
		for _, p := range k.tempRen.libertyLocations {
			k.CheckBoard.UncheckLiberty(p)
		}
	}

	return k.tempRen, true
}

// 石の連の探索
// - 再帰関数
func (k *Kernel) searchStoneRen(here Point) {
	k.CheckBoard.CheckStone(here)
	k.tempRen.AddLocation(here)

	var setAdjacent = func(dir int, p Point) {
		// 呼吸点と壁のチェック
		var stone = k.Board.GetStoneAt(p)
		switch stone {
		case Space: // 空点
			if !k.CheckBoard.IsLibertyChecked(p) { // まだチェックしていない呼吸点なら
				k.CheckBoard.CheckLiberty(p)
				k.tempRen.libertyLocations = append(k.tempRen.libertyLocations, p) // 呼吸点を追加
			}

			return // あとの処理をスキップ

		case Wall: // 壁
			return // あとの処理をスキップ
		}

		// 探索済みの石ならスキップ
		if k.CheckBoard.IsStoneChecked(p) {
			return
		}

		var color = stone.GetColor()
		// 隣接する色、追加
		k.tempRen.adjacentColor = k.tempRen.adjacentColor.GetAdded(color)

		if stone == k.tempRen.stone { // 同じ石
			k.searchStoneRen(p) // 再帰
		}
	}

	// 隣接する４方向
	k.Board.ForeachNeumannNeighborhood(here, setAdjacent)
}

// 空点の連の探索
// - 再帰関数
func (k *Kernel) searchSpaceRen(here Point) {
	k.CheckBoard.CheckStone(here)
	k.tempRen.AddLocation(here)

	var setAdjacent = func(dir int, p Point) {
		// 探索済みならスキップ
		if k.CheckBoard.IsStoneChecked(p) {
			return
		}

		var stone = k.Board.GetStoneAt(p)
		if stone != Space { // 空点でなければスキップ
			return
		}

		var color = stone.GetColor()
		// 隣接する色、追加
		k.tempRen.adjacentColor = k.tempRen.adjacentColor.GetAdded(color)
		k.searchSpaceRen(p) // 再帰
	}

	// 隣接する４方向
	k.Board.ForeachNeumannNeighborhood(here, setAdjacent)
}

// EOF [O1o1o0g22o2o4o0]
