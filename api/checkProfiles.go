package api

import (
	"Algo/globals"
	"fmt"
	"os"
	"sync"
	"time"
)

const runningTimeMin = 180

func ReadProfiles() {

	isOne := false

	for !globals.QuitProgram {

		time.Sleep(5 * time.Second)

		if !isOne && globals.ProfilesToRun.Users.Len() > 0 {
			isOne = true
			// kick off prices
			var wg sync.WaitGroup
			wg.Add(1)
			go GetPrice(&wg)
		} else if isOne && globals.ProfilesToRun.Users.Len() <= 0 {
			isOne = false
			// close prices thread by sending a quit is true channel
			fmt.Println("[Closing prices]")
			globals.QuitPrice <- true
		}

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

func CalculatePriceLoss(userid string) float64 {
	initial := globals.MoneyInitial[userid]
	current := globals.Money[userid]
	change := current - initial

	return change
}
