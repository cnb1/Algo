package algorithms

import (
	"Algo/globals"
	"time"
)

func readProfiles() {

	for !globals.QuitProgram {
		time.Sleep(5 * time.Second)

		for e := globals.ProfilesToRun.Users.Front(); e != nil; e = e.Next() {
			// go StartStopCommand()
		}

	}
}
