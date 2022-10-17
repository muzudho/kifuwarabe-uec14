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
	k.Position.CheckBoard.Init(k.Position.Board.coordinate)

	var libertySearchAlgorithm = NewLibertySearchAlgorithm(k.Position.Board, k.Position.CheckBoard, k.Position.foundRen)

	return libertySearchAlgorithm.findRen(arbitraryPoint)
}

// LibertySearchAlgorithm - 呼吸点探索アルゴリズム
type LibertySearchAlgorithm struct {
	// 盤
	board *Board
	// チェック盤
	checkBoard *CheckBoard
	// foundRen - 呼吸点の探索時に使います
	foundRen *Ren
}

// NewLibertySearchAlgorithm - 新規作成
func NewLibertySearchAlgorithm(board *Board, checkBoard *CheckBoard, foundRen *Ren) *LibertySearchAlgorithm {
	var ls = new(LibertySearchAlgorithm)

	ls.board = board
	ls.checkBoard = checkBoard
	ls.foundRen = foundRen

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
	ls.foundRen = NewRen(ls.board.GetStoneAt(arbitraryPoint))

	if ls.foundRen.stone == Stone_Space {
		ls.searchSpaceRen(arbitraryPoint)
	} else {
		ls.searchStoneRenRecursive(arbitraryPoint)

		// チェックボードの「呼吸点」チェックのみクリアー
		var eachPoint = func(point Point) {
			ls.checkBoard.Erase(point, Mark_BitLiberty)
		}
		ls.board.coordinate.ForeachCellWithoutWall(eachPoint)
	}

	return ls.foundRen, true
}

// 石の連の探索
//
// * 再帰関数
func (ls *LibertySearchAlgorithm) searchStoneRenRecursive(here Point) {

	// 石のチェック
	ls.checkBoard.Overwrite(here, Mark_BitStone)

	ls.foundRen.AddLocation(here)

	// 隣接する交点毎に
	var eachAdjacent = func(dir Cell_4Directions, p Point) {

		var stone = ls.board.GetStoneAt(p) // 石の色
		switch stone {

		case Stone_Space: // 空点
			if !ls.checkBoard.Contains(p, Mark_BitLiberty) { // まだチェックしていない呼吸点なら
				ls.checkBoard.Overwrite(p, Mark_BitLiberty)
				ls.foundRen.libertyLocations = append(ls.foundRen.libertyLocations, p) // 呼吸点を追加
			}

			return // あとの処理をスキップ

		case Stone_Wall: // 枠
			return // あとの処理をスキップ
		}

		// 探索済みの石ならスキップ
		if ls.checkBoard.Contains(p, Mark_BitStone) {
			return
		}

		var color = stone.GetColor()
		// 隣接する色、追加
		ls.foundRen.adjacentColor = ls.foundRen.adjacentColor.GetAdded(color)

		if stone == ls.foundRen.stone { // 同じ石
			ls.searchStoneRenRecursive(p) // 再帰
		}
	}

	// 隣接する４方向
	ls.board.ForeachNeumannNeighborhood(here, eachAdjacent)
}

// 空点の連の探索
// - 再帰関数
func (ls *LibertySearchAlgorithm) searchSpaceRen(here Point) {
	ls.checkBoard.Overwrite(here, Mark_BitStone)
	ls.foundRen.AddLocation(here)

	var eachAdjacent = func(dir Cell_4Directions, p Point) {
		// 探索済みならスキップ
		if ls.checkBoard.Contains(p, Mark_BitStone) {
			return
		}

		var stone = ls.board.GetStoneAt(p)
		if stone != Stone_Space { // 空点でなければスキップ
			return
		}

		var color = stone.GetColor()
		// 隣接する色、追加
		ls.foundRen.adjacentColor = ls.foundRen.adjacentColor.GetAdded(color)
		ls.searchSpaceRen(p) // 再帰
	}

	// 隣接する４方向
	ls.board.ForeachNeumannNeighborhood(here, eachAdjacent)
}

// EOF [O22o2o4o0]
