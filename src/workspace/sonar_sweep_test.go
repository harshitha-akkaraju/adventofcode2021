package workspace

import (
	"fmt"
	"testing"
)

func TestCountDepthPressureIncrease(t *testing.T) {
	t.Skip("Meant to be run locally")

	scanner := NewScanner()
	depths := scanner.ToIntArray("<input-file>")
	count := countDepthPressureIncrease(depths)
	fmt.Printf("Part 1: %d \n", count)

	compressed := slidingWindow(depths)
	count = countDepthPressureIncrease(compressed)
	fmt.Printf("Part 2: %d \n", count)
}