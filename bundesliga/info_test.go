package bundesliga

import (
	"fmt"
	"testing"
)

func TestInfo(t *testing.T) {
	result := MatchInfo("13")
	fmt.Println(result)
	return
}
