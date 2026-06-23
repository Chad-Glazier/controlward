/*
This package implements stores for Riot data by using the system memory.
*/
package memstore

import (
	"errors"
	"sync"
	"time"

	"github.com/Chad-Glazier/controlward/riot"
)

type MatchStore struct {
	Capacity   uint
	matches    map[string]*riot.Match
	lastAccess map[string]int64 // maps matchId -> last access timestamp
	mu         sync.Mutex
}

func NewMatchStore(capacity uint) *MatchStore {
	return &MatchStore{
		Capacity: capacity,
		matches:  make(map[string]*riot.Match, capacity),
		mu:       sync.Mutex{},
	}
}

func (m *MatchStore) GetMatch(matchId string) (*riot.Match, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	match := m.matches[matchId]
	if match != nil {
		return match, nil
	} else {
		return nil, errors.New("match not found")
	}
}

func (m *MatchStore) SetMatch(matchId string, match *riot.Match) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.matches[matchId] = match
	if len(m.matches) > int(m.Capacity) {
		var oldestAccessTime int64 = time.Now().Unix()
		var oldestAccessGame string = ""
		for matchId := range m.matches {
			if m.lastAccess[matchId] < oldestAccessTime {
				oldestAccessTime = m.lastAccess[matchId]
				oldestAccessGame = matchId
			}
		}
		delete(m.matches, oldestAccessGame)
	}
	return nil
}

func (m *MatchStore) DeleteMatch(matchId string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.matches, matchId)
}
