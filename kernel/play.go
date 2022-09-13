// BOF [O1o1o0g19o0]

package kernel

import "strings"

// DoPlay - 打つ
//
// * `command` - Example: `play black A19`
func (k *Kernel) DoPlay(command string, logg *Logger) {
	var tokens = strings.Split(command, " ")

	var stone Stone
	switch tokens[1] {
	case "empty":
		stone = Empty
	case "black":
		stone = Black
	case "white":
		stone = White
	case "wall":
		stone = Wall
	default:
		logg.C.Infof("? unexpected stone:%s\n", tokens[1])
		logg.J.Infow("error", "stone", tokens[1])
		return
	}

	var point = k.Board.GetPointFromCode(tokens[2])

	// [O1o1o0g22o1o2o0]
	var onMasonryError = func() bool {
		logg.C.Infof("? masonry my_stone:%s point:%d\n", stone, point)
		logg.J.Infow("error", "my_stone", stone, "point", point)
		return false
	}

	var isOk = k.Play(stone, point, logg,
		// [O1o1o0g22o1o2o0] ,onMasonryError
		onMasonryError)
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
func (k *Kernel) Play(stone Stone, point Point, logg *Logger,
	// [O1o1o0g22o1o2o0] onMasonryError
	onMasonryError func() bool) bool {

	// [O1o1o0g22o1o2o0]
	if k.IsMasonryError(stone, point) {
		return onMasonryError()
	}

	k.Board.cells[point] = stone
	return true
}

// EOF [O1o1o0g19o0]
