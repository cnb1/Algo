package globals

import (
	"bytes"
	"container/list"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"strconv"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
type PriceAPI struct {
	MongoDB string `json:"mongoDB"`
}

var QuitPrice = make(chan bool)
var QuitAlgo = make(map[string](chan bool))
var Prices = make(map[string](chan float64))
var Money = make(map[string]float64)
var MoneyInitial = make(map[string]float64)
var ProfilesToRun = Profile{}
var QuitProgram = false
var Mongo mongo.Client
var Context context.Context

const sizeMaxProfiles int = 5

func GetClient() mongo.Client {
	return Mongo
}

func GetContext() context.Context {
	return Context
}

func SetClient() {
	fmt.Println("setting client")
	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload PriceAPI
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	MongoTemp, err := mongo.NewClient(options.Client().ApplyURI(payload.MongoDB))
	if err != nil {
		fmt.Println("panic 1")
		log.Fatal(err)
	}
	Mongo = *MongoTemp
	Context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = MongoTemp.Connect(Context)
	if err != nil {
		fmt.Println("panic 2")
		log.Fatal(err)
	}
	// defer Mongo.Disconnect(Context)
	err = Mongo.Ping(Context, readpref.Primary())
	if err != nil {
		fmt.Println("panic 3")
		log.Fatal(err)
	}
}

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
		MoneyInitial[userid] = money
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
			delete(MoneyInitial, userid)
			delete(Money, userid)
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
