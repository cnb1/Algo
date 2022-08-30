package api

import (
	"Algo/algorithms"
	"Algo/globals"
	"fmt"
	"sync"
	"time"
)

func Start(userid string, money float64, strategy string) {
	globals.Money[userid] = money

	fmt.Println("adding key")
	globals.QuitAlgo[userid] = make(chan bool)
	globals.Prices[userid] = make(chan float64)

	var wg sync.WaitGroup

	fmt.Println("starting stragegy : ", strategy)
	const runningTimeMin = 180

	wg.Add(1)

	start := time.Now()

	// start thread here for the algo 1 min MA and a 5 min MA
	go algorithms.Ma15(start, &wg, runningTimeMin, userid)

	wg.Wait()
	fmt.Println("Finished Trading for user : ", userid)

}

func Stop(userid string) {
	// fmt.Println("(E) Number of goroutines : ", runtime.NumGoroutine(), " id ", globals.GetGID())
	fmt.Println("[STOP FUNC] stopping for user ", userid)
}
