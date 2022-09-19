// BOF [O1o1o0g16o0]

package kernel

import "fmt"

// GetPointFromCode - "A7" や "J13" といった符号を Point へ変換します
//
// * `code` - 座標の符号。 Example: "A7" や "J13"
func (b *Board) GetPointFromCode(code string) Point {
	return b.GetPointFromXy(
		GetXFromFile(GetFileFromCode(code))+oneSideWallThickness,
		GetYFromRank(GetRankFromCode(code))+oneSideWallThickness)
}

// GetCodeFromPoint - `GetPointFromCode` の逆関数
func (b *Board) GetCodeFromPoint(point Point) string {
	return getCodeFromPointOnBoard(b.memoryWidth, point)
}

// 例えば "A1" のように、行番号をゼロサプレスして返す
func getCodeFromPointOnBoard(memoryWidth int, point Point) string {
	var file, rank = getFileRankFromPointOnBoard(memoryWidth, point)
	return fmt.Sprintf("%s%d", file, rank)
}

// 例えば "A01" のように、行番号を一律２桁のゼロ埋めにします
func getCodeZeroPaddingFromPointOnBoard(memoryWidth int, point Point) string {
	var file, rank = getFileRankFromPointOnBoard(memoryWidth, point)
	return fmt.Sprintf("%s%02d", file, rank)
}

func getFileRankFromPointOnBoard(memoryWidth int, point Point) (string, int) {
	var x, y = getXyFromPointOnBoard(memoryWidth, point)
	var file = GetFileFromX(x - oneSideWallThickness)
	var rank = getRankFromY(y) - oneSideWallThickness
	return file, rank
}

// EOF [O1o1o0g16o0]
