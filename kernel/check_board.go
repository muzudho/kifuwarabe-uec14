// BOF [O22o2o2o0]

package kernel

// CheckBoard - チェック盤
type CheckBoard struct {
	// 枠付きの横幅
	memoryWidth int

	// 枠付きの縦幅
	memoryHeight int

	// 交点
	//
	// * 英語で交点は node かも知れないが、表計算でよく使われる cell の方を使う
	cells []uint8
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
	b.cells = make([]uint8, b.getMemoryArea())
}

// CheckStone - 石をチェックした
func (b *CheckBoard) CheckStone(p Point) {
	b.cells[p] |= 0b00000001
}

// IsChecked - 石はチェックされているか？
func (b *CheckBoard) IsStoneChecked(p Point) bool {
	return b.cells[p]&0b00000001 == 0b00000001
}

// CheckLiberty - 呼吸点をチェックした
func (b *CheckBoard) CheckLiberty(p Point) {
	b.cells[p] |= 0b00000010
}

// UncheckLiberty - 呼吸点のチェックを外した
func (b *CheckBoard) UncheckLiberty(p Point) {
	b.cells[p] &= 0b11111101
}

// IsLibertyChecked - 呼吸点はチェックされているか？
func (b *CheckBoard) IsLibertyChecked(p Point) bool {
	return b.cells[p]&0b00000010 == 0b00000010
}

// 枠付き盤の面積
func (b *CheckBoard) getMemoryArea() int {
	return b.memoryWidth * b.memoryHeight
}

// EOF [O22o2o2o0]
