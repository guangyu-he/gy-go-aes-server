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
	t.Logf("Next game of %s: %s (H) vs %s (A)", result.TeamName, result.HomeTeam.Name, result.AwayTeam.Name)
}
