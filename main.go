package main

import (
	
	"encoding/json"
	"github.com/gorilla/mux"		
	"net/http"	
	"log"
)

type Person struct{
	ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

var people []Person


func main() {
	people = append(people,  Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people,  Person{ID: "2", Firstname: "Jon", Lastname: "oe", Address: &Address{City: "CityX", State: "Stae X"}})
	people = append(people,  Person{ID: "3", Firstname: "on", Lastname: "oe", Address: &Address{City: "ityX", State: "Stae X"}})
	people = append(people,  Person{ID: "4", Firstname: "Jn", Lastname: "oe", Address: &Address{City: "CityX", State: "Ste X"}})
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	for _, person:= range people {
		if person.ID == params["id"] {
			 json.NewEncoder(w).Encode(person);
			 break;
		}
	}
}


func CreatePerson(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
    var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)
    person.ID = params["id"]
    people = append(people, person)
    json.NewEncoder(w).Encode(people)	
}

func DeletePerson(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	var position int = -1
	for index, item := range people {
		if item.ID == params["id"]{			
			position = index
		}	
	}
	if(position != -1 ){
		people = append(people[:position], people[position+1:]...)
	}

}

