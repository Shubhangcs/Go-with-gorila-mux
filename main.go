package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"time"
	"github.com/gorilla/mux"
)

type UserDetails struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

var users []UserDetails


func main(){
	router := mux.NewRouter()
	router.HandleFunc("/newuser" , createUser).Methods("POST")
	router.HandleFunc("/loginuser",loginUser).Methods("POST")
	router.HandleFunc("/getusers",getUserData).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000",router))
	fmt.Println("Server Is Running On Port 9000...")
}

func createUser(w http.ResponseWriter , r *http.Request){
	time.Sleep(2 * time.Second)
	var newUser UserDetails
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	users = append(users, newUser)
	json.NewEncoder(w).Encode(newUser)
}

func loginUser(w http.ResponseWriter , r *http.Request){
	var isUserPresent bool

	var existingUser UserDetails
	_ = json.NewDecoder(r.Body).Decode(&existingUser)
	isUserPresent = slices.Contains(users , existingUser)
	if isUserPresent{
		fmt.Fprint(w,"Logged In Successfully.")
	}else{
		fmt.Fprint(w,"User Does not exist.")
	}
}

func getUserData(w http.ResponseWriter , r *http.Request)  {
	json.NewEncoder(w).Encode(users)
}