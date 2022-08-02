package algorithms

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

func Ma15(start time.Time, wg *sync.WaitGroup, ch <-chan int) {
	fmt.Println("ma15 execute")
	end := start.Add(1 * time.Hour)
	l := list.New()

	for {
		if time.Now().After(end) {
			fmt.Println("time to cancel ma15")
			wg.Done()
		}
		val := <-ch

		fmt.Println("ma15 printing the price...")
		fmt.Println(val)

		l.PushBack(val)

		// add price to 1 min slice and a 5 min slice
		// how to comapre the 1 min to 5 min times and how to
		// have a rolling average and removing the last object

	}

}
