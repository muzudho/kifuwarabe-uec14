// BOF [O11o__11o2o_2o0]

package debugger

import "fmt"

// 文字列出力
func (vio *VirtualIO) Printf(format string, a ...interface{}) {
	fmt.Fprintf(vio.writer, format, a...)
}

// EOF [O11o__11o2o_2o0]
