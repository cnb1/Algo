package main

import (
	"Algo/api"
	"Algo/globals"
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func main() {
	fmt.Println("(A) Number of goroutines : ", runtime.NumGoroutine())
	fmt.Println(globals.GetGID())
	fmt.Println("starting rest client...")
	http.HandleFunc("/", api.Trading)
	http.HandleFunc("/stop", api.StopTrading)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
