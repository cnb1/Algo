package main

import (
	"Algo/api"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("starting rest client...")

	go api.ReadProfiles()
	http.HandleFunc("/start", api.AddUser)
	http.HandleFunc("/stop", api.UpdateToRemoveStatus)
	http.HandleFunc("/money", api.GetMoneyForUser)
	http.HandleFunc("/kill", api.KillProgram)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
