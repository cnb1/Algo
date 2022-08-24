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
	// fmt.Println("(B) Number of goroutines : ", runtime.NumGoroutine(), " id ", globals.GetGID())

	if val, ok := globals.QuitPrice[userid]; ok {
		//do something here
		fmt.Println(val)
	} else {
		fmt.Println("adding key")
		globals.QuitPrice[userid] = make(chan bool)
		globals.QuitAlgo[userid] = make(chan bool)
		globals.Prices[userid] = make(chan float64)
	}

	var wg sync.WaitGroup

	fmt.Println("starting stragegy : ", strategy)
	const runningTimeMin = 180

	wg.Add(2)

	start := time.Now()

	// start thread here for prices
	go GetPrice(start, &wg, runningTimeMin, userid)

	// start thread here for the algo 1 min MA and a 5 min MA
	go algorithms.Ma15(start, &wg, runningTimeMin, userid)

	wg.Wait()
	fmt.Println("Finished Trading for user : ", userid)

}

func Stop(userid string) {
	// fmt.Println("(E) Number of goroutines : ", runtime.NumGoroutine(), " id ", globals.GetGID())
	fmt.Println("[STOP FUNC] stopping for user ", userid)
	globals.QuitPrice[userid] <- true

}
