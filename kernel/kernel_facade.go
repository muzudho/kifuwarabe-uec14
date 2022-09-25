// BOF [O1o1o0g12o__11o__10o5o__10o0]

package kernel

import (
	"encoding/json"
	"os"
)

// LoadRenDb - [O1o1o0g12o__11o__10o5o__10o_10o0] 連データベースの外部ファイル読取
func (k *Kernel) LoadRenDb(path string, onError func(error) bool) bool {
	// ファイル読込
	var binary, errA = os.ReadFile(path)
	if errA != nil {
		return onError(errA)
	}

	var db = new(RenDb)
	var errB = json.Unmarshal(binary, db)
	if errB != nil {
		return onError(errB)
	}

	// 外部ファイルからの入力を、内部状態へ適用
	RefreshRenDbToInternal(db)

	// 差し替え
	k.renDb = db
	return true
}

// RemoveRen - 石の連を打ち上げます
func (k *Kernel) RemoveRen(ren *Ren) {
	// 空点をセット
	var setLocation = func(i int, location Point) {
		k.Board.SetStoneAt(location, Space)
	}

	// 場所毎に
	ren.ForeachLocation(setLocation)
}

// FindAllRens - [O1o1o0g23o_2o1o0] 盤上の全ての連を見つけます
// * 見つけた連は、連データベースへ入れます
func (k *Kernel) FindAllRens() {
	// チェックボードの初期化
	k.CheckBoard.Init(k.Board.GetWidth(), k.Board.GetHeight())

	var maxPosNthFigure = k.Record.GetMaxPosNthFigure()

	var setLocation = func(location Point) {
		var ren, isFound = k.findRen(location)
		if isFound {
			k.renDb.RegisterRen(maxPosNthFigure, k.Record.posNum, ren)
		}
	}
	// 盤上の枠の内側をスキャン。筋、段の順
	k.Board.ForeachPayloadLocationOrderByYx(setLocation)
}

// EOF [O1o1o0g12o__11o__10o5o__10o0]
