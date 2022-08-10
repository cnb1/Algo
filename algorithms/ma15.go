package algorithms

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type Price struct {
	price float64
	date  time.Time
}

type Average struct {
	avg   float64
	sum   float64
	total float64
}

const intervalsmall = 60
const intervallarge = 300
const runningTimeMin = 5

var money = 1000000

func Ma15(start time.Time, wg *sync.WaitGroup, ch <-chan float64) {
	fmt.Println("ma15 execute")
	end := start.Add(runningTimeMin * time.Hour)
	l := list.New()
	lLarge := list.New()

	avg := Average{avg: 0.0, sum: 0.0, total: 0.0}
	avgLarge := Average{avg: 0.0, sum: 0.0, total: 0.0}

	for {
		if time.Now().After(end) {
			fmt.Println()
			fmt.Println("time to cancel ma15")
			wg.Done()
		}
		val := <-ch

		var price Price
		price.price = val
		price.date = time.Now()

		avg.sum = avg.sum + price.price
		avg.total++
		avg.avg = avg.sum / avg.total

		avgLarge.sum = avgLarge.sum + price.price
		avgLarge.total++
		avgLarge.avg = avgLarge.sum / avgLarge.total

		fmt.Println("-----------------------------")
		fmt.Println("New Price ", price.price, " -> average: ", avg.avg, " total: ", avg.total, " sum: ", avg.sum)
		l.PushBack(price)
		lLarge.PushBack(price)

		checkInterval(l, lLarge, intervalsmall, intervallarge, wg, &avg, &avgLarge)

		fmt.Println("Average is now: ", avg.avg)
		fmt.Println("Average Large is now: ", avgLarge.avg)
		fmt.Println("-----------------------------")

		// now is the time to decide whether or not to buy or sell
	}

}

func checkInterval(l *list.List, lLarge *list.List, intervalsmall int, intervallarge int, wg *sync.WaitGroup, avg *Average, avgLarge *Average) {
	var start time.Time = time.Now().Add(-time.Second * time.Duration(intervalsmall))
	var startLarge time.Time = time.Now().Add(-time.Second * time.Duration(intervallarge))

	for e := l.Front(); e != nil; e = e.Next() {
		item := Price(e.Value.(Price))

		if item.date.Before(start) {

			fmt.Println("its before")
			fmt.Println("average: ", avg.avg, " total: ", avg.total, " sum: ", avg.sum)
			fmt.Println()
			avg.total--
			avg.sum = avg.sum - item.price
			avg.avg = avg.sum / avg.total

			l.Remove(e)

		} else {
			break
		}
	}

	for e := lLarge.Front(); e != nil; e = e.Next() {
		item := Price(e.Value.(Price))

		if item.date.Before(startLarge) {

			fmt.Println("its before")
			fmt.Println("average large: ", avgLarge.avg, " total: ", avgLarge.total, " sum: ", avgLarge.sum)
			fmt.Println()
			avgLarge.total--
			avgLarge.sum = avgLarge.sum - item.price
			avgLarge.avg = avgLarge.sum / avgLarge.total

			lLarge.Remove(e)

		} else {
			break
		}
	}
}
