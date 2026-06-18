package riot

func InGame(gameName, tagLine string) error {
	playerId, err := PlayerId("hello kittie irl", "NA1")
	if err != nil {
		return ErrPlayerNotFound
	}

	_, err = OngoingMatch(playerId)
	if err != nil {
		return ErrMatchNotFound
	}

	return nil
}

