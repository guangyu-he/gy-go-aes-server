package bundesliga

import (
	"encoding/json"
	"fmt"
	"log"
)

func LastFiveGames(teamID int) ([]Match, error) {

	url := fmt.Sprintf("https://api.football-data.org/v4/teams/%d/matches?limit=5&status=FINISHED&season=2024", teamID)
	body, err := RequestGet(url)
	if err != nil {
		return []Match{}, err
	}

	var responseData Response
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		log.Println(err)
		return []Match{}, err
	}

	var ListOfMatches []Match
	for _, match := range responseData.Matches {
		if match.Score.FullTime.HomeTeam > match.Score.FullTime.AwayTeam {
			match.HomeTeam.Win = true
			match.HomeTeam.GD = match.Score.FullTime.HomeTeam - match.Score.FullTime.AwayTeam
			match.HomeTeam.GS = match.Score.FullTime.HomeTeam
			match.HomeTeam.GA = match.Score.FullTime.AwayTeam

			match.AwayTeam.GD = match.Score.FullTime.AwayTeam - match.Score.FullTime.HomeTeam
			match.AwayTeam.GS = match.Score.FullTime.AwayTeam
			match.AwayTeam.GA = match.Score.FullTime.HomeTeam
		} else if match.Score.FullTime.HomeTeam < match.Score.FullTime.AwayTeam {
			match.AwayTeam.Win = true
			match.AwayTeam.GD = match.Score.FullTime.AwayTeam - match.Score.FullTime.HomeTeam
			match.AwayTeam.GS = match.Score.FullTime.AwayTeam
			match.AwayTeam.GA = match.Score.FullTime.HomeTeam

			match.HomeTeam.GD = match.Score.FullTime.HomeTeam - match.Score.FullTime.AwayTeam
			match.HomeTeam.GS = match.Score.FullTime.HomeTeam
			match.HomeTeam.GA = match.Score.FullTime.AwayTeam
		} else {
			match.HomeTeam.Draw = true
			match.AwayTeam.Draw = true
			match.HomeTeam.GD = 0
			match.HomeTeam.GS = match.Score.FullTime.HomeTeam
			match.HomeTeam.GA = match.Score.FullTime.AwayTeam

			match.AwayTeam.GD = 0
			match.AwayTeam.GS = match.Score.FullTime.AwayTeam
			match.AwayTeam.GA = match.Score.FullTime.HomeTeam
		}
		ListOfMatches = append(ListOfMatches, match)
	}

	return ListOfMatches, nil
}
