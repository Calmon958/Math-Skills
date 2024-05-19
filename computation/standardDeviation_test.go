package compute

import (
	"testing"
)

func TestStandardDeviation(t *testing.T) {
	input := []float64{47, 56, 78, 65}
	output := StandardDeviation(input)
	expected := float64(11.4564392373896)

	if output != expected {
		t.Errorf("Variance test failed:\n Got: %v, Expected: %v", output, expected)
	}
}
