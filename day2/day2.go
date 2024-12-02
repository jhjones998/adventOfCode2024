package day2

import (
	"adventOfCode2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func D2Part1() {
	reportNums := parseFileTextP1(utils.GetFileText("day2/inputp1.txt"))
	safeCount := 0
	for _, report := range reportNums {
		reportSafe := validateWithNoDampener(report)
		if reportSafe {
			safeCount++
		}
	}
	fmt.Println(safeCount)
}

func D2Part2() {
	reportNums := parseFileTextP1(utils.GetFileText("day2/inputp1.txt"))
	safeCount := 0
	for _, report := range reportNums {
		for i := 0; i < len(report); i++ {
			var tmpReport []int
			reportWithIRemoved := append(append(tmpReport, report[:i]...), report[i+1:]...)
			reportSafe := validateWithNoDampener(reportWithIRemoved)
			if reportSafe {
				safeCount++
				break
			}
		}
	}
	fmt.Println(safeCount)
}

func parseFileTextP1(input string) [][]int {
	reports := strings.Split(input, "\n")
	var reportNums [][]int
	for _, report := range reports {
		if strings.Trim(report, " ") == "" {
			continue
		}
		strNums := strings.Split(report, " ")
		nums := make([]int, len(strNums))
		for j, strNum := range strNums {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			nums[j] = num
		}
		reportNums = append(reportNums, nums)
	}
	return reportNums
}

func validateWithNoDampener(report []int) bool {
	firstSign := utils.IntSign(report[0] - report[1])
	for i := 0; i < len(report)-1; i++ {
		diff := report[i] - report[i+1]
		absDiff := utils.IntAbs(diff)
		if absDiff < 1 || absDiff > 3 || firstSign != utils.IntSign(diff) {
			return false
		}
	}
	return true
}
