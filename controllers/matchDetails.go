package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterMatchDetailsHandlers(router *mux.Router, route string) {
	matchDetails := router.PathPrefix(route).Subrouter()
	matchDetails.HandleFunc("/matchDetails/byAccID/{accountId}", GetMatchDetails)
}

func GetMatchDetails(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	accountId := vars["accountId"]

	fmt.Println(accountId)
}
