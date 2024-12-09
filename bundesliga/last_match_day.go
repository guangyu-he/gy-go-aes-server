package bundesliga

import (
	"encoding/json"
	"fmt"
	"log"
)

func LatestMatchDay() (string, error) {

	body, err := RequestGet("https://api.football-data.org/v4/competitions/BL1")

	var competition Competition
	err = json.Unmarshal(body, &competition)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		return "", err
	}

	return fmt.Sprintf("%d", competition.CurrentSeason.CurrentMatchday), nil
}
