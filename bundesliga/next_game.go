package bundesliga

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
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

	if nextMatch.HomeTeam.Power < 0 && nextMatch.AwayTeam.Power > 0 {
		// e.g. H: -5, A: 3
		// then turn home team to 0 and away team to 8
		nextMatch.HomeTeam.Power = 0
		nextMatch.AwayTeam.Power = nextMatch.AwayTeam.Power - nextMatch.HomeTeam.Power
	} else if nextMatch.HomeTeam.Power > 0 && nextMatch.AwayTeam.Power < 0 {
		// e.g. H: 5, A: -3
		// then turn home team to 8 and away team to 0
		nextMatch.AwayTeam.Power = 0
		nextMatch.HomeTeam.Power = nextMatch.HomeTeam.Power - nextMatch.AwayTeam.Power
	} else if nextMatch.HomeTeam.Power < 0 && nextMatch.AwayTeam.Power < 0 {
		if nextMatch.HomeTeam.Power < nextMatch.AwayTeam.Power {
			// e.g. H: -5, A: -3
			// this means home team is weaker than away team
			// then turn home team to 3 and away team to 5
			tempPower := nextMatch.HomeTeam.Power
			nextMatch.HomeTeam.Power = 0 - nextMatch.AwayTeam.Power
			nextMatch.AwayTeam.Power = 0 - tempPower
		} else {
			// e.g. H: -3, A: -5
			// this means away team is weaker than home team
			// then turn home team to 5 and away team to 3
			// or H: -3, A: -3
			// then turn home team to 3 and away team to 3
			tempPower := nextMatch.AwayTeam.Power
			nextMatch.AwayTeam.Power = 0 - nextMatch.HomeTeam.Power
			nextMatch.HomeTeam.Power = 0 - tempPower
		}
	} else if nextMatch.HomeTeam.Power == 0 && nextMatch.AwayTeam.Power == 0 {
		// e.g. H: 0, A: 0
		// this means both teams are equal
		// then turn home team to 3 and away team to 3
		nextMatch.HomeTeam.Power = 3
		nextMatch.AwayTeam.Power = 3
	}

	// give random int number from 0 to power
	nextMatch.Prediction.HomeTeam = rand.IntN(nextMatch.HomeTeam.Power + 1)
	nextMatch.Prediction.AwayTeam = rand.IntN(nextMatch.AwayTeam.Power + 1)

	return nextMatch, nil
}
