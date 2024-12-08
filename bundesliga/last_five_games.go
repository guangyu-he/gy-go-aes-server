package bundesliga

import (
	"encoding/json"
	"log"
)

func LastFiveGames() []Match {

	body := Query("https://api.football-data.org/v4/teams/5/matches?limit=5&status=FINISHED&season=2024")

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
