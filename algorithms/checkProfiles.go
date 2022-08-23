package algorithms

import (
	"Algo/globals"
	"fmt"
	"os"
	"sync"
	"time"
)

func ReadProfiles() {

	// TODO need to add logic for killing program
	// add logic to remove user when prompted from rest to stop
	//	the trading when removed

	var wg sync.WaitGroup

	for !globals.QuitProgram {
		time.Sleep(5 * time.Second)
		fmt.Println("")
		fmt.Print("-> Users trading : ")
		for e := globals.ProfilesToRun.Users.Front(); e != nil; e = e.Next() {
			// go StartStopCommand()
			userTemp := globals.User(e.Value.(globals.User))
			if userTemp.Status == "nt" {
				globals.UpdateStatus(userTemp)
				go GetPrice(start, &wg, chnPrices, quit[ss.Userid], runningTimeMin)

				// start thread here for the algo 1 min MA and a 5 min MA
				go algorithms.Ma15(start, &wg, chnPrices, quit[ss.Userid], runningTimeMin)
			}

			if userTemp.Status == "t" {
				fmt.Print(userTemp.Userid, " | ")
			}
		}
		fmt.Println()

	}
	wg.Done()
	os.Exit(4)
}
