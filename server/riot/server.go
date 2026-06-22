package riot

// A Server represents a League of Legends game server. In the Riot API, this
// is called a "region." However, the Riot API also refers to other things as
// regions too. To avoid confusion, this package calls a region like "na1" or
// "kr" a server--since it is, in fact, representing a game server--while the
// term "region" is reserved for geographical regions like "americas" or
// "asia."
type Server string

const (
	ServerBR1  Server = "br1"
	ServerEUN1 Server = "eun1"
	ServerEUW1 Server = "euw1"
	ServerJP1  Server = "jp1"
	ServerKR   Server = "kr"
	ServerLA1  Server = "la1"
	ServerLA2  Server = "la2"
	ServerME1  Server = "me1"
	ServerNA1  Server = "na1"
	ServerOC1  Server = "oc1"
	ServerPBE1 Server = "pbe1"
	ServerRU   Server = "ru"
	ServerSG2  Server = "sg2"
	ServerTR1  Server = "tr1"
	ServerVN2  Server = "vn2"
)

// Most League of Legends endpoints in the Riot API are divided by the server,
// but some more endpoints (like accessing a Riot account) are based on more 
// general, geographic regions. Namely, "americas," "asia," and "europe."
type Region string

const (
	RegionAmericas Region = "americas"
	RegionAsia     Region = "asia"
	RegionEurope   Region = "europe"
)

// Represents a game for some API endpoints. Either "lol" or "tft".
type Game string

const (
	GameLeagueOfLegends  Game = "lol"
	GameTeamfightTactics Game = "tft"
)
