package bundesliga

import (
	"fmt"
	"testing"
)

func TestInfo(t *testing.T) {
	result, err := MatchDayInfo("13")
	if err != nil {
		t.Error("Error getting matchday info:", err)
	}
	fmt.Println(result)
}
