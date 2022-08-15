package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Price struct {
	Usd int `json:"usd"`
}

type Result struct {
	Crypto Price `json:"bitcoin"`
}

func GetPrice(start time.Time, wg *sync.WaitGroup, ch chan float64, chnQuit chan bool,
	runningTimeMin time.Duration) {
	fmt.Println("getting prices...")
	end := start.Add(runningTimeMin * time.Minute)

	for {
		time.Sleep(4 * time.Second)

		select {
		case isQuit := <-chnQuit:
			fmt.Println("is going to quit ", isQuit)
			if isQuit {
				wg.Done()
			}
		default:
		}

		if time.Now().After(end) {
			fmt.Println("stopping prices...")
			wg.Done()
		}

		resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd")

		if err != nil {
			log.Fatalln(err)
			fmt.Println("request error")
			time.Sleep(60 * time.Second)
			continue
		}
		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
			fmt.Println("request error")
			time.Sleep(60 * time.Second)
			continue
		}

		var result Result
		json.Unmarshal(body, &result)

		ch <- (float64(result.Crypto.Usd))
		chnQuit <- false
	}
}
