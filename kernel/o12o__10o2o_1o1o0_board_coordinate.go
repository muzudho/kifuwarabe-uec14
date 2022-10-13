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

// BoardCoordinate - 盤座標
type BoardCoordinate struct {
	// 枠付きの盤の水平一辺の交点の要素数
	memoryWidth int
	// 枠付きの盤の垂直一辺の交点の要素数
	memoryHeight int

	// ４方向（東、北、西、南）への相対番地。2015年講習会サンプル、GoGo とは順序が違います
	cell4Directions [4]Point
}

// GetMemoryBoardWidth - 枠付きの盤の水平一辺の交点数
func (bc *BoardCoordinate) GetMemoryBoardWidth() int {
	return bc.memoryWidth
}

// GetMemoryBoardWidth - 枠付きの盤の垂直一辺の交点数
func (bc *BoardCoordinate) GetMemoryBoardHeight() int {
	return bc.memoryHeight
}

// GetMemoryBoardArea - 壁付き盤の面積
func (bc *BoardCoordinate) GetMemoryBoardArea() int {
	return bc.GetMemoryBoardWidth() * bc.GetMemoryBoardHeight()
}

func (bc *BoardCoordinate) GetBoardWidth() int {
	// 枠の分、２つ減らす
	return bc.memoryWidth - bothSidesWallThickness
}

func (bc *BoardCoordinate) GetBoardHeight() int {
	// 枠の分、２つ減らす
	return bc.memoryHeight - bothSidesWallThickness
}

// GetBoardArea - 壁無し盤の面積
func (bc *BoardCoordinate) GetBoardArea() int {
	return bc.GetBoardWidth() * bc.GetBoardWidth()
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

// GetPointFromXy - x,y 形式の座標を、 point （配列のインデックス）へ変換します。
// point は壁を含む盤上での座標です
//
// Parameters
// ----------
// x : int
//
//	壁を含まない盤での筋番号。 Example: 19路盤なら0～18
//
// y : int
//
//	壁を含まない盤での段番号。 Example: 19路盤なら0～18
//
// Returns
// -------
// point : Point
//
//	配列インデックス。 Example: 2,3 なら 65
func (bc *BoardCoordinate) GetPointFromXy(x int, y int) Point {
	return Point(y*bc.memoryWidth + x)
}

// GetXyFromPoint - `GetPointFromXy` の逆関数
func (bc *BoardCoordinate) GetXyFromPoint(point Point) (int, int) {
	return getXyFromPointOnBoard(bc.memoryWidth, point)
}

// EOF [O12o__10o2o_1o1o0]
