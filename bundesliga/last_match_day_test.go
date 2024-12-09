package bundesliga

import "testing"

func TestLatestMatchDay(t *testing.T) {
	result, err := LatestMatchDay()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result == "" {
		t.Errorf("Expected a matchday, got empty string")
	}

	t.Logf("Latest matchday: %s", result)
}
