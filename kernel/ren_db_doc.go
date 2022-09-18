// BOF [O1o1o0g12o__11o__101o0]

package kernel

import (
	"encoding/json"
	"os"
)

// LoadRenDbDoc - 連データベースの外部ファイル読取
func LoadRenDbDoc(path string, onError func(error) *RenDbDoc) *RenDbDoc {
	// ファイル読込
	var fileData, err = os.ReadFile(path)
	if err != nil {
		return onError(err)
	}

	var renDbDoc = new(RenDbDoc)
	json.Unmarshal(fileData, renDbDoc)

	return renDbDoc
}

// RenDbDoc - 連データベースの外部ファイル
type RenDbDoc struct {
	// Header - ヘッダー
	Header RenDbDocHeader
	// Rens - 連のハッシュテーブル
	Rens map[string]*RenDbDocRen
}

// RenDbDocHeader - ヘッダー
type RenDbDocHeader struct {
	// BoardWidth - 盤の横幅
	BoardWidth int
	// BoardHeight - 盤の縦幅
	BoardHeight int
}

// GetBoardMemoryArea - 盤の面積
func (h *RenDbDocHeader) GetBoardMemoryArea() int {
	var wallWidth = 2
	return (h.BoardWidth + wallWidth) * (h.BoardHeight + wallWidth)
}

// RenDbDocRen - 連の要素
type RenDbDocRen struct {
	// PosNth - 何手目。序数
	PosNth int
	// Loc - 座標符号の空白区切りリスト
	Loc string
}

// EOF [O1o1o0g12o__11o__101o0]
