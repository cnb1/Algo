package main

import (
	"Algo/algorithms"
	"Algo/api"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	fmt.Println("starting rest client...")
	http.HandleFunc("/", api.Trading)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func runProgram() {
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
