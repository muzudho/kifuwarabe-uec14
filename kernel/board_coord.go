// BOF [O16o0]

package kernel

import "fmt"

// GetPointFromGtpMove - "A7" や "J13" といった符号を Point へ変換します
//
// * `gtp_move` - 座標の符号。 Example: "A7" や "J13"
func (bc *BoardCoordinate) GetPointFromGtpMove(gtp_move string) Point {
	return bc.GetPointFromXy(
		GetXFromFile(GetFileFromCode(gtp_move))+oneSideWallThickness,
		GetYFromRank(GetRankFromCode(gtp_move))+oneSideWallThickness)
}

// GetGtpMoveFromPoint - `GetPointFromGtpMove` の逆関数
func (bc *BoardCoordinate) GetGtpMoveFromPoint(point Point) string {
	return getCodeFromPointOnBoard(bc.memoryWidth, point)
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

func getFileRankFromPointOnBoard(memoryWidth int, point Point) (string, int) {
	var x, y = getXyFromPointOnBoard(memoryWidth, point)
	var file = GetFileFromX(x - oneSideWallThickness)
	var rank = getRankFromY(y) - oneSideWallThickness
	return file, rank
}

// EOF [O16o0]
