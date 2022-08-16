package main

import (
	"Algo/api"
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func main() {
	fmt.Println("(A) Number of goroutines : ", runtime.NumGoroutine())
	fmt.Println("starting rest client...")
	http.HandleFunc("/start", api.AddUser)
	http.HandleFunc("/stop", api.RemoveUser)
	http.HandleFunc("/money", api.GetMoneyForUser)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
