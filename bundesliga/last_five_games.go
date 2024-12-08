package bundesliga

import (
	"encoding/json"
	"fmt"
	"log"
)

func LastFiveGames(teamID int) []Match {

	url := fmt.Sprintf("https://api.football-data.org/v4/teams/%d/matches?limit=5&status=FINISHED&season=2024", teamID)
	body := Query(url)

	var responseData Response
	err := json.Unmarshal(body, &responseData)
	if err != nil {
		log.Fatal(err)
	}

	var ListOfMatches []Match
	for _, match := range responseData.Matches {

		if match.Score.FullTime.HomeTeam > match.Score.FullTime.AwayTeam {
			match.Winner.Name = match.HomeTeam.Name
			match.Winner.ID = match.HomeTeam.ID
		} else if match.Score.FullTime.HomeTeam < match.Score.FullTime.AwayTeam {
			match.Winner.Name = match.AwayTeam.Name
			match.Winner.ID = match.AwayTeam.ID
		} else {
			match.Draw = true
		}

		ListOfMatches = append(ListOfMatches, match)
	}

	return ListOfMatches
}
