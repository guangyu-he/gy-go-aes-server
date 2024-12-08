package bundesliga

import (
	"fmt"
	"testing"
)

func TestLastFiveGames(t *testing.T) {
	result := LastFiveGames()
	if len(result) != 5 {
		t.Errorf("Expected 5 matches, got %d", len(result))
	}
	fmt.Println(result)
	return
}
