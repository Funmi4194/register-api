package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"strings"
	"time"

	"github.com/gorilla/mux"
)

//models

//User structs

type User struct {
	Name       string `json:"name"`
	Profession string `json:"profession"`
	Location   string `json:"location"`
}

var register []User

func getDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r) //get paramas

	for _, detail := range register {
		if detail.Name == p["name"] {
			//detail.Name = strings.ToLower(detail.Name)
			json.NewEncoder(w).Encode(detail)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func getAllDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatiom/json")
	json.NewEncoder(w).Encode((register))
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	register = append(register, user)
	json.NewEncoder(w).Encode(user)
}

func main() {
	//initializing Router
	h := mux.NewRouter()

	//mock data to test database
	register = append(register, User{Name: "Funmi", Profession: "trader", Location: "Abuja"})
	register = append(register, User{Name: "Wunmi", Profession: "doctor", Location: "lagos"})
	//creating a server with two endpoints and route Hnadlers
	h.HandleFunc("/api/register/", createAccount).Methods("POST")
	h.HandleFunc("/api/user/{name}", getDetail).Methods("GET")
	h.HandleFunc("/api/registered", getAllDetails).Methods("GET")

	s := &http.Server{
		Addr:           ":8080",
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Starting server on http://localhost:%d\n", 8080)
	log.Fatal(s.ListenAndServe())

}
