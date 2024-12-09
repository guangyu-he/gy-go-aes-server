package bundesliga

import "testing"

func TestNextGame(t *testing.T) {
	result, err := NextGame(5)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if result.TeamName != "FC Bayern München" {
		t.Errorf("Expected: FC Bayern München, Got: %s", result.TeamName)
	}
	t.Logf("Next game of %s: %s (H) [%d] vs %s (A) [%d]", result.TeamName, result.HomeTeam.Name, result.HomeTeam.Power, result.AwayTeam.Name, result.AwayTeam.Power)
	t.Logf("Prediction: %s - %d vs %s - %d", result.HomeTeam.Name, result.Prediction.HomeTeam, result.AwayTeam.Name, result.Prediction.AwayTeam)
}
