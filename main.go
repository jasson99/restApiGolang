package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Hello world"},
	}
	fmt.Println("EndPoint Hit: All articles endpoint")
	json.NewEncoder(w).Encode(articles)
}

func allPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EndPoint Hit: All Post articles endpoint")
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is Homepage EndPoint")
}

func handleRequests() {
	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/", homePage)
	myrouter.HandleFunc("/articles", allArticles).Methods("GET")
	myrouter.HandleFunc("/articles", allPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8085", myrouter))
}

func main() {
	handleRequests()
}
