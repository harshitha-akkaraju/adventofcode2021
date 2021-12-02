package workspace

import (
	"log"
	"strconv"
	"strings"
)

/// Day 2: Dive
// https://adventofcode.com/2021/day/2

func calculateHorizontalPositionAndDepth(directions []string) (int, int) {
	var horizontalPos, verticalPos int
	for _, line := range directions {
		parts := strings.Split(line, " ")
		direction := parts[0]
		distanceAsString := parts[1]
		distance, err := strconv.Atoi(distanceAsString)
		if err != nil {
			log.Fatal(err)
		}

		if direction == "forward" {
			horizontalPos += distance
		} else if direction == "up" {
			verticalPos -= distance
		} else if direction == "down" {
			verticalPos += distance
		}
	}
	return horizontalPos, verticalPos
}

func calculateHorizontalPositionAndDepthV2(directions []string) (int, int) {
	var horizontalPos, verticalPos, aim int
	for _, line := range directions {
		parts := strings.Split(line, " ")
		direction := parts[0]
		distanceAsString := parts[1]
		distance, err := strconv.Atoi(distanceAsString)
		if err != nil {
			log.Fatal(err)
		}

		if direction == "forward" {
			horizontalPos += distance
			verticalPos += aim * distance
		} else if direction == "up" {
			aim -= distance
		} else if direction == "down" {
			aim += distance
		}
	}
	return horizontalPos, verticalPos
}
