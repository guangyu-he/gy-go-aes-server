package bundesliga

import (
	"encoding/json"
	"fmt"
)

func LastFiveGames() []Match {

	body := Query("https://api.football-data.org/v4/teams/5/matches?limit=5&status=FINISHED&season=2024")

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	ListOfMatches := []Match{}
	for _, match := range data["matches"].([]interface{}) {
		var matchData Match
		matchData.ID = int(match.(map[string]interface{})["id"].(float64))
		matchData.UtcDate = match.(map[string]interface{})["utcDate"].(string)
		matchData.Status = match.(map[string]interface{})["status"].(string)
		matchData.MatchDay = int(match.(map[string]interface{})["matchday"].(float64))

		homeTeam := match.(map[string]interface{})["homeTeam"].(map[string]interface{})["name"].(string)
		homeTeamId := int(match.(map[string]interface{})["homeTeam"].(map[string]interface{})["id"].(float64))
		matchData.HomeTeam.ID = homeTeamId
		matchData.HomeTeam.Name = homeTeam

		awayTeam := match.(map[string]interface{})["awayTeam"].(map[string]interface{})["name"].(string)
		awayTeamId := int(match.(map[string]interface{})["awayTeam"].(map[string]interface{})["id"].(float64))
		matchData.AwayTeam.ID = awayTeamId
		matchData.AwayTeam.Name = awayTeam

		score := match.(map[string]interface{})["score"].(map[string]interface{})["fullTime"].(map[string]interface{})
		matchData.Score.FullTime.HomeTeam = int(score["home"].(float64))
		matchData.Score.FullTime.AwayTeam = int(score["away"].(float64))

		ListOfMatches = append(ListOfMatches, matchData)
	}

	return ListOfMatches
}
