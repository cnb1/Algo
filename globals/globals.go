package globals

import (
	"bytes"
	"runtime"
	"strconv"
)

var QuitPrice = make(map[string](chan bool))
var QuitAlgo = make(map[string](chan bool))
var Prices = make(map[string](chan float64))

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
