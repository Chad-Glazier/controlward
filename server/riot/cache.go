package riot

import "errors"

type MatchStore interface {
	GetMatch(matchId string) (*Match, error)
	SetMatch(matchId string, match *Match) error
	DeleteMatch(matchId string)
}

// Stores the configuration for a cache that stores data related to the Riot
// APIs.
type Cache struct {
	matches MatchStore
}

// Retrieves a match if one is stored in the cache, otherwise returning an
// error. An error may also be returned if there was some issue communicating
// with the underlying storage system.
func (c *Cache) LoadMatch(matchId string) (*Match, error) {
	if c.matches == nil {
		return nil, errors.New("matches cache is not set")
	}

	return c.matches.GetMatch(matchId)
}

// Stores a match in the cache. An error may be returned if there was some
// problem communicating with the underlying storage system.
func (c *Cache) SaveMatch(matchId string, match *Match) error {
	if c.matches == nil {
		return errors.New("matches cache is not set")
	}

	return c.matches.SetMatch(matchId, match)
}

// Creates a new cache with the given stores. You can pass nil without causing
// any problems.
func NewCache(matches MatchStore) Cache {
	return Cache{
		matches: matches,
	}
}
