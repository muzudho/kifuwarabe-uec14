// BOF [O12o__11o__10o5o__10o0]

package kernel

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// LoadRenDb - [O12o__11o__10o5o__10o_10o0] 連データベースの外部ファイル読取
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
	for _, ren := range db.Rens {
		var isOk = k.RefreshRenToInternal(ren)
		if !isOk {
			return false
		}
	}

	// 差し替え
	k.renDb = db
	return true
}

// RefreshRenToInternal - TODO 外部ファイルから入力された内容を内部状態に適用します
func (k *Kernel) RefreshRenToInternal(r *Ren) bool {
	{
		var getDefaultStone = func() (bool, Stone) {
			panic(fmt.Sprintf("unexpected stone:%s", r.Sto))
		}

		// TODO stone from r.Sto
		// Example: "x" --> black
		var isOk, stone = GetStoneFromChar(r.Sto, getDefaultStone)
		if !isOk {
			return false
		}
		r.stone = stone
	}
	{
		// TODO locations from r.Loc
		// Example: "C1 D1 E1"
		if 0 < len(r.Loc) {
			var codes = strings.Split(r.Loc, " ")

			var numbers = []Point{}
			for _, code := range codes {
				var location = k.Position.Board.coordinate.GetPointFromGtpMove(code)
				numbers = append(numbers, location)
			}

			r.locations = numbers
		}
	}
	{
		// TODO libertyLocations from r.LibLoc
		// Example: "F1 E2 D2 B1 C2"
		if 0 < len(r.LibLoc) {
			var codes = strings.Split(r.LibLoc, " ")

			var numbers = []Point{}
			for _, code := range codes {
				var location = k.Position.Board.coordinate.GetPointFromGtpMove(code)
				numbers = append(numbers, location)
			}

			r.libertyLocations = numbers
		}
	}

	return true
}

// RemoveRen - 石の連を打ち上げます
func (k *Kernel) RemoveRen(ren *Ren) {
	// 空点をセット
	var setLocation = func(i int, location Point) {
		k.Position.Board.SetStoneAt(location, Stone_Space)
	}

	// 場所毎に
	ren.ForeachLocation(setLocation)
}

// FindAllRens - [O23o_2o1o0] 盤上の全ての連を見つけます
// * 見つけた連は、連データベースへ入れます
func (k *Kernel) FindAllRens() {
	// チェックボードの初期化
	k.Position.CheckBoard.Init(k.Position.Board.coordinate)

	var maxPosNthFigure = k.Record.GetMaxPosNthFigure()

	var setLocation = func(location Point) {

		var libertySearchAlgorithm = NewLibertySearchAlgorithm(k.Position.Board, k.Position.CheckBoard)
		var ren, isFound = libertySearchAlgorithm.findRen(location)

		if isFound {
			k.renDb.RegisterRen(maxPosNthFigure, k.Record.positionNumber, ren)
		}
	}
	// 盤上の枠の内側をスキャン。筋、段の順
	k.Position.Board.GetCoordinate().ForeachPayloadLocationOrderByYx(setLocation)
}

// EOF [O12o__11o__10o5o__10o0]
