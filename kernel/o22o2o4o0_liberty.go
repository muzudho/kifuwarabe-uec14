// BOF [O22o2o4o0]

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
	k.CheckBoard.Init(k.Board.coordinate)

	var libertySearchAlgorithm = NewLibertySearchAlgorithm(k.Board, k.CheckBoard, k.tempRen)

	return libertySearchAlgorithm.findRen(arbitraryPoint)
}

// LibertySearchAlgorithm - 呼吸点探索アルゴリズム
type LibertySearchAlgorithm struct {
	// 盤
	board *Board
	// チェック盤
	checkBoard *CheckBoard
	// tempRen - 呼吸点の探索時に使います
	tempRen *Ren
}

// NewLibertySearchAlgorithm - 新規作成
func NewLibertySearchAlgorithm(board *Board, checkBoard *CheckBoard, tempRen *Ren) *LibertySearchAlgorithm {
	var ls = new(LibertySearchAlgorithm)

	ls.board = board
	ls.checkBoard = checkBoard
	ls.tempRen = tempRen

	return ls
}

// 連の検索
//
// Returns
// -------
// - *Ren is ren or nil
// - bool is found
func (ls *LibertySearchAlgorithm) findRen(arbitraryPoint Point) (*Ren, bool) {
	// 探索済みならスキップ
	if ls.checkBoard.Contains(arbitraryPoint, Mark_BitStone) {
		return nil, false
	}

	// 連の初期化
	ls.tempRen = NewRen(ls.board.GetStoneAt(arbitraryPoint))

	if ls.tempRen.stone == Space {
		ls.searchSpaceRen(arbitraryPoint)
	} else {
		ls.searchStoneRen(arbitraryPoint)

		// チェックボードの呼吸点のチェックをクリアー
		for _, p := range ls.tempRen.libertyLocations {
			ls.checkBoard.Erase(p, Mark_BitLiberty)
		}
	}

	return ls.tempRen, true
}

// 石の連の探索
// - 再帰関数
func (ls *LibertySearchAlgorithm) searchStoneRen(here Point) {
	ls.checkBoard.Overwrite(here, Mark_BitStone)
	ls.tempRen.AddLocation(here)

	var setAdjacent = func(dir int, p Point) {
		// 呼吸点と枠のチェック
		var stone = ls.board.GetStoneAt(p)
		switch stone {
		case Space: // 空点
			if !ls.checkBoard.Contains(p, Mark_BitLiberty) { // まだチェックしていない呼吸点なら
				ls.checkBoard.Overwrite(p, Mark_BitLiberty)
				ls.tempRen.libertyLocations = append(ls.tempRen.libertyLocations, p) // 呼吸点を追加
			}

			return // あとの処理をスキップ

		case Wall: // 枠
			return // あとの処理をスキップ
		}

		// 探索済みの石ならスキップ
		if ls.checkBoard.Contains(p, Mark_BitStone) {
			return
		}

		var color = stone.GetColor()
		// 隣接する色、追加
		ls.tempRen.adjacentColor = ls.tempRen.adjacentColor.GetAdded(color)

		if stone == ls.tempRen.stone { // 同じ石
			ls.searchStoneRen(p) // 再帰
		}
	}

	// 隣接する４方向
	ls.board.ForeachNeumannNeighborhood(here, setAdjacent)
}

// 空点の連の探索
// - 再帰関数
func (ls *LibertySearchAlgorithm) searchSpaceRen(here Point) {
	ls.checkBoard.Overwrite(here, Mark_BitStone)
	ls.tempRen.AddLocation(here)

	var setAdjacent = func(dir int, p Point) {
		// 探索済みならスキップ
		if ls.checkBoard.Contains(p, Mark_BitStone) {
			return
		}

		var stone = ls.board.GetStoneAt(p)
		if stone != Space { // 空点でなければスキップ
			return
		}

		var color = stone.GetColor()
		// 隣接する色、追加
		ls.tempRen.adjacentColor = ls.tempRen.adjacentColor.GetAdded(color)
		ls.searchSpaceRen(p) // 再帰
	}

	// 隣接する４方向
	ls.board.ForeachNeumannNeighborhood(here, setAdjacent)
}

// EOF [O22o2o4o0]
