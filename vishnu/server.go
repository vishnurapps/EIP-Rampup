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
	Content string `json:"Content"`
}

type Articles []Article

var articles = Articles{
	Article{Title: "Cricket", Desc: "Worldcup", Content: "India won worldcup"},
}

func listArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside allArticles")
	json.NewEncoder(w).Encode(articles)
}

func postArticles(w http.ResponseWriter, r *http.Request) {
	var newArticle Article
	fmt.Println("Inside getArticles")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newArticle)
	if err != nil {
		fmt.Println("Inside error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(newArticle)
	// body, _ := ioutil.ReadAll(r.Body)
	// _ = json.Unmarshal([]byte(body), &newArticle)
	// fmt.Println(newArticle)
	articles = append(articles, newArticle)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Homepage")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", listArticles).Methods("GET")
	myRouter.HandleFunc("/all", postArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":9091", myRouter))
}

func main() {
	handleRequests()
}
