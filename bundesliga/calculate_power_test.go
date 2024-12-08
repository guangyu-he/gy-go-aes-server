package bundesliga

import "testing"

func TestCalculatePower(t *testing.T) {
	result := CalculatePower(5)
	if result < 0 {
		t.Errorf("Expected positive number, got %d", result)
	}
	return
}
