package bundesliga

func CalculatePower(teamID int) int {

	// TODO! update this algorithm later!
	lastfivegames := LastFiveGames(teamID)
	var power int
	for _, match := range lastfivegames {
		if match.Winner.ID == teamID {
			power++
		}
	}
	return power
}
