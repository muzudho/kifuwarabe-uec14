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

	// [O1o1o0g22o3o1o0]
	var renC = k.GetLiberty(pointB)
	if renC.GetArea() == 1 { // 石Aを置いた交点を含む連Cについて、連Cの面積が1である（眼）
		if stoneA.GetColor() == renC.AdjacentColor.GetOpponent() {
			// かつ、連Cに隣接する連の色が、石Aのちょうど反対側の色であったなら、
			// 相手の眼に石を置こうとしたとみなす
			return onOpponentEye()

		} else if k.CanNotPutOnMyEye && stoneA.GetColor() == renC.AdjacentColor {
			// [O1o1o0g22o4o1o0]
			// かつ、連Cに隣接する連の色が、石Aの色であったなら、
			// 自分の眼に石を置こうとしたとみなす
			return onForbiddenMyEye()

		}
	}

	// 石を置く
	k.Board.cells[pointB] = stoneA

	// [O1o1o0g22o6o1o0] 死に石を打ちあげる
	var renToRemove [4]*Ren
	for dir := 0; dir < 4; dir++ { // 東、北、西、南
		var adjacentP = pointB + Point(k.Direction[dir]) // 隣接する交点
		var adjacentR = k.GetLiberty(adjacentP)
		if adjacentR.LibertyArea < 1 {
			renToRemove[dir] = adjacentR
		}
	}

	for dir := 0; dir < 4; dir++ {
		if renToRemove[dir] != nil {
			k.RemoveRen(renToRemove[dir])
		}
	}

	return true
}

// EOF [O1o1o0g19o0]
