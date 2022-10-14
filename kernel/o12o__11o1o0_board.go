// BOF [O12o__11o1o0]

package kernel

// Board - 盤
type Board struct {
	// ゲームルール
	gameRule GameRule
	// 盤座標
	coordinate BoardCoordinate

	// 交点
	//
	// * 英語で交点は node かも知れないが、表計算でよく使われる cell の方を使う
	cells []Stone
}

// NewBoard - 新規作成
func NewBoard(gameRule GameRule, boardWidht int, boardHeight int) *Board {
	var b = new(Board)

	// 設定ファイルから読込むので動的設定
	b.gameRule = gameRule

	// 枠の分、２つ増える
	var memoryBoardWidth = boardWidht + 2
	var memoryBoardHeight = boardHeight + 2

	b.coordinate = BoardCoordinate{
		memoryBoardWidth,
		memoryBoardHeight,
		// ４方向（東、北、西、南）の番地への相対インデックス
		[4]Point{1, Point(-memoryBoardWidth), -1, Point(memoryBoardWidth)}}

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

func getXyFromPointOnBoard(memoryWidth int, point Point) (int, int) {
	var p = int(point)
	return p % memoryWidth, p / memoryWidth
}

// サイズ変更
func (b *Board) resize(width int, height int) {
	b.coordinate.memoryWidth = width + bothSidesWallThickness
	b.coordinate.memoryHeight = height + bothSidesWallThickness
	b.cells = make([]Stone, b.coordinate.GetMemoryArea())

	// ４方向（東、北、西、南）の番地への相対インデックス
	b.coordinate.cell4Directions = [4]Point{1, Point(-b.coordinate.GetMemoryWidth()), -1, Point(b.coordinate.GetMemoryWidth())}
}

// EOF [O12o__11o1o0]
