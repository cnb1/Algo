package api

import (
	"Algo/globals"
	"encoding/json"
	"fmt"
	"net/http"
)

type AddStruct struct {
	Strategy string  `json:"strategy"`
	Userid   string  `json:"userid"`
	Money    float64 `json:"money"`
}

type RemoveStruct struct {
	Userid string `json:"userid"`
}

type UserIDStruct struct {
	Userid string `json:"userid"`
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var newUser AddStruct
	// fmt.Println("(AA) Number of goroutines : ", runtime.NumGoroutine(), " id ", globals.GetGID())
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "POST":
		w.WriteHeader(http.StatusCreated)
		if globals.AddProfile(newUser.Userid, newUser.Money, newUser.Strategy) {
			fmt.Fprint(w, "message : ")
			fmt.Fprint(w, "user ", newUser.Userid, " was added")
		} else {
			fmt.Fprint(w, "message : ")
			fmt.Fprint(w, "user ", newUser.Userid, " was not added, max users added")
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	var removeUser RemoveStruct
	// fmt.Println("(AA) Number of goroutines : ", runtime.NumGoroutine(), " id ", globals.GetGID())
	err := json.NewDecoder(r.Body).Decode(&removeUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "POST":
		w.WriteHeader(http.StatusCreated)

		user, err := globals.GetUser(removeUser.Userid)

		if err != nil {
			fmt.Fprint(w, "message : User ", removeUser.Userid, " doesnt exist in context")
		} else {

			fmt.Fprint(w, "message : ")
			fmt.Fprint(w, "user ", removeUser.Userid, " was removed")
			fmt.Fprint(w, ", money value is : ", user.Money)

			globals.RemoveUser(removeUser.Userid)

		}

		// need to check if a user is in the prices channel
		if globals.CheckUserInPrices(removeUser.Userid) {
			fmt.Println("Stopping the program")
			Stop(removeUser.Userid)
		}

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}
}

func GetMoneyForUser(w http.ResponseWriter, r *http.Request) {
	var user UserIDStruct
	// fmt.Println("(AA) Number of goroutines : ", runtime.NumGoroutine(), " id ", globals.GetGID())
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "POST":
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Money : ")
		fmt.Fprint(w, globals.Money[user.Userid])
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}

}

func KillProgram(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Killed Program")
		globals.KillProgram()
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}
}
