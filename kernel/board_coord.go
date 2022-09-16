// BOF [O1o1o0g16o0]

package kernel

import "fmt"

// GetPointFromCode - "A7" や "J13" といった符号を Point へ変換します
//
// * `code` - 座標の符号。 Example: "A7" や "J13"
func (b *Board) GetPointFromCode(code string) Point {
	// 枠の厚み
	var left_wall = 1
	var top_wall = 1
	return b.GetPointFromXy(
		GetXFromFile(GetFileFromCode(code))+left_wall,
		GetYFromRank(GetRankFromCode(code))+top_wall)
}

// GetCodeFromPoint - `GetPointFromCode` の逆関数
func (b *Board) GetCodeFromPoint(point Point) string {
	// 枠の厚み
	var left_wall = 1
	var top_wall = 1
	var x, y = b.GetXyFromPoint(point)
	var file = GetFileFromX(x - left_wall)
	var rank = GetRankFromY(y - top_wall)
	return fmt.Sprintf("%s%s", file, rank)
}

// EOF [O1o1o0g16o0]
