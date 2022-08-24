package globals

import (
	"bytes"
	"container/list"
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

type Profile struct {
	Users list.List
	mu    sync.Mutex
}

type User struct {
	Userid   string
	Money    float64
	Strategy string
	Status   string
}

var QuitPrice = make(map[string](chan bool))
var QuitAlgo = make(map[string](chan bool))
var Prices = make(map[string](chan float64))
var Money = make(map[string]float64)
var ProfilesToRun = Profile{}
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

func GetUser(userid string) (User, error) {
	fmt.Println("get user print")
	PrintProfiles()
	ProfilesToRun.mu.Lock()

	var user User
	var err error
	for e := ProfilesToRun.Users.Front(); e != nil; e = e.Next() {
		if e.Value.(User).Userid == userid {
			ProfilesToRun.mu.Unlock()
			return e.Value.(User), nil
		}
	}
	ProfilesToRun.mu.Unlock()
	user, err = User{}, errors.New("Nothing found")
	return user, err
}

func AddProfile(userid string, money float64, strategy string) bool {

	ProfilesToRun.mu.Lock()
	var ret bool
	if ProfilesToRun.Users.Len() < sizeMaxProfiles {
		profTemp := User{Userid: userid, Money: money, Strategy: strategy,
			Status: "nt"}
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
		if e.Value.(User).Userid == userid {
			ret = true
			ProfilesToRun.Users.Remove(e)
		}
	}
	ProfilesToRun.mu.Unlock()

	return ret
}

func UpdateStatus(user User) {
	ProfilesToRun.mu.Lock()
	for e := ProfilesToRun.Users.Front(); e != nil; e = e.Next() {
		if e.Value.(User).Userid == user.Userid {
			item := User{Userid: user.Userid, Money: user.Money, Strategy: user.Strategy,
				Status: "t"}
			ProfilesToRun.Users.Remove(e)
			ProfilesToRun.Users.PushBack(item)
		}
	}
	ProfilesToRun.mu.Unlock()
}

func UpdateStatusToRemove(userid string) {
	ProfilesToRun.mu.Lock()
	for e := ProfilesToRun.Users.Front(); e != nil; e = e.Next() {
		if e.Value.(User).Userid == userid {
			item := User{Userid: userid, Money: e.Value.(User).Money, Strategy: e.Value.(User).Strategy,
				Status: "rm"}
			ProfilesToRun.Users.Remove(e)
			ProfilesToRun.Users.PushBack(item)
		}
	}
	ProfilesToRun.mu.Unlock()
}

func PrintProfiles() {
	ProfilesToRun.mu.Lock()
	fmt.Println("Printing the profiles...")
	for e := ProfilesToRun.Users.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(User))
	}
	ProfilesToRun.mu.Unlock()
}

func KillProgram() {
	QuitProgram = true
}

func GetQuitProgram() bool {
	return QuitProgram
}

func CheckUserInPrices(userid string) bool {
	_, ok := Prices[userid]
	if ok {
		return true
	} else {
		fmt.Println("The user was not found in prices")
		return false
	}
}
