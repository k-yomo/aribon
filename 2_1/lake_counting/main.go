package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reached [][]bool

func main() {
	io, flush := NewIO()
	defer flush()

	n, m := io.ScanInt2()
	reached = PrepareEmpty2DBoolArray(n, m)
	graph := io.Scan2DGraph(n)
	var lakeCount int

	for y := 0; y < n; y++ {
		for x := 0; x < m; x++ {
			if graph[y][x] == "." || reached[y][x] {
				continue
			}
			lakeCount++
			Search(graph, y, x)
		}
	}
	fmt.Println(lakeCount)
}

func Search(graph [][]string, y, x int) {
	if y < 0 || x < 0 || y > len(graph)-1 || x > len(graph[0])-1 || graph[y][x] == "." || reached[y][x] {
		return
	}
	reached[y][x] = true
	Search(graph, y+1, x)
	Search(graph, y+1, x+1)
	Search(graph, y, x+1)
	Search(graph, y-1, x)
	Search(graph, y-1, x+1)
	Search(graph, y-1, x-1)
	Search(graph, y, x-1)
	Search(graph, y+1, x-1)
}

func PrepareEmptyBoolArray(n int) []bool {
	arr := make([]bool, n)
	for i := 0; i < n; i++ {
		arr[i] = false
	}
	return arr
}

func PrepareEmpty2DBoolArray(y, x int) [][]bool {
	arr := make([][]bool, y)
	for i := 0; i < y; i++ {
		arr[i] = PrepareEmptyBoolArray(x)
	}
	return arr
}

type IO struct {
	scanner *bufio.Scanner
	writer  *bufio.Writer
}

func NewIO() (*IO, func()) {
	io := &IO{
		scanner: newScanner(),
		writer:  newWriter(),
	}
	return io, func() { io.writer.Flush() }
}

func newScanner() *bufio.Scanner {
	s := bufio.NewScanner(os.Stdin)
	s.Buffer(make([]byte, 10000000), 10000000)
	s.Split(bufio.ScanWords)
	return s
}

func newWriter() *bufio.Writer {
	return bufio.NewWriter(os.Stdout)
}

func (io *IO) ScanBytes() []byte {
	if !io.scanner.Scan() {
		panic("scan string failed")
	}
	return io.scanner.Bytes()
}

func (io *IO) ScanString() string {
	if !io.scanner.Scan() {
		panic("scan string failed")
	}
	return io.scanner.Text()
}

func (io *IO) ScanStrings(n int) []string {
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		strs[i] = io.ScanString()
	}
	return strs
}

func (io *IO) Scan2DStrings(y, x int) [][]string {
	strings := make([][]string, y)
	for i := 0; i < y; i++ {
		strings[i] = io.ScanStrings(x)
	}
	return strings
}

func (io *IO) Scan2DGraph(y int) [][]string {
	strs := make([][]string, y)
	for i := 0; i < y; i++ {
		strs[i] = strings.Split(io.ScanString(), "")
	}
	return strs
}

func (io *IO) ScanInt() int {
	return int(io.ScanInt64())
}

func (io *IO) ScanInt2() (int, int) {
	return int(io.ScanInt64()), int(io.ScanInt64())
}

func (io *IO) ScanInt3() (int, int, int) {
	return int(io.ScanInt64()), int(io.ScanInt64()), int(io.ScanInt64())
}

func (io *IO) ScanInt4() (int, int, int, int) {
	return int(io.ScanInt64()), int(io.ScanInt64()), int(io.ScanInt64()), int(io.ScanInt64())
}

func (io *IO) ScanInts(n int) []int {
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = io.ScanInt()
	}
	return ints
}

func (io *IO) Scan2DInts(y, x int) [][]int {
	ints := make([][]int, y)
	for i := 0; i < y; i++ {
		ints[i] = io.ScanInts(x)
	}
	return ints
}

func (io *IO) ScanInt64() int64 {
	i, err := strconv.ParseInt(io.ScanString(), 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func (io *IO) ScanFloat64() float64 {
	i, _ := strconv.ParseFloat(io.ScanString(), 64)
	return i
}

func (io *IO) ScanFloat64s(n int) []float64 {
	floats := make([]float64, n)
	for i := 0; i < n; i++ {
		floats[i] = io.ScanFloat64()
	}
	return floats
}

func (io *IO) Println(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}
