// BOF [O1o1o0g12o__11o__10o2o0]

package kernel

import (
	"encoding/json"
	"fmt"
	"os"
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

// Save - 連データベースの外部ファイル書込
func (db *RenDb) Save(path string, onError func(error) bool) bool {

	// Marshal関数でjsonエンコード
	// ->返り値jsonDataにはエンコード結果が[]byteの形で格納される
	jsonBinary, errA := json.Marshal(db)
	if errA != nil {
		return onError(errA)
	}

	// ファイル読込
	var errB = os.WriteFile(path, jsonBinary, 0664)
	if errB != nil {
		return onError(errB)
	}

	return true
}

// Load - 連データベースの外部ファイル読取
func Load(path string, onError func(error) *RenDb) *RenDb {
	// ファイル読込
	var binary, err = os.ReadFile(path)
	if err != nil {
		return onError(err)
	}

	var renDb = new(RenDb)
	var errB = json.Unmarshal(binary, renDb)
	if errB != nil {
		return onError(err)
	}

	return renDb
}

// NewRenDb - 連データベースを新規作成
func NewRenDb(boardWidth int, boardHeight int) *RenDb {
	var r = new(RenDb)
	r.Header.BoardWidth = boardWidth
	r.Header.BoardHeight = boardHeight
	return r
}

// FindRen - 連を取得
func (db *RenDb) GetRen(renId RenId) (*Ren, bool) {
	var ren, isOk = db.Rens[renId]

	if isOk {
		return ren, true
	}

	return nil, false
}

// RegisterRen - 連を登録
func (db *RenDb) RegisterRen(positionNumber int, ren *Ren) {
	var renId = GetRenId(db.Header.GetBoardMemoryArea(), positionNumber, ren.minimumLocation)
	db.Rens[renId] = ren
}

// Dump - ダンプ
func (db *RenDb) Dump() string {
	var sb strings.Builder

	// 全ての要素
	for i, item := range db.Rens {
		sb.WriteString(fmt.Sprintf("[%d]%s ", i, item.Dump()))
	}

	var text = sb.String()
	if 0 < len(text) {
		text = text[:len(text)-1]
	}
	return text
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

// EOF [O1o1o0g12o__11o__10o2o0]
