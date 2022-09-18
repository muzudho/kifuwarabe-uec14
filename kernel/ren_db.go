// BOF [O1o1o0g12o__11o__10o2o0]

package kernel

import (
	"fmt"
	"strings"
)

// RenId - 連データベースの要素のId
type RenId int

// GetRenId - 連のIdを取得
func GetRenId(boardMemoryArea int, positionNumber int, minimumLocation Point) RenId {
	return RenId(positionNumber*boardMemoryArea + int(minimumLocation))
}

type RenDb struct {
	// Header - ヘッダー
	Header RenDbDocHeader

	// 要素
	Rens map[RenId]*Ren
}

// NewRenDb - 連データベースを新規作成
func NewRenDb(boardWidth int, boardHeight int) *RenDb {
	var r = new(RenDb)
	r.Header.BoardWidth = boardWidth
	r.Header.BoardHeight = boardHeight
	return r
}

// FindRen - 連を取得
func (r *RenDb) GetRen(renId RenId) (*Ren, bool) {
	var ren, isOk = r.Rens[renId]

	if isOk {
		return ren, true
	}

	return nil, false
}

// RegisterRen - 連を登録
func (r *RenDb) RegisterRen(positionNumber int, ren *Ren) {
	var renId = GetRenId(r.Header.GetBoardMemoryArea(), positionNumber, ren.minimumLocation)
	r.Rens[renId] = ren
}

// Dump - ダンプ
func (r *RenDb) Dump() string {
	var sb strings.Builder

	// 全ての要素
	for i, item := range r.Rens {
		sb.WriteString(fmt.Sprintf("[%d]%s ", i, item.Dump()))
	}

	var text = sb.String()
	if 0 < len(text) {
		text = text[:len(text)-1]
	}
	return text
}

// EOF [O1o1o0g12o__11o__10o2o0]
