package riot

type MatchStore interface {
	Get(matchId string) (*Match, error)
	Set(matchId string, match *Match) error
	Delete(matchId string, match *Match)
}

// Stores the configuration for a cache that stores data related to the Riot
// APIs.
type Cache struct {
	Matches MatchStore
}

// Retrieves a match if one is stored in the cache, otherwise returning an
// error. An error may also be returned if there was some issue communicating
// with the underlying storage system.
func (c *Cache) LoadMatch(matchId string) (*Match, error) {
	return c.Matches.Get(matchId)
}

// Stores a match in the cache. An error may be returned if there was some
// problem communicating with the underlying storage system.
func (c *Cache) SaveMatch(matchId string, match *Match) error {
	return c.Matches.Set(matchId, match)
}
