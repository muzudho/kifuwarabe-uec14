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
	// locations - 要素の石の位置
	locations []Point
}

// GetArea - 面積。アゲハマの数
func (r *Ren) GetArea() int {
	return len(r.locations)
}

// AddLocation - 場所の追加
func (r *Ren) AddLocation(location Point) {
	r.locations = append(r.locations, location)
}

// ForeachLocation - 場所毎に
func (r *Ren) ForeachLocation(setLocation func(int, Point)) {
	for i, point := range r.locations {
		setLocation(i, point)
	}
}

// EOF [O1o1o0g11o_4o2o1o0]
