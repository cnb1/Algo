package api

import (
	"Algo/globals"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
)

type StartStop struct {
	Strategy string  `json:"strategy"`
	Userid   string  `json:"userid"`
	Money    float64 `json:"money"`
	Command  string  `json:"command"`
}

func Trading(w http.ResponseWriter, r *http.Request) {
	var startStop StartStop
	fmt.Println("(AA) Number of goroutines : ", runtime.NumGoroutine(), " id ", globals.GetGID())
	err := json.NewDecoder(r.Body).Decode(&startStop)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "GET method requested"}`))
		// get money price from trading application
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "POST method requested"}`))
		StartStopCommand(&startStop)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}
}

func StopTrading(w http.ResponseWriter, r *http.Request) {
	var startStop StartStop
	fmt.Println("(BB) Number of goroutines : ", runtime.NumGoroutine(), " id ", globals.GetGID())
	err := json.NewDecoder(r.Body).Decode(&startStop)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("(E) Number of goroutines : ", runtime.NumGoroutine(), " id ", globals.GetGID())
	fmt.Println("stopping for user ", startStop.Userid)
	globals.QuitPrice[startStop.Userid] <- true
	globals.QuitAlgo[startStop.Userid] <- true
	fmt.Println("map price : ", globals.QuitPrice)
	fmt.Println("map algo : ", globals.QuitAlgo)
}
