package riot

// A Server represents a League of Legends game server. In the Riot API, this
// is called a "region." However, the Riot API also refers to other things as
// regions too. To avoid confusion, this package calls a region like "na1" or
// "kr" a server--since it is, in fact, representing a game server--while the
// term "region" is reserved for geographical regions like "americas" or
// "asia."
type Server string

const (
	BR1  Server = "br1"
	EUN1 Server = "eun1"
	EUW1 Server = "euw1"
	JP1  Server = "jp1"
	KR   Server = "kr"
	LA1  Server = "la1"
	LA2  Server = "la2"
	ME1  Server = "me1"
	NA1  Server = "na1"
	OC1  Server = "oc1"
	PBE1 Server = "pbe1"
	RU   Server = "ru"
	SG2  Server = "sg2"
	TR1  Server = "tr1"
	VN2  Server = "vn2"
)

// Most League of Legends endpoints in the Riot API are divided by the server,
// but some more general endpoints (like accessing a Riot account) are based
// on more general, geographic regions. Namely, "americas," "asia," and
// "europe."
type Region string

const (
	Americas Region = "americas"
	Asia     Region = "asia"
	Europe   Region = "europe"
)

// Represents a game for some API endpoints. Either "lol" or "tft".
type Game string

const (
	LeagueOfLegends  Game = "lol"
	TeamfightTactics Game = "tft"
)
