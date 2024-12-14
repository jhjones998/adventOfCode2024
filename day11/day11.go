package day11

import (
	"adventOfCode2024/utils"
	"container/list"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Part1() int {
	defer utils.TimeTrack(time.Now())
	stones := getInputValuesP1("day11/input.txt")
	for i := 0; i < 25; i++ {
		stones = blinkList(stones)
	}
	return stones.Len()
}

func Part2() int {
	defer utils.TimeTrack(time.Now())
	stones := getInputValuesP2("day11/input.txt")
	preCalcMap := make(map[int][2]int)
	for i := 0; i < 75; i++ {
		stones = blinkMap(stones, &preCalcMap)
	}
	stoneCount := 0
	for val := range stones {
		stoneCount += stones[val]
	}
	return stoneCount
}

func getInputValuesP1(filename string) (stones *list.List) {
	input := strings.Split(utils.GetFileText(filename), " ")
	stones = list.New()
	for _, i := range input {
		intI, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		stones.PushBack(intI)
	}
	return stones
}

func getInputValuesP2(filename string) map[int]int {
	input := strings.Split(utils.GetFileText(filename), " ")
	stones := make(map[int]int)
	for _, i := range input {
		intI, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		stones[intI]++
	}
	return stones
}

func blinkList(stones *list.List) *list.List {
	for e := stones.Front(); e != nil; e = e.Next() {
		eVal := e.Value.(int)
		numDigits := numberOfDigits(eVal)
		if e.Value.(int) == 0 {
			e.Value = 1
		} else if numDigits%2 == 0 {
			leftHalf, rightHalf := splitDigitWithEvenNumDigits(eVal, numDigits)
			stones.InsertBefore(leftHalf, e)
			afterElem := stones.InsertAfter(rightHalf, e)
			stones.Remove(e)
			e = afterElem
		} else {
			e.Value = eVal * 2024
		}
	}
	return stones
}

func blinkMap(stones map[int]int, preCalcMap *map[int][2]int) map[int]int {
	newStones := make(map[int]int)
	for val := range stones {
		numDigits := numberOfDigits(val)
		if val == 0 {
			newVal := 1
			newStones[newVal] += stones[val]
		} else if numDigits%2 == 0 {
			leftHalf, rightHalf := splitDigitWithEvenNumDigitsWithCache(val, numDigits, preCalcMap)
			newStones[leftHalf] += stones[val]
			newStones[rightHalf] += stones[val]
		} else {
			newVal := val * 2024
			newStones[newVal] += stones[val]
		}
	}
	return newStones
}

func splitDigitWithEvenNumDigitsWithCache(d, numDigits int, preCalcMap *map[int][2]int) (int, int) {
	var leftHalf int
	var rightHalf int
	if preCalcMap == nil {
		pcmObj := make(map[int][2]int)
		preCalcMap = &pcmObj
	}
	if _, ok := (*preCalcMap)[d]; ok {
		leftHalf, rightHalf = (*preCalcMap)[d][0], (*preCalcMap)[d][1]
		return leftHalf, rightHalf
	}
	leftHalf, rightHalf = splitDigitWithEvenNumDigits(d, numDigits)
	(*preCalcMap)[d] = [2]int{leftHalf, rightHalf}
	return leftHalf, rightHalf
}

func splitDigitWithEvenNumDigits(d, numDigits int) (leftHalf, rightHalf int) {
	divisor := 1
	for i := 0; i < numDigits/2; i++ {
		divisor *= 10
	}
	leftHalf = d / divisor
	rightHalf = d % divisor
	return leftHalf, rightHalf
}

func numberOfDigits(d int) int {
	if d == 0 {
		return 1
	}
	digitCount := 0
	for d > 0 {
		digitCount++
		d /= 10
	}
	return digitCount
}
