package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func GetPrice(start time.Time, wg *sync.WaitGroup) {
	fmt.Println("getting prices...")
	end := start.Add(8 * time.Second)

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
		//Convert the body to type string
		sb := string(body)
		log.Printf(sb)
	}
}
