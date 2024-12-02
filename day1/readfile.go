package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getFileText(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(content)
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
