package main

import (
	"Algo/algorithms"
	"Algo/api"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	start := time.Now()

	// start thread here for prices
	go api.GetPrice(start, &wg)

	// start thread here for the algo 1 min MA and a 5 min MA
	go algorithms.Ma15(start, &wg)

	wg.Wait()
	fmt.Println("Finished Trading")
}
