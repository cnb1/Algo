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

type Position struct {
	buy        float64
	close      float64
	position   string
	inPosition bool
	newPrice   bool
}

const intervalsmall = 60
const intervallarge = 300

var money float64 = 1000000

func Ma15(start time.Time, wg *sync.WaitGroup, ch <-chan float64, runningTimeMin time.Duration) {
	// Gets the end value that the back value needs to be greater than
	t1 := start.Add(time.Second * intervallarge)
	ifCanTrade := false
	position := Position{buy: 0, close: 0, inPosition: false, newPrice: false}

	fmt.Println("ma15 execute")
	end := start.Add(runningTimeMin * time.Minute)
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

		// check if change in price
		if position.inPosition && position.buy != price.price {
			// fmt.Println("[there is a new price]")
			position.newPrice = true
		}

		// start averages
		avg.sum = avg.sum + price.price
		avg.total++
		avg.avg = avg.sum / avg.total

		avgLarge.sum = avgLarge.sum + price.price
		avgLarge.total++
		avgLarge.avg = avgLarge.sum / avgLarge.total

		l.PushBack(price)
		lLarge.PushBack(price)

		checkInterval(l, lLarge, intervalsmall, intervallarge, wg, &avg, &avgLarge)

		// check if the end value in lLarge is after the t1 value
		if !ifCanTrade && lLarge.Back().Value.(Price).date.After(t1) {
			fmt.Println("IT IS AFTER")
			ifCanTrade = true
		}

		if ifCanTrade {

			if position.inPosition {
				if position.position == "long" && position.newPrice {

					// fmt.Println("long and new price")
					// check if we need to close of a loss or a gain
					if price.price < position.buy || price.price >= position.close {

						fmt.Print("		closing long position, before money: ", money, " | at price : ", price.price, " Time: ", time.Now())

						money += (price.price - position.buy)

						position.position = "none"
						position.inPosition = false
						position.buy = 0
						position.close = 0
						position.newPrice = false

						fmt.Println("		after long money: ", money)
						fmt.Println()
						fmt.Println()
					}

				} else if position.position == "short" && position.newPrice { // "short"
					// fmt.Println("short and new price")
					if price.price > position.buy || price.price <= position.close {
						// close for a loss
						fmt.Print("		closing short position, before money: ", money, " | at price : ", price.price, " Time: ", time.Now())

						money += (position.buy - price.price)

						position.position = "none"
						position.inPosition = false
						position.buy = 0
						position.close = 0
						position.newPrice = false

						fmt.Println("		after short money: ", money)
						fmt.Println()
						fmt.Println()
					}
				}
			} else if avg.avg > avgLarge.avg {
				// bullish
				fmt.Println("buy bitcoin | small avg : ", avg.avg, "   large avg : ", avgLarge.avg)
				position.position = "long"
				position.buy = price.price
				position.close = price.price * 1.01
				position.inPosition = true
				position.newPrice = false

				fmt.Println("position: ", position.position, " buy at: ", position.buy, " close at ", position.close, " Time: ", time.Now())
				fmt.Println()
				fmt.Println()

			} else if avg.avg < avgLarge.avg {
				// bearish
				fmt.Println("short bitcoin | small avg : ", avg.avg, "   large avg : ", avgLarge.avg)
				position.position = "short"
				position.buy = price.price
				position.close = price.price * 0.99
				position.inPosition = true
				position.newPrice = false

				fmt.Println("position: ", position.position, " buy at: ", position.buy, " close at ", position.close, " Time: ", time.Now())
				fmt.Println()
				fmt.Println()

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

			avgLarge.total--
			avgLarge.sum = avgLarge.sum - item.price
			avgLarge.avg = avgLarge.sum / avgLarge.total

			lLarge.Remove(e)

		} else {
			break
		}
	}
}
