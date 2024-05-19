package compute

import (
	"testing"
)

func TestMedian(t *testing.T) {
	input := []float64{56, 47, 65}
	output := Average(input)
	expected := float64(56)

	if output != expected {
		t.Errorf("Test failed. expected: %v, got: %v", expected, output)
	}
}
