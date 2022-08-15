package api

import (
	"Algo/algorithms"
	"fmt"
	"sync"
	"time"
)

func StartStopCommand(ss *StartStop, quit map[string](chan bool)) {
	command := ss.Command

	switch command {
	case "start":

		fmt.Println("map : ", quit)

		if val, ok := quit[ss.Userid]; ok {
			//do something here
			fmt.Println(val)
		} else {
			fmt.Println("adding key")
			quit[ss.Userid] = make(chan bool)
		}

		fmt.Println("map : ", quit)

		var wg sync.WaitGroup

		fmt.Println("starting ", ss.Strategy)
		const runningTimeMin = 180

		wg.Add(2)

		//channel
		chnPrices := make(chan float64)

		start := time.Now()

		// start thread here for prices
		go GetPrice(start, &wg, chnPrices, quit[ss.Userid], runningTimeMin)

		// start thread here for the algo 1 min MA and a 5 min MA
		go algorithms.Ma15(start, &wg, chnPrices, quit[ss.Userid], runningTimeMin)

		wg.Wait()
		wg.Done()
		fmt.Println("Finished Trading")
	case "stop":
		fmt.Println("stopping for user ", ss.Userid)
		quit[ss.Userid] <- true
		quit[ss.Userid] <- true
	}

}
