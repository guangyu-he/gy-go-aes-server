package bundesliga

import (
	"encoding/json"
	"fmt"
	"log"
)

func getTeamName(nextMatch NextMatch, teamID int) string {
	if nextMatch.HomeTeam.ID == teamID {
		return nextMatch.HomeTeam.Name
	}
	return nextMatch.AwayTeam.Name
}

func NextGame(teamID int) (NextMatch, error) {
	url := fmt.Sprintf("https://api.football-data.org/v4/teams/%d/matches?limit=1&status=SCHEDULED&season=2024", teamID)
	body, err := RequestGet(url)
	if err != nil {
		return NextMatch{}, err
	}

	var responseData ResponseNextMatch
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		log.Println(err)
		return NextMatch{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var nextMatch NextMatch
	if len(responseData.Matches) == 0 {
		return NextMatch{}, fmt.Errorf("no upcoming matches found")
	}
	nextMatch = responseData.Matches[0]
	nextMatch.TeamName = getTeamName(nextMatch, teamID)
	nextMatch.HomeTeam.Power = CalculatePower(nextMatch.HomeTeam.ID)
	nextMatch.AwayTeam.Power = CalculatePower(nextMatch.AwayTeam.ID)

	// TODO! update this algorithm later!
	nextMatch.Prediction.HomeTeam = nextMatch.HomeTeam.Power
	nextMatch.Prediction.AwayTeam = nextMatch.AwayTeam.Power

	return nextMatch, nil
}
