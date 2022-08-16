package api

import (
	"Algo/globals"
	"encoding/json"
	"fmt"
	"net/http"
)

type StartStop struct {
	Strategy string  `json:"strategy"`
	Userid   string  `json:"userid"`
	Money    float64 `json:"money"`
	Command  string  `json:"command"`
}

func Trading(w http.ResponseWriter, r *http.Request) {
	var startStop StartStop
	// fmt.Println("(AA) Number of goroutines : ", runtime.NumGoroutine(), " id ", globals.GetGID())
	err := json.NewDecoder(r.Body).Decode(&startStop)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "POST method requested"}`))
		StartStopCommand(&startStop)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}
}

func GetMoneyForUser(w http.ResponseWriter, r *http.Request) {
	var startStop StartStop
	// fmt.Println("(AA) Number of goroutines : ", runtime.NumGoroutine(), " id ", globals.GetGID())
	err := json.NewDecoder(r.Body).Decode(&startStop)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "POST":
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Money : ")
		fmt.Fprint(w, globals.Money[startStop.Userid])
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}

}
