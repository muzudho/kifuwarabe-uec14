// BOF [O12o__11o__10o2o0]

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
func GetRenId(boardMemoryWidth int, positionNthFigure int, positionNumber int, minimumLocation Point) RenId {
	var posNth = positionNumber + geta
	var coord = getRenIdFromPointOnBoard(boardMemoryWidth, minimumLocation)

	return RenId(fmt.Sprintf("%0*d,%s", positionNthFigure, posNth, coord))
}

// RenDb - 連データベース
type RenDb struct {
	// Header - ヘッダー
	Header RenDbDocHeader `json:"header"`

	// 要素
	Rens map[RenId]*Ren `json:"rens"`
}

// NewRenDb - 連データベースを新規作成
func NewRenDb(boardWidth int, boardHeight int) *RenDb {
	var db = new(RenDb)
	db.Header.Init(boardWidth, boardHeight)
	db.Rens = make(map[RenId]*Ren)
	return db
}

// Init - 初期化
func (db *RenDb) Init(boardWidth int, boardHeight int) {
	db.Header.Init(boardWidth, boardHeight)

	// Clear
	for ri := range db.Rens {
		delete(db.Rens, ri)
	}
}

// Save - 連データベースの外部ファイル書込
func (db *RenDb) Save(path string, convertLocation func(Point) string, onError func(error) bool) bool {

	// 外部ファイルに出力するための、内部状態の整形
	db.RefreshToExternalFile(convertLocation)

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
func (db *RenDb) RegisterRen(positionNthFigure int, positionNumber int, ren *Ren) {
	var renId = GetRenId(db.Header.GetBoardMemoryWidth(), positionNthFigure, positionNumber, ren.minimumLocation)

	var _, isExists = db.Rens[renId]
	if !isExists {
		db.Rens[renId] = ren
	}
}

// Dump - ダンプ
func (db *RenDb) Dump() string {
	var sb strings.Builder

	// 全ての要素
	for idStr, ren := range db.Rens {
		sb.WriteString(fmt.Sprintf("[%s]%s \n", idStr, ren.Dump()))
	}

	var text = sb.String()
	if 0 < len(text) {
		text = text[:len(text)-1]
	}
	return text
}

// RefreshToExternalFile - 外部ファイルに出力されてもいいように内部状態を整形します
func (db *RenDb) RefreshToExternalFile(convertLocation func(Point) string) {
	for _, ren := range db.Rens {
		ren.RefreshToExternalFile(convertLocation)
	}
}

// RenDbDocHeader - ヘッダー
type RenDbDocHeader struct {
	// BoardWidth - 盤の横幅
	BoardWidth int `json:"boardWidth"`
	// BoardHeight - 盤の縦幅
	BoardHeight int `json:"boardHeight"`
}

// Init - 初期化
func (h *RenDbDocHeader) Init(boardWidth int, boardHeight int) {
	h.BoardWidth = boardWidth
	h.BoardHeight = boardHeight
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

// EOF [O12o__11o__10o2o0]
