package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type API struct {
	WelcomeMessage string `json:"message"`
}

type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

var users = []User{
	User{1, "Emre", "Öney", 24},
	User{2, "Yusuf", "Öney", 29},
	User{3, "Burak", "Öney", 25},
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", handlerHomePage)
	router.HandleFunc("/users", handlerUsersPage).Methods("GET")
	router.HandleFunc("/users/{id}", HandlerGetUser).Methods("GET")
	router.HandleFunc("/users/delete/{id}", handlerDeleteUser).Methods("DEL")
	router.HandleFunc("/users/create", handlerCreateUser).Methods("POST")
	router.HandleFunc("/users/update/{id}", handlerUpdateUser).Methods("PUT")

	http.ListenAndServe(":8080", router)

}

func handlerHomePage(w http.ResponseWriter, r *http.Request) {
	message := API{"Welcome API Exercises"}
	data, err := json.Marshal(message)
	checkError(err)
	fmt.Fprintf(w, string(data))
}

func handlerUsersPage(w http.ResponseWriter, r *http.Request) {
	message := users

	data, err := json.Marshal(message)
	checkError(err)
	fmt.Fprintf(w, string(data))
}

func HandlerGetUser(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	userID, err := strconv.Atoi(variables["id"])
	checkError(err)

	var targetUser User
	message := getAllUsers()

	for _, user := range message {
		if user.ID == userID {
			targetUser = user
			break
		}
	}

	data, err := json.Marshal(targetUser)
	fmt.Fprintf(w, string(data))
}

func handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	userID, err := strconv.Atoi(variables["id"])
	checkError(err)

	users := getAllUsers()

	var index int 

	for i, user := range users {
		if user.ID == userID {
			index = i
			break
		}
	}

	users = append(users[:index], users[index+1:]...)

	jsonData, err := json.Marshal(users)
	fmt.Fprintf(w, string(jsonData))

}

func handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	checkError(err)

	newUser.ID = len(users) + 1

	users = append(users, newUser)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)

}

func handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	variable := mux.Vars(r)
	userID, err := strconv.Atoi(variable["id"])
	checkError(err)

	var targetUser User

	for _, user := range users {
		if user.ID == userID {
			targetUser = user
			break
		}
	}

	err2 := json.NewDecoder(r.Body).Decode(&targetUser)
	checkError(err2)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(targetUser)

	users[userID-1] = targetUser
}

func checkError(err error) {
	if err != nil {
		fmt.Println("HATA!", err.Error())
	}
}

func getAllUsers() []User {
	message := []User{
		User{1, "Emre", "Öney", 24},
		User{2, "Yusuf", "Öney", 29},
		User{3, "Burak", "Öney", 25},
	}
	return message
}
