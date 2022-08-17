package algorithms

import (
	"Algo/globals"
	"fmt"
	"os"
	"time"
)

func ReadProfiles() {

	// TODO need to add logic for killing program
	// add logic to remove user when prompted from rest to stop
	//	the trading when removed

	for !globals.QuitProgram {
		time.Sleep(5 * time.Second)
		fmt.Println("")
		fmt.Print("-> Users trading : ")
		for e := globals.ProfilesToRun.Users.Front(); e != nil; e = e.Next() {
			// go StartStopCommand()
			userTemp := globals.User(e.Value.(globals.User))
			if userTemp.Status == "nt" {
				globals.UpdateStatus(userTemp)
			}

			if userTemp.Status == "t" {
				fmt.Print(userTemp.Userid, " | ")
			}
		}
		fmt.Println()

	}

	os.Exit(4)
}
