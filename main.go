package main

import (
	"adventOfCode2024/day1"
	"adventOfCode2024/day2"
	"adventOfCode2024/day3"
	"adventOfCode2024/day4"
	"adventOfCode2024/day5"
	"adventOfCode2024/day6"
	"adventOfCode2024/day7"
	"adventOfCode2024/day8"
	"adventOfCode2024/day9"
	"adventOfCode2024/utils"
	"fmt"
	"time"
)

func main() {
	defer utils.TimeTrack(time.Now())
	fmt.Println("Advent of Code 2024")
	fmt.Println()

	fmt.Println("Day 1 Part 1")
	fmt.Println("Answer:", day1.D1Part1())
	fmt.Println("Day 1 Part 2")
	fmt.Println("Answer:", day1.D1Part2())
	fmt.Println()

	fmt.Println("Day 2 Part 1")
	fmt.Println("Answer:", day2.D2Part1())
	fmt.Println("Day 2 Part 2")
	fmt.Println("Answer:", day2.D2Part2())
	fmt.Println()

	fmt.Println("Day 3 Part 1")
	fmt.Println("Answer:", day3.D3Part1())
	fmt.Println("Day 3 Part 2")
	fmt.Println("Answer:", day3.D3Part2())
	fmt.Println()

	fmt.Println("Day 4 Part 1")
	fmt.Println("Answer:", day4.D4Part1())
	fmt.Println("Day 4 Part 2")
	fmt.Println("Answer:", day4.D4Part2())
	fmt.Println()

	fmt.Println("Day 5 Part 1")
	fmt.Println("Answer:", day5.D5Part1())
	fmt.Println("Day 5 Part 2")
	fmt.Println("Answer:", day5.D5Part2())
	fmt.Println()

	fmt.Println("Day 6 Part 1")
	fmt.Println("Answer:", day6.D6Part1())
	fmt.Println("Day 6 Part 2")
	fmt.Println("Map Answer:", day6.D6Part2())
	fmt.Println("Graph Answer:", day6.D6Part2Graph())
	fmt.Println()

	fmt.Println("Day 7 Part 1")
	fmt.Println("Answer:", day7.D7Part1())
	fmt.Println("Day 7 Part 2")
	fmt.Println("Answer:", day7.D7Part2())
	fmt.Println()

	fmt.Println("Day 8 Part 1")
	fmt.Println("Answer:", day8.D8Part1())
	fmt.Println("Day 8 Part 2")
	fmt.Println("Answer:", day8.D8Part2())
	fmt.Println()

	fmt.Println("Day 9 Part 1")
	fmt.Println("Answer:", day9.D9Part1())
	fmt.Println("Day 9 Part 2")
	fmt.Println("Answer:", day9.D9Part2())
	fmt.Println()

	fmt.Println("Done")
}
