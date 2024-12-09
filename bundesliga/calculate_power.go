package bundesliga

func CalculatePower(teamID int) int {

	// TODO! update this algorithm later!
	_, err := LastFiveGames(teamID)
	if err != nil {
		return 0
	}
	var power int
	//for _, match := range lastFiveGames {
	//	if match.Winner.ID == teamID {
	//		power++
	//	}
	//}
	return power
}
