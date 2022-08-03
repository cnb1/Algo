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
	end := start.Add(10 * time.Second)
	l := list.New()

	for {
		if time.Now().After(end) {
			fmt.Println("time to cancel ma15")
			wg.Done()
		}
		val := <-ch

		var price Price
		price.price = val
		price.date = time.Now()

		l.PushBack(price)

		fmt.Println("--------------before-----------------------")
		for e := l.Front(); e != nil; e = e.Next() {
			item := Price(e.Value.(Price))
			fmt.Println("date: ", item.date, " price: ", item.price)
		}

		checkInterval(*l, 4, wg)

		fmt.Println("--------------after-----------------------")
		for e := l.Front(); e != nil; e = e.Next() {
			item := Price(e.Value.(Price))
			fmt.Println("date: ", item.date, " price: ", item.price)
		}
		fmt.Println("\n\n\n")
	}

}

func checkInterval(l list.List, interval int, wg *sync.WaitGroup) {
	var start time.Time = time.Now().Add(-time.Second * time.Duration(interval))
	fmt.Println("start: ", start)
	fmt.Println()
	for e := l.Front(); e != nil; e = e.Next() {
		item := Price(e.Value.(Price))

		fmt.Println("       ", item.date)

		if item.date.Before(start) {
			// remove the first element and recaclutlate the average
			// why is the remove not removing

			fmt.Println("its before")
			l.Remove(e)

		} else {
			fmt.Println("\n[this is not before its fine]\n")
			// do nothing
			break
		}
	}
}
