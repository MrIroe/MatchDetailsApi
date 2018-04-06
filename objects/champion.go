package objects

type ChampionInfo struct {
	Title string `json:"title" bson:"Title"`
	Id    int    `json:"id" bson:"Id"`
	Key   string `json:"key" bson:"Key"`
	Name  string `json:"name" bson:"Name"`
}
