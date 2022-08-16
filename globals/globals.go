package globals

import (
	"bytes"
	"container/list"
	"runtime"
	"strconv"
	"sync"
)

type Profile struct {
	Users list.List
	mu    sync.Mutex
}

var QuitPrice = make(map[string](chan bool))
var QuitAlgo = make(map[string](chan bool))
var Prices = make(map[string](chan float64))
var Money = make(map[string]float64)
var ProfilesToRun = Profile{Users: *list.New()}

const sizeMaxProfiles int = 5

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func GetProfile() string {
	ProfilesToRun.mu.Lock()

	prof := ProfilesToRun.Users.Front().Value.(string)

	ProfilesToRun.mu.Unlock()

	return prof
}

func AddProfile(userid string) bool {

	ProfilesToRun.mu.Lock()
	var ret bool
	if ProfilesToRun.Users.Len() < sizeMaxProfiles {
		ProfilesToRun.Users.PushBack(userid)
		ret = true
	} else {
		ret = false
	}

	ProfilesToRun.mu.Unlock()

	return ret
}
