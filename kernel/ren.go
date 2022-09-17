// BOF [O1o1o0g11o_4o2o1o0]

package kernel

// Ren - 連，れん
type Ren struct {
	// Color - 色
	Color Color
	// AdjacentColor - 隣接する石の色
	AdjacentColor Color
	// LibertyArea - 呼吸点の面積
	LibertyArea int
	// Elements - 要素の石の位置
	Elements []Point
}

// GetArea - 面積。アゲハマの数
func (r *Ren) GetArea() int {
	return len(r.Elements)
}

// EOF [O1o1o0g11o_4o2o1o0]
