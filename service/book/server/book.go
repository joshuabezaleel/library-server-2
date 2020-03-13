package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/joshuabezaleel/library-server-2/service/book/book"

	"github.com/gorilla/mux"
)

type bookHandler struct {
	bookService book.Service
}

func (handler *bookHandler) registerRouter(router *mux.Router) {
	router.HandleFunc("/books", handler.getAllBooks).Methods("GET")
}

func (handler *bookHandler) getAllBooks(w http.ResponseWriter, r *http.Request) {
	var topicIDs []int

	queryValues := r.URL.Query()

	topics := queryValues.Get("topics")
	if topics != "" {
		topicIDsString := strings.Split(topics, " ")
		for _, topicIDString := range topicIDsString {
			topicID, _ := strconv.Atoi(topicIDString)

			topicIDs = append(topicIDs, topicID)
		}
	}

	books, err := handler.bookService.GetAll(topicIDs)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, books)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"Error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
