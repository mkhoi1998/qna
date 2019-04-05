package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k0kubun/pp"
)

// our main function
func main() {
	port := "8000"
	router := mux.NewRouter()
	router.HandleFunc("/question", CreateQuestion).Methods("POST")
	router.HandleFunc("/question", GetQuestionList).Methods("GET")
	router.HandleFunc("/answer", AnswerQuestion).Methods("POST")
	pp.Println("Server is running on ", port)
	defer pp.Println("Server is stopped.")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// CreateQuestionRequest ..
type CreateQuestionRequest struct {
	Content string `json:"content,omitempty"`
}

// CreateQuestion ..
func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var q CreateQuestionRequest
	err := json.NewDecoder(r.Body).Decode(&q)
	if err != nil {
		// TODO: return error 400
		return
	}

	pp.Println("create Q")
	// qRs := Question{
	// 	Content: q.Content,
	// }
	// json.NewEncoder(w).Encode(qRs)
}

// GetQuestionList ..
func GetQuestionList(w http.ResponseWriter, r *http.Request) {
	pp.Println("Get Q")
}

// AnswerQuestion ..
func AnswerQuestion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Answer Q")
}
