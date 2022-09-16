// BOF [O1o1o0g15o__14o1o0]

package kernel

import (
	"os"
	"strings"
)

// DoSetBoard - 盤面を設定する
//
// コマンドラインの複数行入力は難しいので、ファイルから取ることにする
// * `command` - Example: `set_board file data/board.txt`
// ........................--------- ---- --------------
// ........................0         1    2
func (k *Kernel) DoSetBoard(command string, logg *Logger) {
	var tokens = strings.Split(command, " ")

	if tokens[1] == "file" {
		var filePath = tokens[2]

		var fileData, err = os.ReadFile(filePath)
		if err != nil {
			logg.C.Infof("? unexpected file:%s\n", filePath)
			logg.J.Infow("error", "file", filePath)
			return
		}

		var getDefaultStone = func() (bool, Stone) {
			return false, Space
		}

		var size = k.Board.getMemoryArea()
		var i Point = 0
		for _, c := range string(fileData) {
			var str = string([]rune{c})
			var isOk, stone = GetStoneFromChar(str, getDefaultStone)

			if isOk {
				if size <= int(i) {
					// 配列サイズ超過
					logg.C.Infof("? board out of bounds i:%d size:%d\n", i, size)
					logg.J.Infow("error board out of bounds", "i", i, "size", size)
					return
				}

				k.Board.SetStoneAt(i, stone)
				i++
			}
		}

		// サイズが足りていないなら、エラー
		if int(i) != size {
			logg.C.Infof("? not enough size i:%d size:%d\n", i, size)
			logg.J.Infow("error not enough size", "i", i, "size", size)
			return
		}
	}

}

// EOF [O1o1o0g15o__14o1o0]
