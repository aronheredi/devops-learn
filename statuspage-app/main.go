package main

import (
	"fmt"
	"log"
	"net/http"
)

func echoHealthJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Status", "200")
	jsonResponse := `{"status": "ok"}`
	fmt.Fprint(w, jsonResponse)
}

var CommitSHA string
var buildDate string

func echoInfoJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Status", "200")
	jsonResponse := fmt.Sprintf(`{"CommitSHA": "%s","BuildDate": "%s"}`, CommitSHA, buildDate)
	fmt.Fprint(w, jsonResponse)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/health", echoHealthJson)
	http.HandleFunc("/info", echoInfoJson)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
