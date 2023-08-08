package main

import (
	"encoding/json"
	"net/http"
	"quiz_assessment/data"

	"github.com/gorilla/mux"
)

var QuizResults []int

func getQuestionsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	jsonBytes, err := json.Marshal(data.Quiz)
	if err != nil {
		http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
		return
	}
	w.Write(jsonBytes)

}

func submitQuizResultsHandler(w http.ResponseWriter, r *http.Request) {
	var score int
	err := json.NewDecoder(r.Body).Decode(&score)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	QuizResults = append(QuizResults, score)
	// Return the updated QuizResults as the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(QuizResults)

}

func getResultsHandler(w http.ResponseWriter, r *http.Request) {
	// Return the  QuizResults as the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(QuizResults)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/questions", getQuestionsHandler).Methods("GET")
	router.HandleFunc("/results", getResultsHandler).Methods("GET")
	router.HandleFunc("/submit", submitQuizResultsHandler).Methods("POST")
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
