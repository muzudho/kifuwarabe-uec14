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
	Rens map[string]*Ren
}

// EOF [O1o1o0g12o__11o__101o0]
