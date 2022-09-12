// BOF [O1o1o0g19o0]

package kernel

import "strings"

// DoPlay - 打つ
//
// * `command` - Example: `play black A19`
func (k *Kernel) DoPlay(command string, logg *SugaredLoggerForGame) {
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
	k.Play(stone, point)
	logg.C.Info("=\n")
	logg.J.Infow("ok")
}

func (k *Kernel) Play(stone Stone, point Point) {
	k.Board.nodes[point] = stone
}

// EOF [O1o1o0g19o0]
