package riot

// Represents a League of Legends game server. In the Riot API, this is called
// a "region." However, the Riot API also refers to other things as regions 
// too. To avoid confusion, this package calls a region like "na1" or "kr" a 
// server--since it is, in fact, representing a game server--while the term
// "region" is reserved for geographical regions like "americas" or "asia."
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
