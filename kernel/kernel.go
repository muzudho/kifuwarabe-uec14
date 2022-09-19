// BOF [O1o1o0g11o_3o0]

package kernel

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const geta = 1 // Japanese wooden clogs. Used to convert bases and ordinals.

// Kernel - カーネル
type Kernel struct {
	// Board - 盤
	Board *Board

	// [O1o1o0g22o2o3o0]
	// CheckBoard - 呼吸点の探索時に使います
	CheckBoard *CheckBoard
	// tempRen - 呼吸点の探索時に使います
	tempRen *Ren

	// CanNotPutOnMyEye - [O1o1o0g22o4o1o0] 自分の眼に石を置くことはできません
	CanNotPutOnMyEye bool

	// Record - [O1o1o0g12o__11o_3o0] 棋譜
	Record Record

	// RenDb - [O1o1o0g12o__11o__10o3o0] 連データベース
	renDb *RenDb
}

// NewKernel - カーネルの新規作成
func NewKernel(boardWidht int, boardHeight int,
	// [O1o1o0g12o__11o_2o0] 棋譜の初期化
	maxMoves int, playFirst Stone) *Kernel {

	var k = new(Kernel)
	k.Board = NewBoard(boardWidht, boardHeight)

	// [O1o1o0g22o2o3o0]
	k.CheckBoard = NewCheckBoard()

	// [O1o1o0g12o__11o_2o0] 棋譜の初期化
	k.Record = *NewRecord(maxMoves, playFirst)

	// RenDb - [O1o1o0g12o__11o__10o3o0] 連データベース
	k.renDb = NewRenDb(k.Board.GetWidth(), k.Board.GetHeight())

	return k
}

// Execute - 実行
//
// Returns
// -------
// isHandled : bool
// 正常終了またはエラーなら真、無視したら偽
func (k *Kernel) Execute(command string, logg *Logger) bool {

	var tokens = strings.Split(command, " ")
	switch tokens[0] {

	// この下にコマンドを挟んでいく
	// -------------------------

	case "board": // [O1o1o0g13o0]
		// 人間向けの出力
		{
			// 25列まで対応
			const fileSimbols = "ABCDEFGHJKLMNOPQRSTUVWXYZ"
			// 25行まで対応
			const rankSimbols = "   1 2 3 4 5 6 7 8 910111213141516171819202122232425"

			var filesMax = int(math.Min(25, float64(k.Board.GetWidth())))
			var filesLabel = fileSimbols[:filesMax]

			var sb strings.Builder
			// 枠の上辺
			sb.WriteString(fmt.Sprintf(`= board:'''
.     %s
.    `, filesLabel))

			var setStone = func(s Stone) {
				sb.WriteString(fmt.Sprintf("%v", s))
			}
			var doNewline = func() {
				sb.WriteString("\n.    ")
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

	case "can_not_put_on_my_eye": // [O1o1o0g22o4o2o_1o0]
		// Example 1: "can_not_put_on_my_eye get"
		// Example 2: "can_not_put_on_my_eye set true"
		var method = tokens[1]
		switch method {
		case "get":
			var value = k.CanNotPutOnMyEye
			logg.C.Infof("= %t\n", value)
			logg.J.Infow("ok", "value", value)
			return true

		case "set":
			var value = tokens[2]
			switch value {
			case "true":
				k.CanNotPutOnMyEye = true
				return true
			case "false":
				k.CanNotPutOnMyEye = false
				return true
			default:
				logg.C.Infof("? unexpected method:%s value:%s\n", method, value)
				logg.J.Infow("error", "method", method, "value", value)
				return true
			}

		default:
			logg.C.Infof("? unexpected method:%s\n", method)
			logg.J.Infow("error", "method", method)
			return true
		}

	case "find_all_rens": // [O1o1o0g23o_2o2o0]
		// Example: `find_all_rens`
		k.FindAllRens()
		logg.C.Infof("=\n")
		logg.J.Infow("ok")
		return true

	case "play": // [O1o1o0g20o0]
		// Example: `play black A19`
		k.DoPlay(command, logg)
		return true

	case "record": // [O1o1o0g12o__11o_5o0]
		// Example: "record"
		var sb strings.Builder

		var setPoint = func(positionNumber int, item *RecordItem) {
			var positionNth = positionNumber + geta // 基数を序数に変換
			var coord = k.Board.GetCodeFromPoint(item.placePlay)
			// sb.WriteString(fmt.Sprintf("[%d]%s ", positionNth, coord))

			// [O1o1o0g22o7o4o0] コウを追加
			var koStr string
			if item.ko == Point(0) {
				koStr = ""
			} else {
				koStr = fmt.Sprintf("(%s)", k.Board.GetCodeFromPoint(item.ko))
			}
			sb.WriteString(fmt.Sprintf("[%d]%s%s ", positionNth, coord, koStr))
		}

		k.Record.ForeachItem(setPoint)

		var text = sb.String()
		if 0 < len(text) {
			text = text[:len(text)-1]
		}
		logg.C.Infof("= record:'%s'\n", text)
		logg.J.Infow("ok", "record", text)
		return true

	case "remove_ren": // [O1o1o0g22o5o2o0]
		// Example: `remove_ren B2`
		var coord = tokens[1]
		var point = k.Board.GetPointFromCode(coord)
		var ren, isFound = k.GetLiberty(point)
		if isFound {
			k.RemoveRen(ren)
			logg.C.Infof("=\n")
			logg.J.Infow("ok")
			return true
		}

		logg.C.Infof("? not found ren coord:%s%\n", coord)
		logg.J.Infow("error not found ren", "coord", coord)
		return false

	case "rendb_dump": // [O1o1o0g12o__11o__10o4o0]
		var text = k.renDb.Dump()
		logg.C.Infof("= dump'''%s\n'''\n", text)
		logg.J.Infow("ok", "dump", text)
		return true

	case "rendb_load": // [O1o1o0g12o__11o__10o4o0]
		// Example: `rendb_load data/ren_db1_temp.json`
		// * ファイルパスにスペースがはいっていてはいけない
		var path = tokens[1]
		var onError = func(err error) (*RenDb, bool) {
			logg.C.Infof("? error:%s\n", err)
			logg.J.Infow("error", "err", err)
			return nil, false
		}
		var renDb, isOk = LoadRenDb(path, onError)
		if isOk {
			k.renDb = renDb
			logg.C.Infof("=\n")
			logg.J.Infow("ok")
			return true
		}
		return false

	case "rendb_save": // [O1o1o0g12o__11o__10o4o0]
		// Example: `rendb_save data/ren_db1_temp.json`
		// * ファイルパスにスペースがはいっていてはいけない
		var path = tokens[1]

		var convertLocation = func(location Point) string {
			return fmt.Sprintf("%s ", k.Board.GetCodeFromPoint(location))
		}

		var onError = func(err error) bool {
			logg.C.Infof("? error:%s\n", err)
			logg.J.Infow("error", "err", err)
			return false
		}

		var isOk = k.renDb.Save(path, convertLocation, onError)
		if isOk {
			logg.C.Infof("=\n")
			logg.J.Infow("ok")
			return true
		}

		return false

	case "set_board": // [O1o1o0g15o__14o2o0]
		// Example: `set_board file data/board1.txt`
		k.DoSetBoard(command, logg)
		logg.C.Infof("=\n")
		logg.J.Infow("ok")
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
		var coord = tokens[1]
		var point = k.Board.GetPointFromCode(coord)
		var ren, isFound = k.GetLiberty(point)
		if isFound {
			logg.C.Infof("= ren color:%s area:%d libertyArea:%d adjacentColor:%s\n", ren.color, ren.GetArea(), ren.LibertyArea, ren.adjacentColor)
			logg.J.Infow("output ren", "color", ren.color, "area", ren.GetArea(), "libertyArea", ren.LibertyArea, "adjacentColor", ren.adjacentColor)
			return true
		}

		logg.C.Infof("? not found ren coord:%s%\n", coord)
		logg.J.Infow("error not found ren", "coord", coord)
		return false

	case "test_get_point_from_code": // [O1o1o0g16o1o0]
		// Example: "test_get_point_from_code A1"
		var point = k.Board.GetPointFromCode(tokens[1])
		var code = k.Board.GetCodeFromPoint(point)
		logg.C.Infof("= %d %s", point, code)
		logg.J.Infow("ok", "point", point, "code", code)
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
