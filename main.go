package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Struct informações das pessoas
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Struct informação de endereço
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

// GetPeople mostra todos os contatos da variável people
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// GetPerson função para mostrar só uma pessoa na requisição
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// CreatePerson função para criar um novo contato
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// DeletePerson função que deleta um contato
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

// função principal
func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Ygor", Lastname: "Pereira", Address: &Address{City: "Bariloche", State: "Rio Negro"}})
	people = append(people, Person{ID: "2", Firstname: "Eduardo", Lastname: "Dicarte", Address: &Address{City: "Belo Horizonte", State: "Minas Gerais"}})
	router.HandleFunc("/contato", GetPeople).Methods("GET")            // fazendo a requisição
	router.HandleFunc("/contato/{id}", GetPerson).Methods("GET")       // fazendo a requisição
	router.HandleFunc("/contato/{id}", CreatePerson).Methods("POST")   // fazendo a requisição
	router.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE") // fazendo a requisição
	log.Fatal(http.ListenAndServe(":8000", router))                    // definindo a porta :8000
}
