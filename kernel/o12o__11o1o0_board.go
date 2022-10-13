// BOF [O12o__11o1o0]

package kernel

// 片方の枠の厚み。東、北、西、南
const oneSideWallThickness = 1

// 両側の枠の厚み。南北、または東西
const bothSidesWallThickness = 2

// Board - 盤
type Board struct {
	// ゲームルール
	gameRule GameRule
	// 盤座標
	coordinate BoardCoordinate

	// 枠付きの横幅
	memoryWidth int
	// 枠付きの縦幅
	memoryHeight int

	// 交点
	//
	// * 英語で交点は node かも知れないが、表計算でよく使われる cell の方を使う
	cells []Stone

	// Direction - ４方向（東、北、西、南）の番地への相対インデックス
	Direction [4]int
}

// NewBoard - 新規作成
func NewBoard(gameRule GameRule, boardWidht int, boardHeight int) *Board {
	var b = new(Board)

	// 設定ファイルから読込むので動的設定
	b.gameRule = gameRule

	// 枠の分、２つ増える
	var memoryBoardWidth = boardWidht + 2
	var memoryBoardHeight = boardHeight + 2

	b.coordinate = BoardCoordinate{memoryBoardWidth, memoryBoardHeight}

	// 盤のサイズ指定と、盤面の初期化
	b.resize(boardWidht, boardHeight)

	return b
}

// GetGameRule - ゲームルール取得
func (b *Board) GetGameRule() *GameRule {
	return &b.gameRule
}

// SetGameRule - ゲームルール設定
func (b *Board) SetGameRule(gameRule *GameRule) {
	b.gameRule = *gameRule
}

// GetCoordinate - 盤座標取得
func (b *Board) GetCoordinate() *BoardCoordinate {
	return &b.coordinate
}

// GetMemoryWidth - 枠の厚みを含んだ横幅
func (b *Board) GetMemoryWidth() int {
	return b.memoryWidth
}

// GetMemoryHeight - 枠の厚みを含んだ縦幅
func (b *Board) GetMemoryHeight() int {
	return b.memoryHeight
}

// GetWidth - 枠の厚みを含まない横幅
func (b *Board) GetWidth() int {
	return b.memoryWidth - bothSidesWallThickness
}

// GetHeight - 枠の厚みを含まない縦幅
func (b *Board) GetHeight() int {
	return b.memoryHeight - bothSidesWallThickness
}

// GetStoneAt - 指定座標の石を取得
func (b *Board) GetStoneAt(i Point) Stone {
	return b.cells[i]
}

// SetStoneAt - 指定座標の石を設定
func (b *Board) SetStoneAt(i Point, s Stone) {
	b.cells[i] = s
}

// GetColorAt - 指定座標の石の色を取得
func (b *Board) GetColorAt(i Point) Color {
	return b.cells[i].GetColor()
}

// GetPointFromXy - 座標変換 (x,y) → Point
//
// Parameters
// ----------
// x : int
//	筋番号。 Example: 19路盤なら0～18
// y : int
//	段番号。 Example: 19路盤なら0～18
//
// Returns
// -------
// point : Point
//  配列インデックス。 Example: 2,3 なら 65
func (b *Board) GetPointFromXy(x int, y int) Point {
	return Point(y*b.memoryWidth + x)
}

// GetXyFromPoint - `GetPointFromXy` の逆関数
func (b *Board) GetXyFromPoint(point Point) (int, int) {
	return getXyFromPointOnBoard(b.memoryWidth, point)
}

func getXyFromPointOnBoard(memoryWidth int, point Point) (int, int) {
	var p = int(point)
	return p % memoryWidth, p / memoryWidth
}

// サイズ変更
func (b *Board) resize(width int, height int) {
	b.memoryWidth = width + bothSidesWallThickness
	b.memoryHeight = height + bothSidesWallThickness
	b.cells = make([]Stone, b.getMemoryArea())

	// ４方向（東、北、西、南）の番地への相対インデックス
	b.Direction = [4]int{1, -b.GetMemoryWidth(), -1, b.GetMemoryWidth()}
}

// 枠付き盤の面積
func (b *Board) getMemoryArea() int {
	return b.memoryWidth * b.memoryHeight
}

// EOF [O12o__11o1o0]
