// BOF [O1o1o0g22o5o1o0]

package kernel

// RemoveRen - 石の連を打ち上げます
func (k *Kernel) RemoveRen(ren *Ren) {
	for _, point := range ren.Elements {
		k.Board.SetStoneAt(point, Space)
	}
}

// EOF [O1o1o0g22o5o1o0]
