package api

import (
	"Algo/globals"
	"fmt"
	"os"
	"time"
)

const runningTimeMin = 180

func ReadProfiles() {

	// TODO need to add logic for killing program
	// add logic to remove user when prompted from rest to stop
	//	the trading when removed

	// var wg sync.WaitGroup

	for !globals.QuitProgram {
		time.Sleep(5 * time.Second)
		fmt.Println("")
		fmt.Print("-> Users trading : ")
		for e := globals.ProfilesToRun.Users.Front(); e != nil; e = e.Next() {
			userTemp := globals.User(e.Value.(globals.User))

			if userTemp.Status == "rm" {
				removeUser(userTemp.Userid)

			} else if userTemp.Status == "nt" {
				globals.UpdateStatus(userTemp)
				go Start(userTemp.Userid, userTemp.Money, userTemp.Strategy)
			} else if userTemp.Status == "t" {
				fmt.Print(userTemp.Userid, " | ")
			}
		}
		fmt.Println()

	}
	os.Exit(4)
}

func removeUser(userid string) {
	user, err := globals.GetUser(userid)

	if err != nil {
		fmt.Println("message : User ", userid, " doesnt exist in context")
	} else {

		fmt.Print("message : ")
		fmt.Print("user ", userid, " was removed")
		fmt.Println(", money value is : ", user.Money)

		globals.RemoveUser(userid)

	}

	// need to check if a user is in the prices channel
	if globals.CheckUserInPrices(userid) {
		fmt.Println("Stopping the program")
		Stop(userid)
	}

}
