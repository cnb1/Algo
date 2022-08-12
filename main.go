package main

import (
	"Algo/algorithms"
	"Algo/api"
	"fmt"
	"sync"
	"time"
)

func main() {

	const runningTimeMin = 180

	var wg sync.WaitGroup
	wg.Add(2)

	//channel
	chn := make(chan float64)

	start := time.Now()

	// start thread here for prices
	go api.GetPrice(start, &wg, chn, runningTimeMin)

	// start thread here for the algo 1 min MA and a 5 min MA
	go algorithms.Ma15(start, &wg, chn, runningTimeMin)

	wg.Wait()
	fmt.Println("Finished Trading")
}
