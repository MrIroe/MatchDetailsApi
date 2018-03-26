package objects

type MatchDetails struct {
	PhysicalDamageDealt            int64 `json:"physicalDamageDealt"`
	MagicDamageDealt               int64 `json:"magicDamageDealt"`
	Deaths                         int   `json:"deaths"`
	Win                            bool  `json:"win"`
	AltarsCaptured                 int   `json:"altarsCaptured"`
	TotalDamageDealt               int64 `json:"totalDamageDealt"`
	MagicDamageDealtToChampions    int64 `json:"magicDamageDealtToChampions"`
	DamageDealtToObjectives        int64 `json:"damageDealtToObjectives"`
	TeamObjective                  int   `json:"teamObjective"`
	WardsPlaced                    int   `json:"wardsPlaced"`
	TurretKills                    int   `json:"turretKills"`
	ChampLevel                     int   `json:"champLevel"`
	GoldEarned                     int   `json:"goldEarned"`
	MagicalDamageTaken             int64 `json:"magicalDamageTaken"`
	Kills                          int   `json:"kills"`
	Assists                        int   `json:"assists"`
	NeutralMinionsKilled           int   `json:"neutralMinionsKilled"`
	PhysicalDamageDealtToChampions int64 `json:"physicalDamageDealtToChampions"`
	GoldSpent                      int   `json:"goldSpent"`
	TrueDamageDealt                int64 `json:"trueDamageDealt"`
	TrueDamageDealtToChampions     int64 `json:"trueDamageDealtToChampions"`
	ParticipantId                  int   `json:"participantId"`
	TotalHeal                      int64 `json:"totalHeal"`
	TotalMinionsKilled             int   `json:"totalMinionsKilled"`
	LargestMultiKill               int   `json:"largestMultiKill"`
	TotalDamageDealtToChampions    int64 `json:"totalDamageDealtToChampions"`
	TotalDamageTaken               int64 `json:"totalDamageTaken"`
	KillingSprees                  int   `json:"killingSprees"`
	PhysicalDamageTaken            int64 `json:"physicalDamageTaken"`
	CreepsPerMin                   int   `json:"creepsPerMin"`
	CsDiffPerMin                   int   `json:"csDiffPerMin"`
	ExpDiffPerMin                  int   `json:"expDiffPerMin"`
}

type ParticipantStatsDto struct {
	PhysicalDamageDealt            int64 `json:"physicalDamageDealt"`
	MagicDamageDealt               int64 `json:"magicDamageDealt"`
	Deaths                         int   `json:"deaths"`
	Win                            bool  `json:"win"`
	AltarsCaptured                 int   `json:"altarsCaptured"`
	TotalDamageDealt               int64 `json:"totalDamageDealt"`
	MagicDamageDealtToChampions    int64 `json:"magicDamageDealtToChampions"`
	DamageDealtToObjectives        int64 `json:"damageDealtToObjectives"`
	TeamObjective                  int   `json:"teamObjective"`
	WardsPlaced                    int   `json:"wardsPlaced"`
	TurretKills                    int   `json:"turretKills"`
	ChampLevel                     int   `json:"champLevel"`
	GoldEarned                     int   `json:"goldEarned"`
	MagicalDamageTaken             int64 `json:"magicalDamageTaken"`
	Kills                          int   `json:"kills"`
	Assists                        int   `json:"assists"`
	NeutralMinionsKilled           int   `json:"neutralMinionsKilled"`
	PhysicalDamageDealtToChampions int64 `json:"physicalDamageDealtToChampions"`
	GoldSpent                      int   `json:"goldSpent"`
	TrueDamageDealt                int64 `json:"trueDamageDealt"`
	TrueDamageDealtToChampions     int64 `json:"trueDamageDealtToChampions"`
	ParticipantId                  int   `json:"participantId"`
	TotalHeal                      int64 `json:"totalHeal"`
	TotalMinionsKilled             int   `json:"totalMinionsKilled"`
	LargestMultiKill               int   `json:"largestMultiKill"`
	TotalDamageDealtToChampions    int64 `json:"totalDamageDealtToChampions"`
	TotalDamageTaken               int64 `json:"totalDamageTaken"`
	KillingSprees                  int   `json:"killingSprees"`
	PhysicalDamageTaken            int64 `json:"physicalDamageTaken"`
}

type ParticipantTimelineDto struct {
	Lane                        string             `json:"lane"`
	ParticipantId               int                `json:"participantId"`
	CsDiffPerMinDeltas          map[string]float64 `json:"csDiffPerMinDeltas"`
	GoldPerMinDeltas            map[string]float64 `json:"goldPerMinDeltas"`
	XpDiffPerMinDeltas          map[string]float64 `json:"xpDiffPerMinDeltas"`
	CreepsPerMinDeltas          map[string]float64 `json:"creepsPerMinDeltas"`
	XpPerMinDeltas              map[string]float64 `json:"xpPerMinDeltas"`
	Role                        string             `json:"role"`
	DamageTakenDiffPerMinDeltas map[string]float64 `json:"damageTakenDiffPerMinDeltas"`
	DamageTakenPerMinDeltas     map[string]float64 `json:"damageTakenPerMinDeltas"`
}
