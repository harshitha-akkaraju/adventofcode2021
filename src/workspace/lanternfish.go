package workspace

/// Day 6 :: Lanternfish
// https://adventofcode.com/2021/day/6

func simulateLanternfish(fish []int, n int) int {
	temp := []int{}
	for i := 0; i < n; i++ {
		for j := 0; j < len(fish); j++ {
			updated := fish[j] - 1
			if updated < 0 {
				updated = 6
				temp = append(temp, 8)
			}
			fish[j] = updated
		}

		fish = append(fish, temp...)
		temp = []int{}
	}

	return len(fish)
}
