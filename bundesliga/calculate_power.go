package bundesliga

func CalculatePower(teamID int) int {

	// TODO! update this algorithm later!
	lastFiveGames := LastFiveGames(teamID)
	var power int
	for _, match := range lastFiveGames {
		if match.Winner.ID == teamID {
			power++
		}
	}
	return power
}
