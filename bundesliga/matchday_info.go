package bundesliga

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

func MatchDayInfo(matchDay string) (string, error) {
	var url string
	var err error
	isLatestMatchday := false

	if matchDay == "latest" {
		matchDay, err = LatestMatchDay()
		if err != nil {
			log.Println("Error getting latest matchDay:", err)
			return "", err
		}
		isLatestMatchday = true
	}

	url = fmt.Sprintf("https://api.football-data.org/v4/competitions/BL1/matches?season=2024&matchday=%s", matchDay)
	body, err := RequestGet(url)

	//var data map[string]interface{}
	//if err := json.Unmarshal(body, &data); err != nil {
	//	fmt.Println("Error decoding JSON:", err)
	//	return
	//}

	var responseData Response
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return "", err
	}

	output := ""
	if isLatestMatchday {
		output += fmt.Sprintf("Matchday %s (latest) - Results:\n", matchDay)
	} else {
		output += fmt.Sprintf("Matchday %s - Results:\n", matchDay)
	}

	cet, err := CETFmt()
	if err != nil {
		log.Println("Error loading location:", err)
		return "", err
	}

	for _, match := range responseData.Matches {

		if strings.Contains(match.HomeTeam.Name, "FC Bayern") || strings.Contains(match.AwayTeam.Name, "FC Bayern") {
			output += "############################################\n"
		}

		output += fmt.Sprintf("主队: %s\n", match.HomeTeam.Name)
		output += fmt.Sprintf("客队: %s\n", match.AwayTeam.Name)

		// parse utctime string to time.Time
		utcTime, err := time.Parse(time.RFC3339, match.UtcDate)
		if err != nil {
			log.Printf("Error parsing time: %v", err)
			return "", err
		}
		// convert to CET time
		cetTime := utcTime.In(cet)
		output += fmt.Sprintf("比赛时间: %s\n", cetTime)

		output += fmt.Sprintf("比赛状态: %s\n", match.Status)
		output += fmt.Sprintf("比分: %d - %d\n", match.Score.FullTime.HomeTeam, match.Score.FullTime.AwayTeam)

		if strings.Contains(match.HomeTeam.Name, "FC Bayern") || strings.Contains(match.AwayTeam.Name, "FC Bayern") {
			output += "############################################\n"
		}

		output += "\n"
	}

	return output, nil
}

func CETFmt() (*time.Location, error) {
	cet, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		log.Println("Error loading location:", err)
		return nil, err
	} else {
		return cet, nil
	}
}
