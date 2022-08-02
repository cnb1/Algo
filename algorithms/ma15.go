package algorithms

import (
	"fmt"
	"sync"
	"time"
)

func Ma15(start time.Time, wg *sync.WaitGroup, ch <-chan string) {
	fmt.Println("ma15 execute")
	end := start.Add(1 * time.Hour)

	for {
		if time.Now().After(end) {
			fmt.Println("time to cancel ma15")
			wg.Done()
		}
		val := <-ch

		fmt.Println("ma15 printing the price...")
		fmt.Println(val)

	}

}
