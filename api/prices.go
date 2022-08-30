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
	Usd float64 `json:"usd"`
}

type Result struct {
	Crypto Price `json:"bitcoin"`
}

type PriceAPI struct {
	Url string `json:"url"`
}

func GetPrice(wg *sync.WaitGroup) {

	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	var payload PriceAPI
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	fmt.Println("url:", payload.Url)

	fmt.Println("[GETTING PRICES]")
	isQuit := false
	var r Result

	for !isQuit {
		time.Sleep(4 * time.Second)

		select {
		case isQuit = <-globals.QuitPrice:
			fmt.Println("[Is going to quit prices] : ", isQuit)
			if isQuit {
				fmt.Println("After print wg done")
				wg.Done()
				return
			}
		default:
		}

		resp, err := http.Get(payload.Url)

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

		priceToSend := (float64(r.Crypto.Usd))

		// for loop through the users
		for e := globals.ProfilesToRun.Users.Front(); e != nil; e = e.Next() {
			userTemp := globals.User(e.Value.(globals.User))
			fmt.Println("sending price : ", r.Crypto.Usd, " at time : ", time.Now(), " FOR USER [", userTemp.Userid, "]")
			globals.Prices[userTemp.Userid] <- priceToSend

		}

	}

}
