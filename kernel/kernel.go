// BOF [O1o1o0g11o_3o0]

package kernel

import (
	"fmt"
	"strconv"
	"strings"
)

type Kernel struct {
	// Board - 盤
	Board *Board
}

func NewKernel() *Kernel {
	var k = new(Kernel)
	k.Board = NewBoard()
	return k
}

// Execute - 実行
func (k *Kernel) Execute(command string, logg *Logger) bool {

	var tokens = strings.Split(command, " ")
	switch tokens[0] {

	// この下にコマンドを挟んでいく
	// -------------------------

	case "board": // [O1o1o0g13o0]
		// 人間向けの出力
		{
			var sb strings.Builder
			sb.WriteString(`= board:'''
. `)

			var setStone = func(s Stone) {
				sb.WriteString(fmt.Sprintf("%v", s))
			}
			var doNewline = func() {
				sb.WriteString("\n. ")
			}
			k.Board.ForeachLikeText(setStone, doNewline)
			sb.WriteString("\n. '''\n")
			logg.C.Info(sb.String())
		}
		// コンピューター向けの出力
		{
			var sb strings.Builder

			var setStone = func(s Stone) {
				sb.WriteString(fmt.Sprintf("%v", s))
			}
			var doNewline = func() {
				// pass
			}
			k.Board.ForeachLikeText(setStone, doNewline)
			logg.J.Infow("output", "board", sb.String())
		}
		return true

	case "boardsize": // [O1o1o0g15o_11o0]
		// Example: `boardsize 19`
		var sideLength, err = strconv.Atoi(tokens[1])

		if err != nil {
			logg.C.Infof("? unexpected sideLength:%s\n", tokens[1])
			logg.J.Infow("error", "sideLength", tokens[1])
			return true
		}

		k.Board.Resize(sideLength, sideLength)
		logg.C.Info("=\n")
		logg.J.Infow("ok")

		return true

	case "play": // [O1o1o0g20o0]
		// Example: `play black A19`
		k.DoPlay(command, logg)
		return true

	case "test_coord": // [O1o1o0g17o0]
		// Example: "test_coord A13"
		var point = k.Board.GetPointFromCode(tokens[1])
		logg.C.Infof("= %d\n", point)
		logg.J.Infow("output", "point", point)
		return true

	case "test_file": // [O1o1o0g17o0]
		// Example: "test_file A"
		var file = GetFileFromCode(tokens[1])
		logg.C.Infof("= %s\n", file)
		logg.J.Infow("output", "file", file)
		return true

	case "test_rank": // [O1o1o0g17o0]
		// Example: "test_rank 13"
		var rank = GetRankFromCode(tokens[1])
		logg.C.Infof("= %s\n", rank)
		logg.J.Infow("output", "rank", rank)
		return true

	case "test_x": // [O1o1o0g17o0]
		// Example: "test_x 18"
		var x, err = strconv.Atoi(tokens[1])
		if err != nil {
			logg.C.Infof("? unexpected x:%s\n", tokens[1])
			logg.J.Infow("error", "x", tokens[1])
			return true
		}
		var file = GetFileFromX(x)
		logg.C.Infof("= %s\n", file)
		logg.J.Infow("output", "file", file)
		return true

	case "test_y": // [O1o1o0g17o0]
		// Example: "test_y 18"
		var y, err = strconv.Atoi(tokens[1])
		if err != nil {
			logg.C.Infof("? unexpected y:%s\n", tokens[1])
			logg.J.Infow("error", "y", tokens[1])
			return true
		}
		var rank = GetRankFromY(y)
		logg.C.Infof("= %s\n", rank)
		logg.J.Infow("output", "rank", rank)
		return true

	// この上にコマンドを挟んでいく
	// -------------------------

	default:
	}

	return false
}

// EOF [O1o1o0g11o_3o0]
