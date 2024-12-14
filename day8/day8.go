package day8

import (
	"adventOfCode2024/utils"
	"strings"
	"time"
)

type antennaType rune
type mapPos [2]int

func Part1() int {
	defer utils.TimeTrack(time.Now())
	antennas, numRows, numCols := getInputValues("day8/input.txt")
	antipodesSeen := make(map[mapPos]bool)
	for _, antennaPositions := range antennas {
		for i := 0; i < len(antennaPositions)-1; i++ {
			antennaPos := antennaPositions[i]
			for _, otherAntennaPos := range antennaPositions[i+1:] {
				antipodePos1 := mapPos{
					2*antennaPos[0] - otherAntennaPos[0],
					2*antennaPos[1] - otherAntennaPos[1],
				}
				if utils.IntInRange(antipodePos1[0], 0, numRows) &&
					utils.IntInRange(antipodePos1[1], 0, numCols) &&
					!antipodesSeen[antipodePos1] {
					antipodesSeen[antipodePos1] = true
				}
				antipodePos2 := mapPos{
					2*otherAntennaPos[0] - antennaPos[0],
					2*otherAntennaPos[1] - antennaPos[1],
				}
				if utils.IntInRange(antipodePos2[0], 0, numRows) &&
					utils.IntInRange(antipodePos2[1], 0, numCols) &&
					!antipodesSeen[antipodePos2] {
					antipodesSeen[antipodePos2] = true
				}
			}
		}
	}
	return len(antipodesSeen)
}

func Part2() int {
	defer utils.TimeTrack(time.Now())
	antennas, numRows, numCols := getInputValues("day8/input.txt")
	antinodesSeen := make(map[mapPos]bool)
	for _, antennaPositions := range antennas {
		for i := 0; i < len(antennaPositions)-1; i++ {
			antennaPos := antennaPositions[i]
			antinodesSeen[antennaPos] = true
			for _, otherAntennaPos := range antennaPositions[i+1:] {
				antinodesSeen[otherAntennaPos] = true
				dy := otherAntennaPos[1] - antennaPos[1]
				dx := otherAntennaPos[0] - antennaPos[0]
				hasIntSlope := dx != 0 && dy%dx == 0
				if dx == 0 {
					for y := 0; y < numCols; y++ {
						antinodesSeen[mapPos{antennaPos[0], y}] = true
					}
				} else {
					for x := 0; x < numRows; x++ {
						if hasIntSlope || (x-antennaPos[0])%dx == 0 {
							y := dy*(x-antennaPos[0])/dx + antennaPos[1]
							if utils.IntInRange(y, 0, numCols) {
								antinodesSeen[mapPos{x, y}] = true
							}
						}
					}
				}
			}
		}
	}
	return len(antinodesSeen)
}

func getInputValues(filename string) (antennas map[antennaType][]mapPos, numRows, numCols int) {
	input := strings.Trim(utils.GetFileText(filename), " \n")
	rows := strings.Split(input, "\n")
	numRows = len(rows)
	numCols = len(rows[0])
	antennas = make(map[antennaType][]mapPos)
	for i, row := range rows {
		for j, char := range row {
			if char != '.' {
				if _, ok := antennas[antennaType(char)]; !ok {
					antennas[antennaType(char)] = []mapPos{}
				}
				antennas[antennaType(char)] = append(antennas[antennaType(char)], mapPos{i, j})
			}
		}
	}
	return antennas, numRows, numCols
}
