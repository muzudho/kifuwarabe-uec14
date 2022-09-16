// BOF [O1o1o0g11o_4o1o0]

package kernel

import "fmt"

type Color uint

const (
	Color_None Color = iota
	Color_Black
	Color_White
	Color_Mixed
)

// String - 文字列化
func (s Color) String() string {
	switch s {
	case Color_None:
		return ""
	case Color_Black:
		return "x"
	case Color_White:
		return "o"
	case Color_Mixed:
		return "xo"
	default:
		panic(fmt.Sprintf("%d", int(s)))
	}
}

// EOF [O1o1o0g11o_4o1o0]
