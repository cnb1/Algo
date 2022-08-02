package algorithms

import (
	"fmt"
	"sync"
	"time"
)

func Ma15(start time.Time, wg *sync.WaitGroup) {
	fmt.Println("ma15 execute")
	wg.Done()

	// create e algo that calculates the 1 min and 5 min MAs
	// makes trades on these
	// when 1 goes below 5 then sell
	// when 1 crosses 5 then buy
	// create a chanel to communicate in bewtwwn
}
