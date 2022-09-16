// BOF [O1o1o0g15o__14o1o0]

package kernel

import (
	"os"
	"strings"
)

// DoSetBoard - 盤面を設定する
//
// コマンドラインの複数行入力は難しいので、ファイルから取ることにする
//
//   - `command` - Example: `set_board file input.txt`
//     --------- ---- ---------
//     0         1    2
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

		var getDefaultStone = func() Stone {
			return Empty
		}

		var i Point = 0
		for _, c := range string(fileData) {
			var str = string([]rune{c})
			var stone = GetStoneFromString(str, getDefaultStone)
			k.Board.SetStoneAt(i, stone)
			i++
		}
	}

}

// EOF [O1o1o0g15o__14o1o0]
