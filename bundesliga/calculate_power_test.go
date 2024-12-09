package bundesliga

import "testing"

func TestCalculatePower(t *testing.T) {
	result := CalculatePower(5)
	t.Logf("Power of Bayern Munich: %d", result)
}
