package objects

type SummonerMatchStats struct {
	AccountId  int64                  `json:"accountId" bson:"AccountId"`
	MatchId    int64                  `json:"matchId" bson:"MatchId"`
	ChampionId int                    `json:"championId" bson:"ChampionId"`
	Stats      ParticipantStatsDto    `json:"stats" bson:"Stats"`
	Timeline   ParticipantTimelineDto `json:"timeline" bson:"Timeline"`
}
