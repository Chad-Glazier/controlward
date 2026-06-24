package memstore

import (
	"fmt"
	"testing"

	"github.com/Chad-Glazier/controlward/riot"
)

func TestMatchStore(t *testing.T) {
	s := NewMatchStore(100)

	for i := range s.Capacity {
		s.SetMatch(fmt.Sprintf("%d", i), &riot.Match{})
	}

	if len(s.matches) != int(s.Capacity) {
		t.Fatalf("expected %d matches, got %d", s.Capacity, len(s.matches))
	}

	for i := range s.Capacity {
		if match, _ := s.GetMatch(fmt.Sprintf("%d", i)); match == nil {
			t.Fatalf("expected match %d but was not found", i)
		}
	}

	for i := range s.Capacity * 3 {
		s.SetMatch(fmt.Sprintf("%d", i), &riot.Match{})
		if len(s.matches) != int(s.Capacity) {
			t.Fatalf("expected %d matches, got %d", s.Capacity, len(s.matches))
		}
	}
}
