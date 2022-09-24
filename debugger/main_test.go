// BOF [O1o1o0g11o__11o3o0]

package debugger

import (
	"testing"
)

func TestMain(t *testing.T) {
	virtualIo.ReplaceInputToFileLines("./test.input.txt")
	main()
}

// EOF [O1o1o0g11o__11o3o0]
