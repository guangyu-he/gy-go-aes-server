package bundesliga

import (
	"fmt"
	"testing"
)

func TestLastFiveGames(t *testing.T) {
	result, _ := LastFiveGames(5) // Bayern Munich
	if len(result) != 5 {
		t.Errorf("Expected 5 matches, got %d", len(result))
	}
	for _, match := range result {
		fmt.Println(match)
	}
}
