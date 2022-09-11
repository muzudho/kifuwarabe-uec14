// BOF [O1o1o0g16o0]

package kernel

// GetPointFromXy - 座標変換 (x,y) → Point
func (b *Board) GetPointFromXy(x int, y int) Point {
	// 枠の厚み 1 を考慮
	return Point((y+1)*b.memoryWidth + x + 1)
}

// GetPointFromCode - "A7" や "J13" といった符号を Point へ変換します
//
// * `code` - 座標の符号。 Example: "A7" や "J13"
func (b *Board) GetPointFromCode(code string) Point {
	return b.GetPointFromXy(
		GetXFromFile(GetFileFromCode(code)),
		GetYFromRank(GetRankFromCode(code)))
}

// EOF [O1o1o0g16o0]
