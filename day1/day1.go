package day1

import (
	"adventOfCode2024/utils"
	"fmt"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"
)

func Part1() int {
	defer utils.TimeTrack(time.Now())
	abs, _ := filepath.Abs("day1/input.txt")
	nums1, nums2, err := parseFileText(utils.GetFileText(abs))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	slices.Sort(nums1)
	slices.Sort(nums2)
	sum := 0
	for i := 0; i < len(nums1); i++ {
		sum += utils.IntAbs(nums1[i] - nums2[i])
	}
	return sum
}

func Part2() int {
	defer utils.TimeTrack(time.Now())
	abs, _ := filepath.Abs("day1/input.txt")
	nums1, nums2, err := parseFileText(utils.GetFileText(abs))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	nums2map := make(map[int]int)
	for _, num := range nums2 {
		num2count, ok := nums2map[num]
		if !ok {
			nums2map[num] = 0
		}
		nums2map[num] = num2count + 1
	}
	similarityScore := 0
	for _, num := range nums1 {
		similarityScore += nums2map[num] * num
	}
	return similarityScore
}

func parseFileText(text string) ([]int, []int, error) {
	var nums1, nums2 []int
	for _, line := range strings.Split(text, "\n") {
		if line == "" {
			continue
		}
		nums := strings.Split(line, "   ")
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			return nil, nil, err
		}
		nums1 = append(nums1, num1)
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			return nil, nil, err
		}
		nums2 = append(nums2, num2)
	}

	return nums1, nums2, nil
}
