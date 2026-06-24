package riot

import (
	"errors"
	"unicode/utf8"
)

//
// This file defines some simple validator functions for values that might be
// sent to/from the Riot API.
//

// Returns nil if and only if the given game name is valid. Otherwise, an
// error is returned that describes the problem. The error is appropriate for
// a plaintext error response.
func ValidateGameName(gameName string) error {

	runeCount := utf8.RuneCountInString(gameName)
	if runeCount > 16 {
		return errors.New("gameName cannot have more than 16 runes")
	}
	if runeCount < 3 {
		return errors.New("gameName cannot have fewer than 3 runes")
	}

	return nil
}

// Returns nil if and only if the given tag line is valid. Otherwise, an error
// is returned that describes the problem. The error is appropriate for a
// plaintext error response.
func ValidateTagLine(tagLine string) error {

	runeCount := utf8.RuneCountInString(tagLine)
	if runeCount > 5 {
		return errors.New("tagLine cannot have more than 5 runes")
	}
	if runeCount < 3 {
		return errors.New("tagLine cannot have fewer than 4 runes")
	}

	return nil
}
