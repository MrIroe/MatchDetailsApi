package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	mongo "matchDetailsApi/mongo"
)

func RegisterMatchDetailsHandlers(router *mux.Router, route string) {
	matchDetails := router.PathPrefix(route).Subrouter()
	matchDetails.HandleFunc("/matchDetails/byAccID/{accountId}", GetMatchDetails)
}

func GetMatchDetails(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	accountId := vars["accountId"]
	intAccountId, err := strconv.ParseInt(accountId, 10, 64)
	if err != nil {
		errLog := fmt.Sprintf("Error parsing string: %v", err)
		WriteResponse(w, 400, false, errLog, nil)
		return
	}

	summonerMatchStats, err := mongo.GetSummonerStats(intAccountId)
	if err != nil {
		errLog := fmt.Sprintf("Error in GetMatchDetails: %v", err)
		WriteResponse(w, 400, false, errLog, nil)
		return
	}
	fmt.Println(summonerMatchStats)
}
