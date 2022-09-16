// BOF [O1o1o0g19o0]

package kernel

import "strings"

// DoPlay - 打つ
//
// * `command` - Example: `play black A19`
func (k *Kernel) DoPlay(command string, logg *Logger) {
	var tokens = strings.Split(command, " ")

	var isErr = false
	var getDefaultStone = func() Stone {
		isErr = true
		return Empty
	}

	var stone = GetStoneFromString(tokens[1], logg, getDefaultStone)
	if isErr {
		return
	}

	var point = k.Board.GetPointFromCode(tokens[2])

	// [O1o1o0g22o1o2o0]
	var onMasonry = func() bool {
		logg.C.Infof("? masonry my_stone:%s point:%d\n", stone, point)
		logg.J.Infow("error masonry", "my_stone", stone, "point", point)
		return false
	}

	// [O1o1o0g22o3o1o0]
	var onOpponentEye = func() bool {
		logg.C.Infof("? opponent_eye my_stone:%s point:%d\n", stone, point)
		logg.J.Infow("error opponent_eye", "my_stone", stone, "point", point)
		return false
	}

	var isOk = k.Play(stone, point, logg,
		// [O1o1o0g22o1o2o0] ,onMasonry
		onMasonry,
		// [O1o1o0g22o3o1o0] ,onOpponentEye
		onOpponentEye)

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
	onOpponentEye func() bool) bool {

	// [O1o1o0g22o1o2o0]
	if k.IsMasonryError(stoneA, pointB) {
		return onMasonry()
	}

	// [O1o1o0g22o3o1o0]
	var renC = k.GetLiberty(pointB)
	if renC.Area == 1 && stoneA.GetColor() == renC.AdjacentColor.GetOpponent() {
		// 石Aを置いた交点を含む連Cについて、
		// 連Cの面積が1であり、かつ、
		// 連Cに隣接する連の色が、石Aのちょうど反対側の色であったなら、
		// 相手の眼に石を置こうとしたとみなし、この手をエラーとする
		return onOpponentEye()
	}

	k.Board.cells[pointB] = stoneA
	return true
}

// EOF [O1o1o0g19o0]
