package api

import (
	"net/http"
)

type StartStop struct {
	Strategy string  `json:"strategy"`
	Userid   string  `json:"userid"`
	Money    float64 `json:"money"`
	Command  string  `json:"command"`
}

func HandleRequests() {
	http.HandleFunc("/", trading)
}

func trading(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

}
