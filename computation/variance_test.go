package compute

import (
	"testing"
)

func TestVariance(t *testing.T) {
	input := []float64{47, 56, 65, 78}
	output := Variance(input)
	expected := float64(131.25)

	if output != expected {
		t.Errorf("Standard test failed:\n Got: %v, Expected: %v", output, expected)
	}
}
