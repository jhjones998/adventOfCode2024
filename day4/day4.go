package day4

import (
	"adventOfCode2024/utils"
	"fmt"
	"strings"
)

/* Not sure if there's a better solution to part 2 that looks like part 1. I realized we
were just looking for and testing A's in part 2, which is a lot simpler than the XMAS
search in part 1. Look into it */

func D4Part1() {
	xmasRunes := []rune("XMAS")
	wordSearch := getWordSearch("day4/inputp1.txt")
	visitedCells := map[[2]int]bool{}
	xmasCount := 0
	for i := 0; i < len(wordSearch); i++ {
		for j := 0; j < len(wordSearch[i]); j++ {
			if visitedCells[[2]int{i, j}] {
				continue
			}
			if wordSearch[i][j] != xmasRunes[0] {
				continue
			}
			// check northerly directions
			if i >= 3 {
				// check north
				if wordSearch[i-1][j] == xmasRunes[1] &&
					wordSearch[i-2][j] == xmasRunes[2] &&
					wordSearch[i-3][j] == xmasRunes[3] {
					visitedCells[[2]int{i - 1, j}] = true
					visitedCells[[2]int{i - 2, j}] = true
					visitedCells[[2]int{i - 3, j}] = true
					xmasCount++
				}
				// check northwest
				if j >= 3 {
					if wordSearch[i-1][j-1] == xmasRunes[1] &&
						wordSearch[i-2][j-2] == xmasRunes[2] &&
						wordSearch[i-3][j-3] == xmasRunes[3] {
						visitedCells[[2]int{i - 1, j - 1}] = true
						visitedCells[[2]int{i - 2, j - 2}] = true
						visitedCells[[2]int{i - 3, j - 3}] = true
						xmasCount++
					}
				}
				//check northeast
				if j < len(wordSearch[i])-3 {
					if wordSearch[i-1][j+1] == xmasRunes[1] &&
						wordSearch[i-2][j+2] == xmasRunes[2] &&
						wordSearch[i-3][j+3] == xmasRunes[3] {
						visitedCells[[2]int{i - 1, j + 1}] = true
						visitedCells[[2]int{i - 2, j + 2}] = true
						visitedCells[[2]int{i - 3, j + 3}] = true
						xmasCount++
					}
				}
			}
			// check southerly directions
			if i < len(wordSearch)-3 {
				// check south
				if wordSearch[i+1][j] == xmasRunes[1] &&
					wordSearch[i+2][j] == xmasRunes[2] &&
					wordSearch[i+3][j] == xmasRunes[3] {
					visitedCells[[2]int{i + 1, j}] = true
					visitedCells[[2]int{i + 2, j}] = true
					visitedCells[[2]int{i + 3, j}] = true
					xmasCount++
				}

				// check southwest
				if j >= 3 {
					if wordSearch[i+1][j-1] == xmasRunes[1] &&
						wordSearch[i+2][j-2] == xmasRunes[2] &&
						wordSearch[i+3][j-3] == xmasRunes[3] {
						visitedCells[[2]int{i + 1, j - 1}] = true
						visitedCells[[2]int{i + 2, j - 2}] = true
						visitedCells[[2]int{i + 3, j - 3}] = true
						xmasCount++
					}
				}

				// check southeast
				if j < len(wordSearch[i])-3 {
					if wordSearch[i+1][j+1] == xmasRunes[1] &&
						wordSearch[i+2][j+2] == xmasRunes[2] &&
						wordSearch[i+3][j+3] == xmasRunes[3] {
						visitedCells[[2]int{i + 1, j + 1}] = true
						visitedCells[[2]int{i + 2, j + 2}] = true
						visitedCells[[2]int{i + 3, j + 3}] = true
						xmasCount++
					}
				}
			}

			// check east
			if j < len(wordSearch[i])-3 {
				if wordSearch[i][j+1] == xmasRunes[1] &&
					wordSearch[i][j+2] == xmasRunes[2] &&
					wordSearch[i][j+3] == xmasRunes[3] {
					visitedCells[[2]int{i, j + 1}] = true
					visitedCells[[2]int{i, j + 2}] = true
					visitedCells[[2]int{i, j + 3}] = true
					xmasCount++
				}
			}

			//check west
			if j >= 3 {
				if wordSearch[i][j-1] == xmasRunes[1] &&
					wordSearch[i][j-2] == xmasRunes[2] &&
					wordSearch[i][j-3] == xmasRunes[3] {
					visitedCells[[2]int{i, j - 1}] = true
					visitedCells[[2]int{i, j - 2}] = true
					visitedCells[[2]int{i, j - 3}] = true
					xmasCount++
				}
			}
		}
	}
	fmt.Println(xmasCount)
}

func D4Part2() {
	masRunes := []rune("MAS")
	wordSearch := getWordSearch("day4/inputp1.txt")
	masxCount := 0
	for i := 0; i < len(wordSearch); i++ {
		for j := 0; j < len(wordSearch[i]); j++ {
			if wordSearch[i][j] != masRunes[1] {
				continue
			}
			if i >= 1 && i < len(wordSearch)-1 && j >= 1 && j < len(wordSearch[i])-1 {
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
	fmt.Println(masxCount)
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
