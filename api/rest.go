package api

import (
	"Algo/globals"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	PortID string `json:"portfolioId"`
}

type Response struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type Value struct {
	Price float64
	Date  time.Time
	ID    primitive.ObjectID `bson:"_id"`
}

type ResultValue struct {
	ID    primitive.ObjectID
	Value float64
}

func AddUser(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	var newUser AddStruct

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check if user already exists
	_, errUsr := globals.GetUser(newUser.Userid)

	if errUsr == nil {
		fmt.Println("USER IS ALREADY TRADING! ")
		w.Header().Set("Content-Type", "application/json")
		resptemp := Response{Message: "User [" + newUser.Userid + "] is already trading",
			Success: false}
		json.NewEncoder(w).Encode(resptemp)
		return
	}

	switch r.Method {
	case "POST":
		// w.WriteHeader(http.StatusCreated)
		if globals.AddProfile(newUser.Userid, newUser.Money, newUser.Strategy) {
			fmt.Println("Added Profile for {user: ", newUser.Userid, " strategy: ", newUser.Strategy,
				" money to trade: ", newUser.Money, "}")
			w.Header().Set("Content-Type", "application/json")

			resptemp := Response{Message: "User [" + newUser.Userid + "] was added",
				Success: true}
			json.NewEncoder(w).Encode(resptemp)
			return
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
	setupResponse(&w, r)
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
		_, err := globals.GetUser(removeUser.Userid)

		if err != nil {
			fmt.Fprint(w, "message : User ", removeUser.Userid, " doesnt exist in context")
		} else {
			w.Header().Set("Content-Type", "application/json")
			resp := make(map[string]string)

			profitLoss := CalculatePriceLoss(removeUser.Userid)

			resp["message"] = "user " + removeUser.Userid + " was updated to rm" +
				", money value is : " + strconv.FormatFloat(profitLoss, 'f', 2, 32)
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
	setupResponse(&w, r)

	var userJSON UserIDStruct
	// fmt.Println("(AA) Number of goroutines : ", runtime.NumGoroutine(), " id ", globals.GetGID())
	err := json.NewDecoder(r.Body).Decode(&userJSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "POST":
		user, err := globals.GetUser(userJSON.Userid)

		if err != nil {
			fmt.Println("this user is not trading.....")
			fmt.Fprint(w, "message : User ", user.Userid, " doesnt exist in context")
		} else {
			w.Header().Set("Content-Type", "application/json")

			resp := make(map[string]string)
			tempMoney := CalculatePriceLoss(user.Userid)
			resp["message"] = strconv.FormatFloat(tempMoney, 'f', 2, 32)
			jsonResp, err := json.Marshal(resp)

			fmt.Println("Getting money for the user: ", user.Userid)
			fmt.Println(jsonResp)

			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}

			// send this money to the mongo db
			fmt.Println("about to send money 1")

			updateValueAndValueHistory(userJSON, user.Money, tempMoney)

			w.Write(jsonResp)
		}
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

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func updateValueAndValueHistory(userJSON UserIDStruct, money float64, pl float64) {

	// Updating the value History
	client := globals.GetClient()

	// ctx := globals.GetContext()

	fmt.Println("---------------------------------------------")
	var result ResultValue
	objID, _ := primitive.ObjectIDFromHex(userJSON.PortID)
	usercol := client.Database("CryptoWebApp").Collection("userportfolios")
	match := bson.M{"_id": objID}

	// get current value
	usercol.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&result)

	// update the value
	fmt.Println("The prfit loss is .. : ", pl)
	fmt.Println("Result val is ", result.Value)

	newValue := pl + result.Value - 100000000

	updateVal := bson.M{"$set": bson.M{"value": newValue}}

	fmt.Println()
	usercol.UpdateOne(context.TODO(), match, updateVal)

	val := Value{
		Price: newValue,
		Date:  time.Now(),
		ID:    primitive.NewObjectID(),
	}
	addVal := bson.M{"$push": bson.M{"valueHistory": val}}
	usercol.UpdateOne(context.TODO(), match, addVal)

}
