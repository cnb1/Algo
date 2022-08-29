package api

import (
	"Algo/globals"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type AddStruct struct {
	Strategy string  `json:"strategy"`
	Userid   string  `json:"userid"`
	Money    float64 `json:"money"`
}

type RemoveStruct struct {
	Userid string `json:"userid"`
}

type UserIDStruct struct {
	Userid string `json:"userid"`
}

type Response struct {
	Message string `json:"message"`
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var newUser AddStruct

	// fmt.Println("(AA) Number of goroutines : ", runtime.NumGoroutine(), " id ", globals.GetGID())
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "POST":
		w.WriteHeader(http.StatusCreated)
		if globals.AddProfile(newUser.Userid, newUser.Money, newUser.Strategy) {
			w.Header().Set("Content-Type", "application/json")
			// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			// w.Header().Set("Access-Control-Allow-Origin", "*")
			resp := make(map[string]string)
			resp["message"] = "user " + newUser.Userid + " was added"
			jsonResp, err := json.Marshal(&resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			in := []byte(`{"id":1,"name":"test","context":{"key1":"value1","key2":2}}`)
			w.Write(in)
			fmt.Println(jsonResp)
			return
			// var resptemp Response
			// errMar := json.Unmarshal(jsonResp, &resptemp)

			// if errMar != nil {
			// 	fmt.Println("error in unmarshalling")
			// }

			// fmt.Println(resptemp)
			// fmt.Fprint(w, "message : ")
			// fmt.Fprint(w, "user ", newUser.Userid, " was added")
		} else {
			fmt.Fprint(w, "message : ")
			fmt.Fprint(w, "user ", newUser.Userid, " was not added, max users added")
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}
}

func UpdateToRemoveStatus(w http.ResponseWriter, r *http.Request) {
	var removeUser RemoveStruct
	// fmt.Println("(AA) Number of goroutines : ", runtime.NumGoroutine(), " id ", globals.GetGID())
	err := json.NewDecoder(r.Body).Decode(&removeUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "POST":
		w.WriteHeader(http.StatusCreated)
		user, err := globals.GetUser(removeUser.Userid)

		if err != nil {
			fmt.Fprint(w, "message : User ", removeUser.Userid, " doesnt exist in context")
		} else {
			w.Header().Set("Content-Type", "application/json")
			resp := make(map[string]string)
			resp["message"] = "user " + removeUser.Userid + " was updated to rm" + ", money value is : " + strconv.FormatFloat(user.Money, 'E', -1, 32)
			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			w.Write(jsonResp)

			globals.UpdateStatusToRemove(removeUser.Userid)

		}

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}

	fmt.Println("[UPDATE] leaving the remove update")

}

func GetMoneyForUser(w http.ResponseWriter, r *http.Request) {
	var user UserIDStruct
	// fmt.Println("(AA) Number of goroutines : ", runtime.NumGoroutine(), " id ", globals.GetGID())
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "POST":
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Money : ")
		fmt.Fprint(w, globals.Money[user.Userid])
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}

}

func KillProgram(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Killed Program")
		globals.KillProgram()
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}
}
