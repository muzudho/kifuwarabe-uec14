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
