package bundesliga

import (
	"encoding/json"
	"fmt"
	"log"
)

func NextGame(teamID int) NextMatch {
	url := fmt.Sprintf("https://api.football-data.org/v4/teams/%d/matches?limit=1&status=SCHEDULED&season=2024", teamID)
	body := Query(url)

	var responseData ResponseNextMatch
	err := json.Unmarshal(body, &responseData)
	if err != nil {
		log.Fatal(err)
	}

	var nextMatch NextMatch
	nextMatch = responseData.Matches[0]
	nextMatch.HomeTeam.Power = CalculatePower(nextMatch.HomeTeam.ID)
	nextMatch.AwayTeam.Power = CalculatePower(nextMatch.AwayTeam.ID)

	// TODO! update this algorithm later!
	nextMatch.Prediction.HomeTeam = nextMatch.HomeTeam.Power
	nextMatch.Prediction.AwayTeam = nextMatch.AwayTeam.Power

	return nextMatch

}
