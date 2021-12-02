package workspace

import (
	"fmt"
	"testing"
)

func TestCountDepthPressureIncrease(t *testing.T) {
	t.Skip("meant to be run locally")

	scanner := NewScanner()
	depths := scanner.ToIntArray("<input-file>")
	count := countDepthPressureIncrease(depths)
	fmt.Printf("Part 1: %d \n", count)

	compressed := slidingWindow(depths)
	count = countDepthPressureIncrease(compressed)
	fmt.Printf("Part 2: %d \n", count)
}

func TestDive(t *testing.T) {
	t.Skip("meant to be run locally")

	scanner := NewScanner()
	directions := scanner.ToStringArray("<input-file>")

	horizontalPos, verticalPos := calculateHorizontalPositionAndDepth(directions)
	fmt.Printf(
		"Part 1 :: Horizontal Position: %d	Vertical Position: %d	Product: %d \n",
		horizontalPos,
		verticalPos,
		horizontalPos * verticalPos)

	horizontalPos, verticalPos = calculateHorizontalPositionAndDepthV2(directions)
	fmt.Printf(
		"Part 2 :: Horizontal Position: %d	Vertical Position: %d	Product: %d \n",
		horizontalPos,
		verticalPos,
		horizontalPos * verticalPos)
}