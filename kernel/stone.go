// BOF [O1o1o0g11o0]

package kernel

import "fmt"

type Stone uint

const (
	Empty Stone = iota
	Black
	White
	Wall
)

// GetStoneFromString - 文字列を元に値を返します
func GetStoneFromString(stoneName string, logg *Logger, getDefaultStone func() Stone) Stone {
	switch stoneName {
	case "empty":
		return Empty
	case "black":
		return Black
	case "white":
		return White
	case "wall":
		return Wall
	default:
		logg.C.Infof("? unexpected stone:%s\n", stoneName)
		logg.J.Infow("error", "stone", stoneName)
		return getDefaultStone()
	}
}

// String - 文字列化
func (s Stone) String() string {
	switch s {
	case Empty:
		return "."
	case Black:
		return "x"
	case White:
		return "o"
	case Wall:
		return "+"
	default:
		panic(fmt.Sprintf("%d", int(s)))
	}
}

// GetColor - 色の取得
func (s Stone) GetColor() Color {
	switch s {
	case Empty:
		return Color_None
	case Black:
		return Color_Black
	case White:
		return Color_White
	case Wall:
		return Color_None
	default:
		panic(fmt.Sprintf("%d", int(s)))
	}
}

// EOF [O1o1o0g11o0]
