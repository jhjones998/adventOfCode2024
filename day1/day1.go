package day1

import (
	"fmt"
	"math"
	"path/filepath"
	"slices"
)

func D1Part1() {
	abs, _ := filepath.Abs("day1/input.txt")
	nums1, nums2, err := parseFileText(getFileText(abs))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	slices.Sort(nums1)
	slices.Sort(nums2)
	sum := 0
	for i := 0; i < len(nums1); i++ {
		sum += int(math.Abs(float64(nums1[i]) - float64(nums2[i])))
	}
	fmt.Println(sum)
}

func D1Part2() {
	abs, _ := filepath.Abs("day1/input.txt")
	nums1, nums2, err := parseFileText(getFileText(abs))
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
	fmt.Println(similarityScore)
}
