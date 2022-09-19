// BOF [O1o1o0g11o_4o2o1o0]

package kernel

import (
	"fmt"
	"math"
	"strings"
)

// Ren - 連，れん
type Ren struct {
	// Loc - （外部ファイル向け）石の盤上の座標符号の空白区切りのリスト
	Loc string

	// LibertyArea - 呼吸点の面積
	LibertyArea int

	// 色
	color Color
	// 隣接する石の色
	adjacentColor Color
	// 要素の石の位置
	locations []Point
	// 最小の場所。Idとして利用することを想定
	minimumLocation Point
}

// NewRen - 連を新規作成
//
// Parameters
// ----------
// color - 色
func NewRen(color Color) *Ren {
	var r = new(Ren)
	r.color = color
	r.adjacentColor = Color_None
	r.minimumLocation = math.MaxInt
	return r
}

// GetArea - 面積。アゲハマの数
func (r *Ren) GetArea() int {
	return len(r.locations)
}

// GetColor - 色
func (r *Ren) GetColor() Color {
	return r.color
}

// GetAdjacentColor - 隣接する石の色
func (r *Ren) GetAdjacentColor() Color {
	return r.adjacentColor
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
//
// Example: `A1 B2 C3 D4`
func (r *Ren) Dump() string {
	return r.createCoordBelt()
}

// Example: `A1 B2 C3 D4`
func (r *Ren) createCoordBelt() string {
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

// RefreshToExternalFile - 外部ファイルに出力されてもいいように内部状態を整形します
func (r *Ren) RefreshToExternalFile() {
	r.Loc = r.createCoordBelt()
}

// EOF [O1o1o0g11o_4o2o1o0]
