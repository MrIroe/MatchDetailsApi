package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	mongo "matchDetailsApi/mongo"
	obj "matchDetailsApi/objects"
	util "matchDetailsApi/util"
)

//Queue is a map of QueueTypes and their IDs
var Queue = map[string]int{
	"NormalBlind": 430,
	"NormalDraft": 400,
	"RankedSolo":  420,
	"RankedFlex":  440,
}

func RegisterMatchDetailsHandlers(router *mux.Router, route string) {
	matchDetails := router.PathPrefix(route).Subrouter()
	matchDetails.HandleFunc("/matchDetails/all/{accountId}", GetMatchAllDetails)
	matchDetails.HandleFunc("/matchDetails/byQueue", GetMatchDetailsByQueue).Methods("POST")
}

func GetMatchAllDetails(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	accountId := vars["accountId"]
	intAccountId, err := strconv.ParseInt(accountId, 10, 64)
	if err != nil {
		errLog := fmt.Sprintf("Error parsing string: %v", err)
		WriteResponse(w, 400, false, errLog, nil)
		return
	}

	summonerMatchStats, err := mongo.GetAllSummonerStats(intAccountId)
	if err != nil {
		errLog := fmt.Sprintf("Error in GetMatchDetails: %v", err)
		WriteResponse(w, 400, false, errLog, nil)
		return
	}

	var returnMatchDetails []obj.MatchDetails

	for _, stats := range summonerMatchStats {
		creeps, csDiff, expDiff := util.AverageDeltas(stats.Timeline)
		if err != nil {
			fmt.Printf("Error averaging deltas for matchId: %v", stats.MatchId)
		}

		matchDetails := obj.MatchDetails{
			PhysicalDamageDealt:            stats.Stats.PhysicalDamageDealt,
			MagicDamageDealt:               stats.Stats.MagicDamageDealt,
			Deaths:                         stats.Stats.Deaths,
			Win:                            stats.Stats.Win,
			AltarsCaptured:                 stats.Stats.AltarsCaptured,
			TotalDamageDealt:               stats.Stats.TotalDamageDealt,
			MagicDamageDealtToChampions:    stats.Stats.MagicDamageDealtToChampions,
			DamageDealtToObjectives:        stats.Stats.DamageDealtToObjectives,
			TeamObjective:                  stats.Stats.TeamObjective,
			WardsPlaced:                    stats.Stats.WardsPlaced,
			TurretKills:                    stats.Stats.TurretKills,
			ChampLevel:                     stats.Stats.ChampLevel,
			GoldEarned:                     stats.Stats.GoldEarned,
			MagicalDamageTaken:             stats.Stats.MagicalDamageTaken,
			Kills:                          stats.Stats.Kills,
			Assists:                        stats.Stats.Assists,
			NeutralMinionsKilled:           stats.Stats.NeutralMinionsKilled,
			PhysicalDamageDealtToChampions: stats.Stats.PhysicalDamageDealtToChampions,
			GoldSpent:                      stats.Stats.GoldSpent,
			TrueDamageDealt:                stats.Stats.TrueDamageDealt,
			TrueDamageDealtToChampions:     stats.Stats.TrueDamageDealtToChampions,
			ParticipantId:                  stats.Stats.ParticipantId,
			TotalHeal:                      stats.Stats.TotalHeal,
			TotalMinionsKilled:             stats.Stats.TotalMinionsKilled,
			LargestMultiKill:               stats.Stats.LargestMultiKill,
			TotalDamageDealtToChampions:    stats.Stats.TotalDamageDealtToChampions,
			TotalDamageTaken:               stats.Stats.TotalDamageTaken,
			KillingSprees:                  stats.Stats.KillingSprees,
			PhysicalDamageTaken:            stats.Stats.PhysicalDamageTaken,
			CreepsPerMin:                   creeps,
			CsDiffPerMin:                   csDiff,
			ExpDiffPerMin:                  expDiff,
		}

		returnMatchDetails = append(returnMatchDetails, matchDetails)

	}

	WriteResponse(w, 200, true, "", returnMatchDetails)
}

func GetMatchDetailsByQueue(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var filter obj.MatchDetailFilter
	decoder.Decode(&filter)
	accountId := filter.AccountId
	queue := Queue[filter.Queue]

	intAccountId, err := strconv.ParseInt(accountId, 10, 64)
	if err != nil {
		errLog := fmt.Sprintf("Error parsing string: %v", err)
		WriteResponse(w, 400, false, errLog, nil)
		return
	}

	summonerMatchStats, err := mongo.GetSummonerStatsByQueue(intAccountId, queue)
	if err != nil {
		errLog := fmt.Sprintf("Error in GetMatchDetails: %v", err)
		WriteResponse(w, 400, false, errLog, nil)
		return
	}

	var returnMatchDetails []obj.MatchDetails

	for _, stats := range summonerMatchStats {
		creeps, csDiff, expDiff := util.AverageDeltas(stats.Timeline)
		if err != nil {
			fmt.Printf("Error averaging deltas for matchId: %v", stats.MatchId)
		}

		matchDetails := obj.MatchDetails{
			PhysicalDamageDealt:            stats.Stats.PhysicalDamageDealt,
			MagicDamageDealt:               stats.Stats.MagicDamageDealt,
			Deaths:                         stats.Stats.Deaths,
			Win:                            stats.Stats.Win,
			AltarsCaptured:                 stats.Stats.AltarsCaptured,
			TotalDamageDealt:               stats.Stats.TotalDamageDealt,
			MagicDamageDealtToChampions:    stats.Stats.MagicDamageDealtToChampions,
			DamageDealtToObjectives:        stats.Stats.DamageDealtToObjectives,
			TeamObjective:                  stats.Stats.TeamObjective,
			WardsPlaced:                    stats.Stats.WardsPlaced,
			TurretKills:                    stats.Stats.TurretKills,
			ChampLevel:                     stats.Stats.ChampLevel,
			GoldEarned:                     stats.Stats.GoldEarned,
			MagicalDamageTaken:             stats.Stats.MagicalDamageTaken,
			Kills:                          stats.Stats.Kills,
			Assists:                        stats.Stats.Assists,
			NeutralMinionsKilled:           stats.Stats.NeutralMinionsKilled,
			PhysicalDamageDealtToChampions: stats.Stats.PhysicalDamageDealtToChampions,
			GoldSpent:                      stats.Stats.GoldSpent,
			TrueDamageDealt:                stats.Stats.TrueDamageDealt,
			TrueDamageDealtToChampions:     stats.Stats.TrueDamageDealtToChampions,
			ParticipantId:                  stats.Stats.ParticipantId,
			TotalHeal:                      stats.Stats.TotalHeal,
			TotalMinionsKilled:             stats.Stats.TotalMinionsKilled,
			LargestMultiKill:               stats.Stats.LargestMultiKill,
			TotalDamageDealtToChampions:    stats.Stats.TotalDamageDealtToChampions,
			TotalDamageTaken:               stats.Stats.TotalDamageTaken,
			KillingSprees:                  stats.Stats.KillingSprees,
			PhysicalDamageTaken:            stats.Stats.PhysicalDamageTaken,
			CreepsPerMin:                   creeps,
			CsDiffPerMin:                   csDiff,
			ExpDiffPerMin:                  expDiff,
		}

		returnMatchDetails = append(returnMatchDetails, matchDetails)

	}

	WriteResponse(w, 200, true, "", returnMatchDetails)
}
