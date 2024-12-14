package day9

import (
	"adventOfCode2024/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Part1() int {
	defer utils.TimeTrack(time.Now())
	diskArray, emptyPositions, numCount := getInputValuesP1("day9/input.txt")
	sum := 0
	j := 0
	for i := len(diskArray) - 1; (i >= 0) && (j < len(emptyPositions)) && (i-numCount >= 0); i-- {
		for diskArray[i] == -1 {
			i--
		}
		diskArray[emptyPositions[j]] = diskArray[i]
		diskArray[i] = -1
		j++
	}
	for i, v := range diskArray {
		if v > 0 {
			sum += v * i
		}
	}
	return sum
}

func Part2() int {
	defer utils.TimeTrack(time.Now())
	diskArray, _, _ := getInputValuesP1("day9/input.txt")
	firstEmptyPos := 0
	for firstEmptyPos < len(diskArray) && diskArray[firstEmptyPos] != -1 {
		firstEmptyPos++
	}
	if firstEmptyPos == len(diskArray) {
		return calculateChecksum(diskArray)
	}
	revStart := len(diskArray) - 1
	for revStart >= 0 && diskArray[revStart] == -1 {
		revStart--
	}
	if revStart < 0 {
		return 0
	}

	maxId := diskArray[revStart]
	for maxId >= 0 && revStart >= 0 {
		for revStart >= 0 && diskArray[revStart] != maxId {
			revStart--
		}
		if firstEmptyPos > revStart {
			break
		}

		endBlockIdx := revStart
		startBlockIdx := revStart
		for startBlockIdx > 0 && diskArray[startBlockIdx-1] == maxId {
			startBlockIdx--
		}
		blockSize := endBlockIdx - startBlockIdx + 1

		startEmptyBlockIdx := firstEmptyPos
		endEmptyBlockIdx := startEmptyBlockIdx + blockSize - 1
		for endEmptyBlockIdx < startBlockIdx {
			usableBlockFound := true
			actualEndEmptyBlockIdx := startEmptyBlockIdx
			for ; actualEndEmptyBlockIdx <= endEmptyBlockIdx; actualEndEmptyBlockIdx++ {
				if diskArray[actualEndEmptyBlockIdx] != -1 {
					usableBlockFound = false
					break
				}
			}
			if usableBlockFound {
				for i := 0; i < blockSize; i++ {
					diskArray[startEmptyBlockIdx+i] = diskArray[startBlockIdx+i]
					diskArray[startBlockIdx+i] = -1
				}
				break
			} else {
				startEmptyBlockIdx = actualEndEmptyBlockIdx + 1
				for startEmptyBlockIdx+blockSize-1 < startBlockIdx && diskArray[startEmptyBlockIdx] != -1 {
					startEmptyBlockIdx++
				}
				endEmptyBlockIdx = startEmptyBlockIdx + blockSize - 1
			}
		}
		if startEmptyBlockIdx == firstEmptyPos {
			firstEmptyPos = startEmptyBlockIdx + blockSize
			for firstEmptyPos < startBlockIdx && diskArray[firstEmptyPos] != -1 {
				firstEmptyPos++
			}
		}
		maxId--
	}
	return calculateChecksum(diskArray)
}

func calculateChecksum(diskArray []int) (checksum int) {
	for i := 0; i < len(diskArray); i++ {
		if diskArray[i] > 0 {
			checksum += diskArray[i] * i
		}
	}
	return checksum
}

func getInputValuesP1(filename string) ([]int, []int, int) {
	input := strings.Split(strings.Trim(utils.GetFileText(filename), " \n"), "")
	var diskArray []int
	var emptyPositions []int
	id := 0
	for i, v := range input {
		intV, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		if i%2 == 0 {
			for j := 0; j < intV; j++ {
				diskArray = append(diskArray, id)
			}
			id++
		} else {
			for j := 0; j < intV; j++ {
				diskArray = append(diskArray, -1)
				emptyPositions = append(emptyPositions, len(diskArray)-1)
			}
		}
	}
	emptyPosLen := len(emptyPositions)
	daLen := len(diskArray)
	numCount := daLen - emptyPosLen
	var newEmptyPositions []int
	for _, emptyPos := range emptyPositions {
		if emptyPos <= numCount {
			newEmptyPositions = append(newEmptyPositions, emptyPos)
		} else {
			break
		}
	}

	for i := numCount; i < len(diskArray); i++ {
		newEmptyPositions = append(newEmptyPositions, i)
	}
	return diskArray, newEmptyPositions, numCount
}
