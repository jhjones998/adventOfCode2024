package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"time"
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

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func TimeTrack(start time.Time) {
	elapsed := time.Since(start)

	// Skip this function, and fetch the PC and file for its parent.
	pc, _, _, _ := runtime.Caller(1)

	// Retrieve a function object this functions parent.
	funcObj := runtime.FuncForPC(pc)

	// Regex to extract just the function name (and not the module path).
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")

	fmt.Println(fmt.Sprintf("%s took %s", name, elapsed))
}
