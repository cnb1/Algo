package main

import (
	"Algo/algorithms"
	"Algo/api"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	fmt.Println("starting rest client...")

	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/article", createNewArticle)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in creating article")
	// r.ParseForm()
	// for key, value := range r.Form {
	// 	fmt.Println(key, value)
	// }
	// fmt.Fprintln(w, "success message")
	// decoder := json.NewDecoder(r.Body)
	// var t Message
	// err := decoder.Decode(&t)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println(t)
	// fmt.Println(r)

	// get the body of our POST request
	// return the string response containing the request body
	// w.Header().Set("Content-Type", "application/json")

	// var m Message

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	// err := json.NewDecoder(r.Body).Decode(&m)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	fmt.Println("there was an error")
	// 	return
	// }

	// Do something with the Person struct...
	// fmt.Fprintf(w, "Message: %+v", m)
	// fmt.Println(m.message)

	// err := r.ParseForm()
	// if err == nil {
	// 	fmt.Println("fail")
	// }

	// temp := r.ParseForm("message")
	// fmt.Println(temp)

	// Change the response depending on the method being requested
	// switch r.Method {
	// case "GET":
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte(`{"message": "GET method requested"}`))
	// case "POST":
	// 	fmt.Println("successful post")
	// 	w.WriteHeader(http.StatusCreated)
	// 	w.Write([]byte(`{"message": "POST method requested"}`))
	// default:
	// 	w.WriteHeader(http.StatusNotFound)
	// 	w.Write([]byte(`{"message": "Can't find method requested"}`))
	// }
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func runProgram() {
	const runningTimeMin = 180

	var wg sync.WaitGroup
	wg.Add(2)

	//channel
	chn := make(chan float64)

	start := time.Now()

	// start thread here for prices
	go api.GetPrice(start, &wg, chn, runningTimeMin)

	// start thread here for the algo 1 min MA and a 5 min MA
	go algorithms.Ma15(start, &wg, chn, runningTimeMin)

	wg.Wait()
	fmt.Println("Finished Trading")
}
