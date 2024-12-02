package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetFileText(filename string) string {
	abs, err := filepath.Abs(filename)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	content, err := os.ReadFile(abs)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(content)
}

func IntAbs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func IntSign(num int) int {
	if num < 0 {
		return -1
	}
	if num > 0 {
		return 1
	}
	return 0
}
