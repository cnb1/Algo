package api

import (
	"Algo/algorithms"
	"Algo/globals"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func StartStopCommand(ss *StartStop) {
	command := ss.Command
	globals.Money[ss.Userid] = ss.Money
	// fmt.Println("(B) Number of goroutines : ", runtime.NumGoroutine(), " id ", globals.GetGID())

	switch command {
	case "start":

		// fmt.Println("Price map : ", globals.QuitPrice)
		// fmt.Println("Algo map : ", globals.QuitPrice)

		if val, ok := globals.QuitPrice[ss.Userid]; ok {
			//do something here
			fmt.Println(val)
		} else {
			fmt.Println("adding key")
			globals.QuitPrice[ss.Userid] = make(chan bool)
			globals.QuitAlgo[ss.Userid] = make(chan bool)
			globals.Prices[ss.Userid] = make(chan float64)

		}

		// fmt.Println("map price : ", globals.QuitPrice)
		// fmt.Println("map algo : ", globals.QuitAlgo)
		// fmt.Println("map prices : ", globals.Prices)

		var wg sync.WaitGroup

		fmt.Println("starting ", ss.Strategy)
		const runningTimeMin = 180

		wg.Add(2)

		start := time.Now()

		// fmt.Println("(C) Number of goroutines : ", runtime.NumGoroutine(), " id ", globals.GetGID())

		// start thread here for prices
		go GetPrice(start, &wg, runningTimeMin, ss.Userid)

		// start thread here for the algo 1 min MA and a 5 min MA
		go algorithms.Ma15(start, &wg, runningTimeMin, ss.Userid)

		// fmt.Println("(D) Number of goroutines : ", runtime.NumGoroutine(), " id ", globals.GetGID())

		wg.Wait()
		fmt.Println("Finished Trading")
		// fmt.Println("Number of goroutines : ", runtime.NumGoroutine())
	case "stop":
		fmt.Println("(E) Number of goroutines : ", runtime.NumGoroutine(), " id ", globals.GetGID())
		fmt.Println("stopping for user ", ss.Userid)
		globals.QuitPrice[ss.Userid] <- true
		globals.QuitAlgo[ss.Userid] <- true
		fmt.Println("map price : ", globals.QuitPrice)
		fmt.Println("map algo : ", globals.QuitAlgo)
	}

}
