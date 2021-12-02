package workspace

/// Day 1: Sonar Sweep
/// https://adventofcode.com/2021/day/1

func countDepthPressureIncrease(depths []int) int {
	count := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i - 1] {
			count++
		}
	}
	return count
}

func slidingWindow(depths []int) []int {
	result := []int{}
	for i := 0; i < len(depths) - 2; i++ {
		sum := depths[i] + depths[i + 1] + depths[i + 2]
		result = append(result, sum)
	}
	return result
}