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

func GetPrice(start time.Time, wg *sync.WaitGroup, ch chan int) {
	fmt.Println("getting prices...")
	end := start.Add(1 * time.Hour)

	for {
		time.Sleep(2 * time.Second)

		if time.Now().After(end) {
			fmt.Println("time to cacnel")
			wg.Done()
		}

		resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd")
		if err != nil {
			log.Fatalln(err)
		}
		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var result Result
		json.Unmarshal(body, &result)

		ch <- result.Crypto.Usd
	}
}
