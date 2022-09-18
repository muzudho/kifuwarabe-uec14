// BOF [O1o1o0g11o_4o2o1o0]

package kernel

import (
	"fmt"
	"math"
	"strings"
)

// Ren - 連，れん
type Ren struct {
	// Color - 色
	Color Color
	// AdjacentColor - 隣接する石の色
	AdjacentColor Color
	// LibertyArea - 呼吸点の面積
	LibertyArea int
	// 要素の石の位置
	locations []Point
	// 最小の場所。Idとして利用することを想定
	minimumLocation Point
}

// NewRen - 連を新規作成
func NewRen() *Ren {
	var r = new(Ren)
	r.minimumLocation = math.MaxInt
	return r
}

// GetArea - 面積。アゲハマの数
func (r *Ren) GetArea() int {
	return len(r.locations)
}

// GetMinimumLocation - 最小の場所。Idとして利用することを想定
func (r *Ren) GetMinimumLocation() Point {
	return r.minimumLocation
}

// AddLocation - 場所の追加
func (r *Ren) AddLocation(location Point) {
	r.locations = append(r.locations, location)

	// 最小の数を更新
	r.minimumLocation = Point(math.Min(float64(r.minimumLocation), float64(location)))
}

// ForeachLocation - 場所毎に
func (r *Ren) ForeachLocation(setLocation func(int, Point)) {
	for i, point := range r.locations {
		setLocation(i, point)
	}
}

// Dump - ダンプ
func (r *Ren) Dump() string {
	var sb strings.Builder

	// 全ての要素
	for _, location := range r.locations {
		sb.WriteString(fmt.Sprintf("%d ", location))
	}

	var text = sb.String()
	if 0 < len(text) {
		text = text[:len(text)-1]
	}
	return text
}

// EOF [O1o1o0g11o_4o2o1o0]
