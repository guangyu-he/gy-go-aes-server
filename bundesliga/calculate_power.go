package bundesliga

import "math"

func CalculatePower(teamID int) int {

	lastFiveGames, err := LastFiveGames(teamID)
	if err != nil {
		return 0
	}
	var power int
	for _, match := range lastFiveGames {
		if match.HomeTeam.ID == teamID {
			if match.HomeTeam.Win {
				power += 2
			} else if match.HomeTeam.Draw {
				power += 0
			} else {
				power -= 2
			}
			power += match.HomeTeam.GD * 2
			power += match.HomeTeam.GS
			power -= match.HomeTeam.GA

		} else if match.AwayTeam.ID == teamID {
			if match.AwayTeam.Win {
				power += 2
			} else if match.AwayTeam.Draw {
				power += 0
			} else {
				power -= 2
			}
			power += match.AwayTeam.GD * 2
			power += match.AwayTeam.GS
			power -= match.AwayTeam.GA
		}
	}
	power = int(math.Round(float64(power / 5)))

	return power
}
