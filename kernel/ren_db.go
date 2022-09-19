// BOF [O1o1o0g12o__11o__10o2o0]

package kernel

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// RenId - 連データベースに格納される連のId
// - 外部ファイルの可読性を優先して数値型ではなく文字列
// - 昇順に並ぶように前ゼロを付ける
type RenId string

// GetRenId - 連のIdを取得
func GetRenId(boardMemoryWidth int, positionNumber int, minimumLocation Point) RenId {
	var posNth = positionNumber + geta
	var coord = getCodeZeroPaddingFromPointOnBoard(boardMemoryWidth, minimumLocation)
	return RenId(fmt.Sprintf("%d,%s", posNth, coord))
}

type RenDb struct {
	// Header - ヘッダー
	Header RenDbDocHeader

	// 要素
	Rens map[RenId]*Ren
}

// Save - 連データベースの外部ファイル書込
func (db *RenDb) Save(path string, onError func(error) bool) bool {

	// 外部ファイルに出力するための、内部状態の整形
	db.RefreshToExternalFile()

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

// LoadRenDb - 連データベースの外部ファイル読取
func LoadRenDb(path string, onError func(error) (*RenDb, bool)) (*RenDb, bool) {
	// ファイル読込
	var binary, errA = os.ReadFile(path)
	if errA != nil {
		return onError(errA)
	}

	var renDb = new(RenDb)
	var errB = json.Unmarshal(binary, renDb)
	if errB != nil {
		return onError(errB)
	}

	return renDb, true
}

// NewRenDb - 連データベースを新規作成
func NewRenDb(boardWidth int, boardHeight int) *RenDb {
	var r = new(RenDb)
	r.Header.BoardWidth = boardWidth
	r.Header.BoardHeight = boardHeight
	r.Rens = make(map[RenId]*Ren)
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
// * すでに Id が登録されているなら、上書きしない
func (db *RenDb) RegisterRen(positionNumber int, ren *Ren) {
	var renId = GetRenId(db.Header.GetBoardMemoryWidth(), positionNumber, ren.minimumLocation)

	var _, isExists = db.Rens[renId]
	if !isExists {
		db.Rens[renId] = ren
	}
}

// Dump - ダンプ
func (db *RenDb) Dump() string {
	var sb strings.Builder

	// 全ての要素
	for idStr, item := range db.Rens {
		sb.WriteString(fmt.Sprintf("[%s]%s ", idStr, item.Dump()))
	}

	var text = sb.String()
	if 0 < len(text) {
		text = text[:len(text)-1]
	}
	return text
}

// RefreshToExternalFile - 外部ファイルに出力されてもいいように内部状態を整形します
func (db *RenDb) RefreshToExternalFile() {
	for _, ren := range db.Rens {
		ren.RefreshToExternalFile()
	}
}

// RenDbDocHeader - ヘッダー
type RenDbDocHeader struct {
	// BoardWidth - 盤の横幅
	BoardWidth int
	// BoardHeight - 盤の縦幅
	BoardHeight int
}

// GetBoardMemoryArea - 枠付き盤の面積
func (h *RenDbDocHeader) GetBoardMemoryArea() int {
	return h.GetBoardMemoryWidth() * h.GetBoardMemoryHeight()
}

// GetBoardMemoryWidth - 枠付き盤の横幅
func (h *RenDbDocHeader) GetBoardMemoryWidth() int {
	return h.BoardWidth + bothSidesWallThickness
}

// GetBoardMemoryHeight - 枠付き盤の縦幅
func (h *RenDbDocHeader) GetBoardMemoryHeight() int {
	return h.BoardHeight + bothSidesWallThickness
}

// EOF [O1o1o0g12o__11o__10o2o0]
