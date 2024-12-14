package day4

import (
	"adventOfCode2024/utils"
	"strings"
	"time"
)

func Part1() int {
	defer utils.TimeTrack(time.Now())
	xmasRunes := []rune("XMAS")
	wordSearch := getWordSearch("day4/input.txt")
	xmasCount := 0
	for i := 0; i < len(wordSearch); i++ {
		for j := 0; j < len(wordSearch[i]); j++ {
			// check north
			xmasCount += utils.BoolToInt(checkForWord(&wordSearch, xmasRunes, i, -1, j, 0))
			// check northeast
			xmasCount += utils.BoolToInt(checkForWord(&wordSearch, xmasRunes, i, -1, j, 1))
			// check east
			xmasCount += utils.BoolToInt(checkForWord(&wordSearch, xmasRunes, i, 0, j, 1))
			// check southeast
			xmasCount += utils.BoolToInt(checkForWord(&wordSearch, xmasRunes, i, 1, j, 1))
			// check south
			xmasCount += utils.BoolToInt(checkForWord(&wordSearch, xmasRunes, i, 1, j, 0))
			// check southwest
			xmasCount += utils.BoolToInt(checkForWord(&wordSearch, xmasRunes, i, 1, j, -1))
			// check west
			xmasCount += utils.BoolToInt(checkForWord(&wordSearch, xmasRunes, i, 0, j, -1))
			// check northwest
			xmasCount += utils.BoolToInt(checkForWord(&wordSearch, xmasRunes, i, -1, j, -1))
		}
	}
	return xmasCount
}

func Part2() int {
	defer utils.TimeTrack(time.Now())
	masRunes := []rune("MAS")
	diagonalDistance := (len(masRunes) - 1) / 2
	wordSearch := getWordSearch("day4/input.txt")
	masxCount := 0
	for i := 0; i < len(wordSearch); i++ {
		for j := 0; j < len(wordSearch[i]); j++ {
			if wordSearch[i][j] != masRunes[1] {
				continue
			}
			if i >= diagonalDistance &&
				i < len(wordSearch)-diagonalDistance &&
				j >= diagonalDistance &&
				j < len(wordSearch[i])-diagonalDistance {
				upperleftRune := wordSearch[i-1][j-1]
				upperrightRune := wordSearch[i-1][j+1]
				lowerleftRune := wordSearch[i+1][j-1]
				lowerrightRune := wordSearch[i+1][j+1]
				rightDiagonalLegExists := (upperleftRune == masRunes[0] &&
					lowerrightRune == masRunes[2]) ||
					(upperleftRune == masRunes[2] &&
						lowerrightRune == masRunes[0])
				leftDiagonalLegExists := (upperrightRune == masRunes[0] &&
					lowerleftRune == masRunes[2]) ||
					(upperrightRune == masRunes[2] &&
						lowerleftRune == masRunes[0])
				if rightDiagonalLegExists && leftDiagonalLegExists {
					masxCount++
				}
			}
		}
	}
	return masxCount
}

func getWordSearch(input string) [][]rune {
	lines := strings.Split(utils.GetFileText(input), "\n")
	var runeArray [][]rune
	for _, line := range lines {
		if strings.Trim(line, " ") == "" {
			continue
		}
		runesInLine := []rune(line)
		runeArray = append(runeArray, runesInLine)
	}
	return runeArray
}

func checkForWord(wordSearch *[][]rune, runes []rune, rowIdx, rowInc, colIdx, colInc int) bool {
	runeLenMinusOne := len(runes) - 1
	if rowIdx+rowInc*runeLenMinusOne < 0 ||
		rowIdx+rowInc*runeLenMinusOne >= len(*wordSearch) ||
		colIdx+colInc*runeLenMinusOne < 0 ||
		colIdx+colInc*runeLenMinusOne >= len((*wordSearch)[rowIdx]) {
		return false
	}
	for i := 0; i <= runeLenMinusOne; i++ {
		if (*wordSearch)[rowIdx][colIdx] != runes[i] {
			return false
		}
		rowIdx += rowInc
		colIdx += colInc
	}
	return true
}
