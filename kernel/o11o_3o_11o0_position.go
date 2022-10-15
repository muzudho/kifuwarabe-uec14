// BOF [O11o_3o_11o1o0]

package kernel

type Position struct {
	// Board - 盤
	Board *Board

	// [O22o2o3o0]
	// CheckBoard - 呼吸点の探索時に使います
	CheckBoard *CheckBoard
	// tempRen - 呼吸点の探索時に使います
	tempRen *Ren

	// CanNotPutOnMyEye - [O22o4o1o0] 自分の眼に石を置くことはできません
	CanNotPutOnMyEye bool
}

// NewDirtyKernel - カーネルの新規作成
// - 一部のメンバーは、初期化されていないので、別途初期化処理が要る
func NewDirtyPosition(gameRule GameRule, boardWidht int, boardHeight int) *Position {
	var p = new(Position)

	p.Board = NewBoard(gameRule, boardWidht, boardHeight)

	// [O22o2o3o0] チェックボード
	p.CheckBoard = NewDirtyCheckBoard()

	return p
}

// EOF [O11o_3o_11o1o0]
