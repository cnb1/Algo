package api

import (
	"Algo/algorithms"
	"fmt"
	"sync"
	"time"
)

func StartStopCommand(ss *StartStop) {
	command := ss.Command
	chnQuit := make(chan bool)

	switch command {
	case "start":
		fmt.Println("starting ", ss.Strategy)
		const runningTimeMin = 180

		var wg sync.WaitGroup
		wg.Add(2)

		//channel
		chnPrices := make(chan float64)

		start := time.Now()

		// start thread here for prices
		go GetPrice(start, &wg, chnPrices, chnQuit, runningTimeMin)

		// start thread here for the algo 1 min MA and a 5 min MA
		go algorithms.Ma15(start, &wg, chnPrices, chnQuit, runningTimeMin)

		wg.Wait()
		fmt.Println("Finished Trading")
	case "stop":
		fmt.Println("stopping for user ", ss.Userid)
		chnQuit <- false
		chnQuit <- false
	}

}
