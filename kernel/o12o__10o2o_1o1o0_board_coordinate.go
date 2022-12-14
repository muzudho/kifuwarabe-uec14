// BOF [O12o__10o2o_1o1o0]

package kernel

import (
	"fmt"
	"strconv"
)

// 片方の枠の厚み。東、北、西、南
const oneSideWallThickness = 1

// 両側の枠の厚み。南北、または東西
const bothSidesWallThickness = 2

// Cell_4Directions - 東、北、西、南を指す配列のインデックスに対応
type Cell_4Directions int

// 東、北、西、南を指す配列のインデックスに対応
const (
	Cell_East Cell_4Directions = iota
	Cell_North
	Cell_West
	Cell_South
)

// BoardCoordinate - 盤座標
type BoardCoordinate struct {
	// 枠付きの盤の水平一辺の交点の要素数
	memoryWidth int
	// 枠付きの盤の垂直一辺の交点の要素数
	memoryHeight int

	// ４方向（東、北、西、南）への相対番地。2015年講習会サンプル、GoGo とは順序が違います
	cell4Directions [4]Point
}

// GetMemoryWidth - 枠付きの盤の水平一辺の交点数
func (bc *BoardCoordinate) GetMemoryWidth() int {
	return bc.memoryWidth
}

// GetMemoryWidth - 枠付きの盤の垂直一辺の交点数
func (bc *BoardCoordinate) GetMemoryHeight() int {
	return bc.memoryHeight
}

// GetMemoryArea - 枠付き盤の面積
func (bc *BoardCoordinate) GetMemoryArea() int {
	return bc.GetMemoryWidth() * bc.GetMemoryHeight()
}

// GetWidth - 枠無し盤の横幅
func (bc *BoardCoordinate) GetWidth() int {
	// 枠の分、２つ減らす
	return bc.memoryWidth - bothSidesWallThickness
}

// GetHeight - 枠無し盤の縦幅
func (bc *BoardCoordinate) GetHeight() int {
	// 枠の分、２つ減らす
	return bc.memoryHeight - bothSidesWallThickness
}

// GetBoardArea - 枠無し盤の面積
func (bc *BoardCoordinate) GetBoardArea() int {
	return bc.GetWidth() * bc.GetHeight()
}

// GetRelativePointOf - ４方向（東、北、西、南）の先の番地
func (bc *BoardCoordinate) GetRelativePointOf(dir4 Cell_4Directions) Point {
	return bc.cell4Directions[dir4]
}

// GetEastOf - 東
func (bc *BoardCoordinate) GetEastOf(point Point) Point {
	return point + bc.cell4Directions[Cell_East]
}

// GetNorthEastOf - 北東
func (bc *BoardCoordinate) GetNorthEastOf(point Point) Point {
	return point + bc.cell4Directions[Cell_North] + bc.cell4Directions[Cell_East]
}

// GetNorthOf - 北
func (bc *BoardCoordinate) GetNorthOf(point Point) Point {
	return point + bc.cell4Directions[Cell_North]
}

// GetNorthWestOf - 北西
func (bc *BoardCoordinate) GetNorthWestOf(point Point) Point {
	return point + bc.cell4Directions[Cell_North] + bc.cell4Directions[Cell_West]
}

// GetWestOf - 西
func (bc *BoardCoordinate) GetWestOf(point Point) Point {
	return point + bc.cell4Directions[Cell_West]
}

// GetSouthWestOf - 南西
func (bc *BoardCoordinate) GetSouthWestOf(point Point) Point {
	return point + bc.cell4Directions[Cell_South] + bc.cell4Directions[Cell_West]
}

// GetSouthOf - 南
func (bc *BoardCoordinate) GetSouthOf(point Point) Point {
	return point + bc.cell4Directions[Cell_South]
}

// GetSouthEastOf - 南東
func (bc *BoardCoordinate) GetSouthEastOf(point Point) Point {
	return point + bc.cell4Directions[Cell_South] + bc.cell4Directions[Cell_East]
}

// GetPointFromXy - x,y 形式の座標を、 point （配列のインデックス）へ変換します。
// point は枠を含む盤上での座標です
//
// Parameters
// ----------
// x : int
// 枠を含む盤での筋番号。 Example: 19路盤なら0～20
// y : int
// 枠を含む盤での段番号。 Example: 19路盤なら0～20
//
// Returns
// -------
// point : Point
// 配列インデックス
func (bc *BoardCoordinate) GetPointFromXy(x int, y int) Point {
	return Point(y*bc.memoryWidth + x)
}

// GetXyFromPoint - `GetPointFromXy` の逆関数
func (bc *BoardCoordinate) GetXyFromPoint(point Point) (int, int) {
	return getXyFromPointOnBoard(bc.memoryWidth, point)
}

// GetXFromFile - `A` ～ `Z` を 0 ～ 24 へ変換します。 国際囲碁連盟のルールに倣い、筋の符号に `I` は使いません
func GetXFromFile(file string) int {
	// 筋
	var x = file[0] - 'A'
	if file[0] >= 'I' {
		x--
	}
	return int(x)
}

// GetFileFromX - GetXFromFile の逆関数
func GetFileFromX(x int) string {
	// ABCDEFGHI
	// 012345678
	if 7 < x {
		// 'I' を飛ばす
		x++
	}
	// 筋
	return fmt.Sprintf("%c", 'A'+x)
}

// GetYFromRank - '1' ～ '99' を 0 ～ 98 へ変換します
func GetYFromRank(rank string) int {
	// 段
	var y = int(rank[0] - '0')
	if 1 < len(rank) {
		y *= 10
		y += int(rank[1] - '0')
	}
	return y - 1
}

// GetRankFromY - GetYFromRank の逆関数
//
// Parameters
// ----------
// y : int
//
//	0 .. 98
//
// Returns
// -------
// rank : string
//
//	"1" .. "99"
func GetRankFromY(y int) string {
	return strconv.Itoa(getRankFromY(y))
}

func getRankFromY(y int) int {
	return y + 1
}

// GetFileFromCode - 座標の符号の筋の部分を抜き出します
//
// * `code` - 座標の符号。 Example: "A7" や "J13"
func GetFileFromCode(code string) string {
	return code[0:1]
}

// GetRankFromCode - 座標の符号の段の部分を抜き出します
//
// * `code` - 座標の符号。 Example: "A7" や "J13"
func GetRankFromCode(code string) string {
	if 2 < len(code) {
		return code[1:3]
	}

	return code[1:2]
}

// ForeachLikeText - 枠を含めた各セルの石
func (bc *BoardCoordinate) ForeachLikeText(setPoint func(Point), doNewline func()) {
	for y := 0; y < bc.memoryHeight; y++ {
		if y != 0 {
			doNewline()
		}

		for x := 0; x < bc.memoryWidth; x++ {
			var i = bc.GetPointFromXy(x, y)
			setPoint(i)
		}
	}
}

// GetPointFromGtpMove - "A7" や "J13" といった符号を Point へ変換します
//
// * `gtp_move` - 座標の符号。 Example: "A7" や "J13"
func (bc *BoardCoordinate) GetPointFromGtpMove(gtp_move string) Point {
	return bc.GetPointFromXy(
		GetXFromFile(GetFileFromCode(gtp_move))+oneSideWallThickness,
		GetYFromRank(GetRankFromCode(gtp_move))+oneSideWallThickness)
}

func getFileRankFromPointOnBoard(memoryWidth int, point Point) (string, int) {
	var x, y = getXyFromPointOnBoard(memoryWidth, point)
	var file = GetFileFromX(x - oneSideWallThickness)
	var rank = getRankFromY(y) - oneSideWallThickness
	return file, rank
}

// 例えば "A1" のように、行番号をゼロサプレスして返す
func getCodeFromPointOnBoard(memoryWidth int, point Point) string {
	var file, rank = getFileRankFromPointOnBoard(memoryWidth, point)
	return fmt.Sprintf("%s%d", file, rank)
}

// 例えば "01A" のように、符号を行、列の順とし、かつ、行番号を一律２桁のゼロ埋めにします
func getRenIdFromPointOnBoard(memoryWidth int, point Point) string {
	var file, rank = getFileRankFromPointOnBoard(memoryWidth, point)
	return fmt.Sprintf("%02d%s", rank, file)
}

// GetGtpMoveFromPoint - `GetPointFromGtpMove` の逆関数
func (bc *BoardCoordinate) GetGtpMoveFromPoint(point Point) string {
	return getCodeFromPointOnBoard(bc.memoryWidth, point)
}

// ForeachCellWithoutWall - 枠や改行を含めない各セルの番地
func (bc *BoardCoordinate) ForeachCellWithoutWall(setPoint func(Point)) {
	// - x,y は枠無しの盤での0から始まる座標
	// - point は枠無しの盤を配列にしたもので0から始まる配列の添え字
	var heightNth = bc.memoryHeight - 1
	var widthNth = bc.memoryWidth - 1
	for y := 1; y < heightNth; y++ {
		for x := 1; x < widthNth; x++ {
			var point = bc.GetPointFromXy(x, y)
			setPoint(point)
		}
	}
}

// ForeachPayloadLocationOrderByYx - 枠や改行を含めない各セルの番地。筋、段の順
func (bc *BoardCoordinate) ForeachPayloadLocationOrderByYx(setLocation func(Point)) {
	var height = bc.memoryHeight - 1
	var width = bc.memoryWidth - 1
	for x := 1; x < width; x++ {
		for y := 1; y < height; y++ {
			var i = bc.GetPointFromXy(x, y)
			setLocation(i)
		}
	}
}

// EOF [O12o__10o2o_1o1o0]
