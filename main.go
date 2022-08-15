package main

import (
	"Algo/api"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("starting rest client...")
	http.HandleFunc("/", api.Trading)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
