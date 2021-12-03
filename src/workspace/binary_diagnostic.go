package workspace

import (
	"log"
	"strconv"
)

/// Day 3: Binary Diagnostic
// https://adventofcode.com/2021/day/3

func newZeroArray(size int) []int {
	zeros := []int{}
	for i := 0; i < size; i++ {
		zeros = append(zeros, 0)
	}
	return zeros
}

func constructBitCountMap(report []string) [][]int {
	sample := report[0]
	bitCountMap := [][]int{
		newZeroArray(len(sample)), // 0 counts
		newZeroArray(len(sample)), // 1 counts
	}

	for _, line := range report {
		for i := 0; i < len(line); i++ {
			val := string(line[i])
			if val == "0" {
				bitCountMap[0][i]++
			} else {
				bitCountMap[1][i]++
			}
		}
	}

	return bitCountMap
}

// most common bit
func getGammaRate(report []string) int {
	counts := constructBitCountMap(report)
	zeroCounts := counts[0]
	oneCounts := counts[1]
	binary := ""

	length := len(counts[0])
	for i := 0; i < length; i++ {
		if zeroCounts[i] > oneCounts[i] {
			binary += "0"
		} else {
			binary += "1"
		}
	}

	gammaRate, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Fatal(err.Error())
	}

	return int(gammaRate)
}

// least common bit
func getEpsilonRate(report []string) int {
	counts := constructBitCountMap(report)
	zeroCounts := counts[0]
	oneCounts := counts[1]
	binary := ""

	length := len(counts[0])
	for i := 0; i < length; i++ {
		if zeroCounts[i] < oneCounts[i] {
			binary += "0"
		} else {
			binary += "1"
		}
	}

	epsilonRate, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Fatal(err.Error())
	}

	return int(epsilonRate)
}

func getPowerConsumption(report []string) int {
	gammaRate := getGammaRate(report)
	epsilonRate := getEpsilonRate(report)
	return gammaRate * epsilonRate
}

func getLifeSupportRate(report []string) int {
	oxygenGeneratorRating := getOxygenGeneratorRating(report)
	co2ScrubberRating := getCO2ScrubberRating(report)
	return oxygenGeneratorRating * co2ScrubberRating
}

func getOxygenGeneratorRating(report []string) int {
	counts := constructBitCountMap(report)
	length := len(counts[0])

	for pos := 0; pos < length; pos++ {
		if len(report) == 1 {
			break
		}

		maxBit := ""
		if counts[0][pos] > counts[1][pos] {
			maxBit = "0"
		} else {
			maxBit = "1"
		}

		report = filterLinesWithBitAtPos(report, maxBit, pos)
		counts = constructBitCountMap(report)
	}

	oxygenGeneratorRating, err := strconv.ParseInt(report[0], 2, 64)
	if err != nil {
		log.Fatal(err.Error())
	}

	return int(oxygenGeneratorRating)
}

func getCO2ScrubberRating(report []string) int {
	counts := constructBitCountMap(report)
	length := len(counts[0])

	for pos := 0; pos < length; pos++ {
		if len(report) == 1 {
			break
		}

		minBit := ""
		if counts[0][pos] <= counts[1][pos] {
			minBit = "0"
		} else {
			minBit = "1"
		}

		report = filterLinesWithBitAtPos(report, minBit, pos)
		counts = constructBitCountMap(report)
	}

	co2ScrubberRating, err := strconv.ParseInt(report[0], 2, 64)
	if err != nil {
		log.Fatal(err.Error())
	}

	return int(co2ScrubberRating)
}

func filterLinesWithBitAtPos(report []string, bit string, pos int) []string {
	filtered := []string{}
	for _, line := range report {
		val := string(line[pos])
		if val == bit {
			filtered = append(filtered, line)
		}
	}
	return filtered
}