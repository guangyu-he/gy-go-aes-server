package bundesliga

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Match struct {
	ID       int `json:"id"`
	HomeTeam struct {
		Name string `json:"name"`
	} `json:"homeTeam"`
	AwayTeam struct {
		Name string `json:"name"`
	} `json:"awayTeam"`
	UtcDate string `json:"utcDate"`
	Status  string `json:"status"`
	Score   struct {
		FullTime struct {
			HomeTeam int `json:"home"`
			AwayTeam int `json:"away"`
		} `json:"fullTime"`
	} `json:"score"`
}

type Response struct {
	Matches []Match `json:"matches"`
}

func MatchInfo() string {
	apiKey := "adfac7e310f6495f99f1c38883718fd0"
	url := "https://api.football-data.org/v4/competitions/BL1/matches?season=2024&matchday=13"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("X-Auth-Token", apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("请求失败，状态码：%d\n", resp.StatusCode)
	}

	//var data map[string]interface{}
	//if err := json.Unmarshal(body, &data); err != nil {
	//	fmt.Println("Error decoding JSON:", err)
	//	return
	//}

	var responseData Response
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		log.Fatal(err)
	}

	output := fmt.Sprintf("Matchday 13 - Results:\n")

	for _, match := range responseData.Matches {
		output += fmt.Sprintf("比赛 ID: %d\n", match.ID)
		output += fmt.Sprintf("主队: %s\n", match.HomeTeam.Name)
		output += fmt.Sprintf("客队: %s\n", match.AwayTeam.Name)
		output += fmt.Sprintf("比赛时间(UTC): %s\n", match.UtcDate)
		output += fmt.Sprintf("比赛状态: %s\n", match.Status)
		output += fmt.Sprintf("比分: %d - %d\n\n", match.Score.FullTime.HomeTeam, match.Score.FullTime.AwayTeam)
	}

	return output
}
