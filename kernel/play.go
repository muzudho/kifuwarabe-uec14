// BOF [O1o1o0g19o0]

package kernel

import "strings"

// DoPlay - 打つ
//
// * `command` - Example: `play black A19`
//                         ---- ----- ---
//                         0    1     2
func (k *Kernel) DoPlay(command string, logg *Logger) {
	var tokens = strings.Split(command, " ")
	var stoneName = tokens[1]

	var getDefaultStone = func() (bool, Stone) {
		logg.C.Infof("? unexpected stone:%s\n", stoneName)
		logg.J.Infow("error", "stone", stoneName)
		return false, Space
	}

	var isOk1, stone = GetStoneFromName(stoneName, getDefaultStone)
	if !isOk1 {
		return
	}

	var coord = tokens[2]
	var point = k.Board.GetPointFromCode(coord)

	// [O1o1o0g22o1o2o0] 石（または壁）の上に石を置こうとしました
	var onMasonry = func() bool {
		logg.C.Infof("? masonry my_stone:%s point:%s\n", stone, k.Board.GetCodeFromPoint(point))
		logg.J.Infow("error masonry", "my_stone", stone, "point", k.Board.GetCodeFromPoint(point))
		return false
	}

	// [O1o1o0g22o3o1o0] 相手の眼に石を置こうとしました
	var onOpponentEye = func() bool {
		logg.C.Infof("? opponent_eye my_stone:%s point:%s\n", stone, k.Board.GetCodeFromPoint(point))
		logg.J.Infow("error opponent_eye", "my_stone", stone, "point", k.Board.GetCodeFromPoint(point))
		return false
	}

	// [O1o1o0g22o4o1o0] 自分の眼に石を置こうとしました
	var onForbiddenMyEye = func() bool {
		logg.C.Infof("? my_eye my_stone:%s point:%s\n", stone, k.Board.GetCodeFromPoint(point))
		logg.J.Infow("error my_eye", "my_stone", stone, "point", k.Board.GetCodeFromPoint(point))
		return false
	}

	var isOk = k.Play(stone, point, logg,
		// [O1o1o0g22o1o2o0] ,onMasonry
		onMasonry,
		// [O1o1o0g22o3o1o0] ,onOpponentEye
		onOpponentEye,
		// [O1o1o0g22o4o1o0] ,onForbiddenMyEye
		onForbiddenMyEye)

	if isOk {
		logg.C.Info("=\n")
		logg.J.Infow("ok")
	}
}

// Play - 石を打つ
//
// Returns
// -------
// isOk : bool
//		石を置けたら真、置けなかったら偽
func (k *Kernel) Play(stoneA Stone, pointB Point, logg *Logger,
	// [O1o1o0g22o1o2o0] onMasonry
	onMasonry func() bool,
	// [O1o1o0g22o3o1o0] onOpponentEye
	onOpponentEye func() bool,
	// [O1o1o0g22o4o1o0] onForbiddenMyEye
	onForbiddenMyEye func() bool) bool {

	// [O1o1o0g22o1o2o0]
	if k.IsMasonryError(stoneA, pointB) {
		return onMasonry()
	}

	// [O1o1o0g22o7o2o0] コウの判定
	if k.Record.IsKo(pointB) {

	}

	// [O1o1o0g22o6o1o0] Captured ルール
	var isExists4rensToRemove = false
	var o4rensToRemove [4]*Ren
	var isChecked4rensToRemove = false

	// [O1o1o0g22o3o1o0]
	var renC = k.GetLiberty(pointB)
	if renC.GetArea() == 1 { // 石Aを置いた交点を含む連Cについて、連Cの面積が1である（眼）
		if stoneA.GetColor() == renC.AdjacentColor.GetOpponent() {
			// かつ、連Cに隣接する連の色が、石Aのちょうど反対側の色であったなら、
			// 相手の眼に石を置こうとしたとみなす

			// [O1o1o0g22o6o1o0] 打ちあげる死に石の連を取得
			k.Board.SetStoneAt(pointB, stoneA) // いったん、石を置く
			isExists4rensToRemove, o4rensToRemove = k.GetRenToCapture(pointB)
			isChecked4rensToRemove = true
			k.Board.SetStoneAt(pointB, Space) // 石を取り除く

			if !isExists4rensToRemove {
				// `Captured` ルールと被らなければ
				return onOpponentEye()
			}

		} else if k.CanNotPutOnMyEye && stoneA.GetColor() == renC.AdjacentColor {
			// [O1o1o0g22o4o1o0]
			// かつ、連Cに隣接する連の色が、石Aの色であったなら、
			// 自分の眼に石を置こうとしたとみなす
			return onForbiddenMyEye()

		}
	}

	// 石を置く
	k.Board.SetStoneAt(pointB, stoneA)

	// [O1o1o0g22o6o1o0] 打ちあげる死に石の連を取得
	if !isChecked4rensToRemove {
		isExists4rensToRemove, o4rensToRemove = k.GetRenToCapture(pointB)
	}

	// [O1o1o0g22o7o2o0] コウの判定
	var capturedCount = 0

	// [O1o1o0g22o6o1o0] 死に石を打ちあげる
	if isExists4rensToRemove {
		for dir := 0; dir < 4; dir++ {
			var ren = o4rensToRemove[dir]
			if ren != nil {
				k.RemoveRen(ren)

				// [O1o1o0g22o7o2o0] コウの判定
				capturedCount += ren.LibertyArea
			}
		}
	}

	// [O1o1o0g22o7o2o0] コウの判定
	var ko = Point(0)
	if capturedCount == 1 {
		ko = pointB
	}

	// 棋譜に追加
	k.Record.Push(pointB,
		// [O1o1o0g22o7o2o0] コウの判定
		ko)

	return true
}

// GetRenToCapture - 現在、着手後の盤面とします。打ち上げられる石の連を返します
//
// Returns
// -------
// isExists : bool
// renToRemove : [4]*Ren
// 隣接する東、北、西、南にある石を含む連
func (k *Kernel) GetRenToCapture(placePlay Point) (bool, [4]*Ren) {
	// [O1o1o0g22o6o1o0]
	var isExists bool
	var rensToRemove [4]*Ren

	var setAdjacentPoint = func(dir int, adjacentP Point) {
		var adjacentR = k.GetLiberty(adjacentP)
		if adjacentR.LibertyArea < 1 {
			isExists = true
			rensToRemove[dir] = adjacentR
		}
	}

	// 隣接する４方向
	k.Board.ForeachNeumannNeighborhood(placePlay, setAdjacentPoint)

	return isExists, rensToRemove
}

// EOF [O1o1o0g19o0]
