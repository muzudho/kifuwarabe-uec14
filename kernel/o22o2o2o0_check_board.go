// BOF [O22o2o2o0]

package kernel

// Mark - 目印
type Mark uint8

const (
	Mark_BitStone   Mark = 0b00000001
	Mark_BitLiberty Mark = 0b00000010
)

// CheckBoard - チェック盤
type CheckBoard struct {
	// 盤座標
	coordinate BoardCoordinate

	// 長さが可変な盤
	//
	// * 英語で交点は node かも知れないが、表計算でよく使われる cell の方を使う
	cells []Mark
}

// NewCheckBoard - 新規作成
//
// * このメソッドを呼び出した後に Init 関数を呼び出してほしい
func NewCheckBoard(coordinate BoardCoordinate) *CheckBoard {
	var cb = new(CheckBoard)

	cb.coordinate = coordinate

	return cb
}

// Init - 盤面初期化
func (cb *CheckBoard) Init(newBoardCoordinate BoardCoordinate) {
	// 盤面のサイズが異なるなら、盤面を作り直す
	if cb.coordinate.memoryWidth != newBoardCoordinate.memoryWidth || cb.coordinate.memoryHeight != newBoardCoordinate.memoryHeight {
		cb.coordinate = newBoardCoordinate
		cb.cells = make([]Mark, cb.coordinate.GetMemoryArea())
		return
	}

	// 盤面のクリアー
	for i := 0; i < len(cb.cells); i++ {
		cb.cells[i] = 0
	}
}

// Overwrite - 上書き
func (cb *CheckBoard) Overwrite(point Point, mark Mark) {
	cb.cells[point] |= mark
}

// Erase - 消す
func (cb *CheckBoard) Erase(point Point, mark Mark) {
	cb.cells[point] &= ^mark // ^ はビット反転
}

// Contains - 含む
func (cb *CheckBoard) Contains(point Point, mark Mark) bool {
	return cb.cells[point]&mark == mark
}

// EOF [O22o2o2o0]
