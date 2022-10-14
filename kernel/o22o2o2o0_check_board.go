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
	// 枠付きの横幅
	memoryWidth int

	// 枠付きの縦幅
	memoryHeight int

	// 交点
	//
	// * 英語で交点は node かも知れないが、表計算でよく使われる cell の方を使う
	cells []Mark
}

// NewCheckBoard - 新規作成
//
// * 使用する前に Init 関数を呼び出してほしい
func NewCheckBoard() *CheckBoard {
	var b = new(CheckBoard)
	return b
}

// Init - 盤面初期化
func (b *CheckBoard) Init(width int, height int) {
	// 盤面のサイズが異なるなら、盤面を作り直す
	if b.memoryWidth != width || b.memoryHeight != height {
		b.Resize(width, height)
		return
	}

	// 盤面のクリアー
	for i := 0; i < len(b.cells); i++ {
		b.cells[i] = 0
	}
}

// Resize - サイズ変更
func (b *CheckBoard) Resize(width int, height int) {
	b.memoryWidth = width + 2
	b.memoryHeight = height + 2
	b.cells = make([]Mark, b.getMemoryArea())
}

// CheckStone - 石をチェックした
func (b *CheckBoard) CheckStone(p Point) {
	b.cells[p] |= Mark_BitStone
}

// IsChecked - 石はチェックされているか？
func (b *CheckBoard) IsStoneChecked(p Point) bool {
	return b.cells[p]&Mark_BitStone == Mark_BitStone
}

// CheckLiberty - 呼吸点をチェックした
func (b *CheckBoard) CheckLiberty(p Point) {
	b.cells[p] |= Mark_BitLiberty
}

// UncheckLiberty - 呼吸点のチェックを外した
func (b *CheckBoard) UncheckLiberty(p Point) {
	b.cells[p] &= ^Mark_BitLiberty // ^ はビット反転
}

// IsLibertyChecked - 呼吸点はチェックされているか？
func (b *CheckBoard) IsLibertyChecked(p Point) bool {
	return b.cells[p]&Mark_BitLiberty == Mark_BitLiberty
}

// 枠付き盤の面積
func (b *CheckBoard) getMemoryArea() int {
	return b.memoryWidth * b.memoryHeight
}

// EOF [O22o2o2o0]
