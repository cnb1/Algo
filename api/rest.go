package api

import (
	"encoding/json"
	"net/http"
)

type StartStop struct {
	Strategy string  `json:"strategy"`
	Userid   string  `json:"userid"`
	Money    float64 `json:"money"`
	Command  string  `json:"command"`
}

var quit = make(map[string](chan bool))

func Trading(w http.ResponseWriter, r *http.Request) {
	var startStop StartStop

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
		StartStopCommand(&startStop, quit)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}

}
