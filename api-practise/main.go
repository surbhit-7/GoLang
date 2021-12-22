package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type user struct {
	Name   string `json:"name"`
	UserID string `json:"userId"`
}

func homePage(w http.ResponseWriter, r *http.Request) {

	users := []user{
		{"Surbhit", "1234"},
		{"abcd", "2532"},
	}

	byteArr, _ := json.Marshal(users)
	fmt.Fprintf(w, "%s", byteArr)
}
func handleRequests() {

	http.HandleFunc("/users", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
	//log.Println("hello")
}

func main() {

	handleRequests()

}
