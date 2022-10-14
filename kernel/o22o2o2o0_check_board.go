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

	// 交点
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

// CheckStone - 石をチェックした
func (cb *CheckBoard) CheckStone(p Point) {
	cb.cells[p] |= Mark_BitStone
}

// IsChecked - 石はチェックされているか？
func (cb *CheckBoard) IsStoneChecked(p Point) bool {
	return cb.cells[p]&Mark_BitStone == Mark_BitStone
}

// CheckLiberty - 呼吸点をチェックした
func (cb *CheckBoard) CheckLiberty(p Point) {
	cb.cells[p] |= Mark_BitLiberty
}

// UncheckLiberty - 呼吸点のチェックを外した
func (cb *CheckBoard) UncheckLiberty(p Point) {
	cb.cells[p] &= ^Mark_BitLiberty // ^ はビット反転
}

// IsLibertyChecked - 呼吸点はチェックされているか？
func (cb *CheckBoard) IsLibertyChecked(p Point) bool {
	return cb.cells[p]&Mark_BitLiberty == Mark_BitLiberty
}

// EOF [O22o2o2o0]
