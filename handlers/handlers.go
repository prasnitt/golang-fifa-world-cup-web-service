package handlers

import (
	"golang-fifa-world-cup-web-service/data"
	"net/http"
)

// RootHandler returns an empty body status code
func RootHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNoContent)
}

// ListWinners returns winners from the list
func ListWinners(res http.ResponseWriter, req *http.Request) {
	year := req.URL.Query().Get("year")
	res.Header().Set("Content-Type", "application/json")

	if year == "" {

		winners, err := data.ListAllJSON()
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.Write(winners)
	} else {
		winner, err := data.ListAllByYear(year)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
		}
		res.Write(winner)
	}
}

// AddNewWinner adds new winner to the list
func AddNewWinner(res http.ResponseWriter, req *http.Request) {

}

// WinnersHandler is the dispatcher for all /winners URL
func WinnersHandler(res http.ResponseWriter, req *http.Request) {

}
