package compute

import (
	"testing"
)

func TestAverage(t *testing.T) {
	input := []float64{47, 65, 56, 78}
	output := Average(input)
	expected := float64(61.5)

	if output != expected {
		t.Errorf("Average test failed: Got: %v, Expected: %v", output, expected)
	}
}
