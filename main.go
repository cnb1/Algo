package main

import (
	"Algo/algorithms"
	"Algo/api"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("starting rest client...")

	go algorithms.ReadProfiles()
	http.HandleFunc("/start", api.AddUser)
	http.HandleFunc("/stop", api.RemoveUser)
	http.HandleFunc("/money", api.GetMoneyForUser)
	http.HandleFunc("/kill", api.KillProgram)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
