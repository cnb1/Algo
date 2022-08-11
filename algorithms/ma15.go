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

const intervalsmall = 10 //60
const intervallarge = 30 //300
const runningTimeMin = 5

var money = 1000000

func Ma15(start time.Time, wg *sync.WaitGroup, ch <-chan float64) {
	// Gets the end value that the back value needs to be greater than
	t1 := start.Add(time.Second * intervallarge)
	ifCanTrade := false

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
		fmt.Println("small : New Price ", price.price, " -> average: ", avg.avg, " total: ", avg.total, " sum: ", avg.sum)
		fmt.Println("large : New Price ", price.price, " -> average: ", avgLarge.avg, " total: ", avgLarge.total, " sum: ", avgLarge.sum)
		fmt.Println()
		fmt.Println()

		l.PushBack(price)
		lLarge.PushBack(price)

		checkInterval(l, lLarge, intervalsmall, intervallarge, wg, &avg, &avgLarge)

		fmt.Println("Average is now: ", avg.avg)
		fmt.Println("Average Large is now: ", avgLarge.avg)
		fmt.Println("-----------------------------")
		fmt.Println()
		fmt.Println()

		// check if the end value in lLarge is after the t1 value
		if !ifCanTrade && lLarge.Back().Value.(Price).date.After(t1) {
			fmt.Println("IT IS AFTER")
			ifCanTrade = true
		}

		if ifCanTrade {
			/*
				if a position is being taken then set a position and a price and
				dont take a new one unless the actual price is at break even from start
			*/
			// you can check the averages now small and large
			fmt.Println("CHECK THE AVERAGES")

			if avg.avg > avgLarge.avg {
				// bullish
				fmt.Println("buy bitcoin")
			} else if avg.avg < avgLarge.avg {
				// bearish
				fmt.Println("short bitcoin")
			} else {
				// do nothing
				fmt.Println("dont buy anything")
			}
		}

	}

}

func checkInterval(l *list.List, lLarge *list.List, intervalsmall int, intervallarge int, wg *sync.WaitGroup, avg *Average, avgLarge *Average) {
	var start time.Time = time.Now().Add(-time.Second * time.Duration(intervalsmall))
	var startLarge time.Time = time.Now().Add(-time.Second * time.Duration(intervallarge))

	for e := l.Front(); e != nil; e = e.Next() {
		item := Price(e.Value.(Price))

		if item.date.Before(start) {

			fmt.Println("	small its before")
			fmt.Println("	average: ", avg.avg, " total: ", avg.total, " sum: ", avg.sum)
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

			fmt.Println("	large its before")
			fmt.Println("	average large: ", avgLarge.avg, " total: ", avgLarge.total, " sum: ", avgLarge.sum)
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
