package mongo

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	obj "matchDetailsApi/objects"
	util "matchDetailsApi/util"
)

var primarySess *mgo.Session

const entityDbName = "Entities"
const summonerColName = "Summoners"
const summonerStatsColName = "SummonerStats"

func InitMongoRepo() {
	var err error
	primarySess, err = mgo.Dial(util.GetConfigValue("mongoAddress"))
	if err != nil {
		fmt.Println(errors.Wrap(err, "Couldnt connect to mongo"))
		os.Exit(-1)
	}

	primarySess.SetMode(mgo.Monotonic, true)
}

func GetSummonerStats(accountId int64) ([]obj.SummonerMatchStats, error) {
	localSess := *primarySess.Clone()
	defer localSess.Close()
	c := localSess.DB(entityDbName).C(summonerStatsColName)

	var summonerStats []obj.SummonerMatchStats
	err := c.Find(bson.M{"AccountId": accountId}).All(&summonerStats)
	if err != nil {
		return []obj.SummonerMatchStats{}, errors.Wrap(err, "Error getting SummonerMatchStats")
	}

	return summonerStats, nil
}
