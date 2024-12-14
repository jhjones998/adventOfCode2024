package day3

import (
	"adventOfCode2024/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Part1() int {
	defer utils.TimeTrack(time.Now())
	input := utils.GetFileText("day3/input.txt")
	matches := getRegexpMatches(input, `mul\(\d{1,3},\d{1,3}\)`)
	mulSum := 0
	for _, match := range matches {
		mulSum += parseMul(match)
	}
	return mulSum
}

func Part2() int {
	defer utils.TimeTrack(time.Now())
	input := utils.GetFileText("day3/input.txt")
	matches := getRegexpMatches(input, `(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don\'t\(\))`)
	mulSum := 0
	do := true
	for _, match := range matches {
		if match == "do()" {
			do = true
			continue
		}
		if match == "don't()" {
			do = false
			continue
		}
		if strings.Contains(match, "mul") && do {
			mulSum += parseMul(match)
		}
	}
	return mulSum
}

func parseMul(s string) int {
	parts := strings.Split(s, ",")
	num1, err := strconv.Atoi(parts[0][4:])
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	num2, err := strconv.Atoi(parts[1][:len(parts[1])-1])
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return num1 * num2
}

func getRegexpMatches(s, pattern string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindAllString(s, -1)
}
