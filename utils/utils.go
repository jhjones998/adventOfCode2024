package utils

import (
	"container/list"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
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
	return strings.Trim(string(content), " \t\r\n")
}

func IntAbs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func IntMaxMin(num1, num2 int) (int, int) {
	if num1 > num2 {
		return num1, num2
	}
	return num2, num1
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

func IntInRange(num, low, high int) bool {
	return num >= low && num < high
}

func PrintList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}

func DivMod(dividend, divisor int) (quotient, remainder int, err error) {
	if divisor == 0 {
		return 0, 0, errors.New("divisor is zero")
	}
	quotient = dividend / divisor
	remainder = dividend - quotient*divisor
	return quotient, remainder, nil
}

func ExtendedGcd(a, b int) (oldR, oldS, oldT, s, t int) {
	oldR, r := a, b
	oldS, s = 1, 0
	oldT, t = 0, 1

	for r != 0 {
		quotient := oldR / r
		oldR, r = r, oldR-quotient*r
		oldS, s = s, oldS-quotient*s
		oldT, t = t, oldT-quotient*t
	}
	return oldR, oldS, oldT, s, t
}

func ExtendedGcdUint64(a, b uint64) (oldR, oldS, oldT, s, t uint64) {
	oldR, r := a, b
	oldS, s = 1, 0
	oldT, t = 0, 1

	for r != 0 {
		quotient := oldR / r
		oldR, r = r, oldR-quotient*r
		oldS, s = s, oldS-quotient*s
		oldT, t = t, oldT-quotient*t
	}
	return oldR, oldS, oldT, s, t
}

func PositiveMod(a, b int) int {
	return ((a % b) + b) % b
}
