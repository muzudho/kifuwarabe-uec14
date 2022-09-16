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

	// [O1o1o0g22o2o3o0]
	// CheckBoard - 呼吸点の探索時に使います
	CheckBoard *CheckBoard
	// Ren - 呼吸点の探索時に使います
	Ren *Ren
	// Direction - ４方向（東、北、西、南）の番地への相対インデックス
	Direction [4]int
}

func NewKernel() *Kernel {
	var k = new(Kernel)
	k.Board = NewBoard()

	// [O1o1o0g22o2o3o0]
	k.CheckBoard = NewCheckBoard()

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

	case "boardsize": // [O1o1o0g15o__11o0]
		// Example: `boardsize 19`
		var sideLength, err = strconv.Atoi(tokens[1])

		if err != nil {
			logg.C.Infof("? unexpected sideLength:%s\n", tokens[1])
			logg.J.Infow("error", "sideLength", tokens[1])
			return true
		}

		k.Board.Init(sideLength, sideLength)
		logg.C.Info("=\n")
		logg.J.Infow("ok")

		return true

	case "play": // [O1o1o0g20o0]
		// Example: `play black A19`
		k.DoPlay(command, logg)
		return true

	case "set_board": // [O1o1o0g15o__14o2o0]
		// Example: `set_board file data/board.txt`
		k.DoSetBoard(command, logg)
		return true

	case "test_coord": // [O1o1o0g12o__10o2o0]
		// Example: "test_coord A13"
		var point = k.Board.GetPointFromCode(tokens[1])
		logg.C.Infof("= %d\n", point)
		logg.J.Infow("output", "point", point)
		return true

	case "test_file": // [O1o1o0g12o__10o2o0]
		// Example: "test_file A"
		var file = GetFileFromCode(tokens[1])
		logg.C.Infof("= %s\n", file)
		logg.J.Infow("output", "file", file)
		return true

	case "test_get_liberty": // [O1o1o0g22o2o5o0]
		// Example: "test_get_liberty B2"
		var point = k.Board.GetPointFromCode(tokens[1])
		var ren = k.GetLiberty(point)
		logg.C.Infof("= ren color:%s area:%d libertyArea:%d adjacentColor:%s\n", ren.Color, ren.Area, ren.LibertyArea, ren.AdjacentColor)
		logg.J.Infow("output ren", "color", ren.Color, "area", ren.Area, "libertyArea", ren.LibertyArea, "adjacentColor", ren.AdjacentColor)
		return true

	case "test_get_point_from_xy": // [O1o1o0g12o__11o2o0]
		// Example: "test_get_point_from_xy 2 3"
		var x, errX = strconv.Atoi(tokens[1])
		if errX != nil {
			logg.C.Infof("? unexpected x:%s\n", tokens[1])
			logg.J.Infow("error", "x", tokens[1], "err", errX)
			return true
		}
		var y, errY = strconv.Atoi(tokens[2])
		if errY != nil {
			logg.C.Infof("? unexpected y:%s\n", tokens[2])
			logg.J.Infow("error", "y", tokens[2], "err", errY)
			return true
		}

		var point = k.Board.GetPointFromXy(x, y)
		logg.C.Infof("= %d\n", point)
		logg.J.Infow("output", "point", point)
		return true

	case "test_rank": // [O1o1o0g12o__10o2o0]
		// Example: "test_rank 13"
		var rank = GetRankFromCode(tokens[1])
		logg.C.Infof("= %s\n", rank)
		logg.J.Infow("output", "rank", rank)
		return true

	case "test_x": // [O1o1o0g12o__10o2o0]
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

	case "test_y": // [O1o1o0g12o__10o2o0]
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
