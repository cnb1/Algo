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
	Money float64
	mu    sync.Mutex
}

type User struct {
	userid   string
	money    float64
	strategy string
}

var QuitPrice = make(map[string](chan bool))
var QuitAlgo = make(map[string](chan bool))
var Prices = make(map[string](chan float64))
var Money = make(map[string]float64)
var ProfilesToRun = Profile{Users: *list.New()}
var QuitProgram = false

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

func AddProfile(userid string, money float64, strategy string) bool {

	ProfilesToRun.mu.Lock()
	var ret bool
	if ProfilesToRun.Users.Len() < sizeMaxProfiles {
		profTemp := User{userid: userid, money: money, strategy: strategy}
		ProfilesToRun.Users.PushBack(profTemp)
		ret = true
	} else {
		ret = false
	}

	ProfilesToRun.mu.Unlock()

	return ret
}

func RemoveUser(userid string) bool {

	ProfilesToRun.mu.Lock()
	var ret bool = false
	for e := ProfilesToRun.Users.Front(); e != nil; e = e.Next() {
		if e.Value.(User).userid == userid {
			ret = true
			ProfilesToRun.Users.Remove(e)
		}
	}
	ProfilesToRun.mu.Unlock()

	return ret
}

func KillProgram() {
	QuitProgram = true
}

func GetQuitProgram() bool {
	return QuitProgram
}
