package algorithms

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type Price struct {
	price int
	date  time.Time
}

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

		var price Price
		price.price = val
		price.date = time.Now()

		l.PushBack(price)

		checkInterval(*l, 1)
	}

}

func checkInterval(l list.List, interval int) {
	var start time.Time = time.Now().Add(-time.Minute * time.Duration(interval))

	for e := l.Front(); e != nil; e = e.Next() {
		item := Price(e.Value.(Price))

		if item.date.Before(start) {
			fmt.Println("it is before this needs to be removed")
			// remove the first element and recaclutlate the average
		} else {
			fmt.Println("this is not before its fine")
			// do nothing
		}
	}

}
