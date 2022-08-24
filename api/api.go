package api

import (
	"Algo/globals"
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

func GetPrice(start time.Time, wg *sync.WaitGroup, runningTimeMin time.Duration, userid string) {
	fmt.Println("getting prices...")
	end := start.Add(runningTimeMin * time.Minute)
	isQuit := false

	for !time.Now().After(end) && !isQuit {
		// fmt.Println("(Price GR) Number of goroutines : ", runtime.NumGoroutine(), " id ", globals.GetGID())
		time.Sleep(4 * time.Second)
		// fmt.Println("getting prices ....")

		select {
		case isQuit = <-globals.QuitPrice[userid]:
			fmt.Println("is going to quit prices : ", isQuit)
			if isQuit {
				fmt.Println("After print wg done")
				wg.Done()
				return
			}
		default:
		}

		if time.Now().After(end) {
			fmt.Println("stopping prices...")
			wg.Done()
			fmt.Println("After prices wg done")
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

		fmt.Println()
		fmt.Println("sending price : ", result.Crypto.Usd, " at time : ", time.Now(), " FOR USER [", userid, "]")

		// If we can recieve on the channel then it is NOT closed

		globals.Prices[userid] <- (float64(result.Crypto.Usd))
	}

}
