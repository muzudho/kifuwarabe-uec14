package debugger

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// VirtualIO - 入出力を模擬したもの
type VirtualIO struct {
	scanner *bufio.Scanner
	writer  *bufio.Writer
}

// 新規作成
func newVirtualIO() *VirtualIO {
	// 実体をメモリ上に占有させ、そのアドレスを返す
	return &VirtualIO{
		scanner: bufio.NewScanner(os.Stdin),
		writer:  bufio.NewWriter(os.Stdout),
	}
}

// 次の文字列入力を読取る
func (io *VirtualIO) nextLine() string {
	io.scanner.Scan()
	return io.scanner.Text()
}

// 次の整数入力を読取る
func (io *VirtualIO) nextInt() int {
	i, e := strconv.Atoi(io.nextLine())
	if e != nil {
		panic(e)
	}
	return i
}

// 文字列出力
func (io *VirtualIO) printLn(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}

// 新規作成
var virtualIo = newVirtualIO()

func main() {
	virtualIo.scanner.Split(bufio.ScanWords)      // switch to separating by space
	virtualIo.scanner.Buffer([]byte{}, 100000007) // switch to read large size input
	defer virtualIo.writer.Flush()

	N := virtualIo.nextInt()
	hoge := fmt.Sprintf("%d is ok", N) // なんらかの処理
	virtualIo.printLn(hoge)            // 出力
}
