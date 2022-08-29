package api

import (
	"Algo/globals"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Price struct {
	Usd float64 `json:"usd"`
}

type Result struct {
	Crypto Price `json:"bitcoin"`
}

func GetPrice(start time.Time, wg *sync.WaitGroup, runningTimeMin time.Duration, userid string) {
	fmt.Println("getting prices...")
	end := start.Add(runningTimeMin * time.Minute)
	isQuit := false
	var r Result

	for !time.Now().After(end) && !isQuit {
		time.Sleep(4 * time.Second)

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

		decodererr := json.NewDecoder(resp.Body).Decode(&r)

		if decodererr != nil {
			panic(err)
		}

		// fmt.Println("decoder is ", decoder)
		if err != nil {
			log.Fatalln(err)
			fmt.Println("request error")
			time.Sleep(60 * time.Second)
			continue
		}
		//We Read the response body on the line below.
		if err != nil {
			log.Fatalln(err)
			fmt.Println("request error")
			time.Sleep(60 * time.Second)
			continue
		}

		if decodererr != nil {
			fmt.Println("Decoder error")
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Println("sending price : ", r.Crypto.Usd, " at time : ", time.Now(), " FOR USER [", userid, "]")

		globals.Prices[userid] <- (float64(r.Crypto.Usd))
	}

}
