package day10

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
	trailheads := getInputValues("day10/input.txt")
	pathStack := list.New()
	trailCount := 0
	for _, th := range trailheads {
		seenEnds := make(map[[2]int]bool)
		pathStack.PushBack(th)
		for pathStack.Len() > 0 {
			top := pathStack.Back()
			pathStack.Remove(top)
			topTrailPos := top.Value.(*TrailPosition)
			if topTrailPos.value == 9 {
				seenEnds[topTrailPos.location] = true
			}
			for _, edge := range topTrailPos.edges {
				pathStack.PushBack(edge)
			}
		}
		trailCount += len(seenEnds)
	}
	return trailCount
}

func Part2() int {
	defer utils.TimeTrack(time.Now())
	trailheads := getInputValues("day10/input.txt")
	pathStack := list.New()
	trailCount := 0
	for _, th := range trailheads {
		pathStack.PushBack(th)
		for pathStack.Len() > 0 {
			top := pathStack.Back()
			pathStack.Remove(top)
			topTrailPos := top.Value.(*TrailPosition)
			if topTrailPos.value == 9 {
				trailCount++
			}
			for _, edge := range topTrailPos.edges {
				pathStack.PushBack(edge)
			}
		}
	}
	return trailCount
}

type TrailPosition struct {
	value    int
	location [2]int
	edges    []*TrailPosition
}

func getInputValues(filename string) (trailheads []*TrailPosition) {
	input := strings.Split(strings.Trim(utils.GetFileText(filename), " \n"), "\n")
	var tpMatrix [][]*TrailPosition
	for i, row := range input {
		tpMatrix = append(tpMatrix, []*TrailPosition{})
		rowI := strings.Split(strings.Trim(row, " "), "")
		for j, col := range rowI {
			if col == "." {
				tpMatrix[i] = append(tpMatrix[i], nil)
			} else {
				intCol, err := strconv.Atoi(col)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
				tpMatrix[i] = append(
					tpMatrix[i],
					&TrailPosition{intCol, [2]int{i, j}, []*TrailPosition{}},
				)
				if intCol == 0 {
					trailheads = append(trailheads, tpMatrix[i][j])
				}
			}
		}
	}
	for i, row := range tpMatrix {
		for j, col := range row {
			if col != nil {
				if i > 0 {
					if tpMatrix[i-1][j] != nil {
						if tpMatrix[i-1][j].value == col.value+1 {
							col.edges = append(col.edges, tpMatrix[i-1][j])
						}
					}
				}
				if j < len(row)-1 {
					if tpMatrix[i][j+1] != nil {
						if tpMatrix[i][j+1].value == col.value+1 {
							col.edges = append(col.edges, tpMatrix[i][j+1])
						}
					}
				}
				if i < len(tpMatrix)-1 {
					if tpMatrix[i+1][j] != nil {
						if tpMatrix[i+1][j].value == col.value+1 {
							col.edges = append(col.edges, tpMatrix[i+1][j])
						}
					}
				}
				if j > 0 {
					if tpMatrix[i][j-1] != nil {
						if tpMatrix[i][j-1].value == col.value+1 {
							col.edges = append(col.edges, tpMatrix[i][j-1])
						}
					}
				}
			}
		}
	}
	//for i := 0; i < len(tpMatrix); i++ {
	//	for j := 0; j < len(tpMatrix[i]); j++ {
	//		if tpMatrix[i][j] != nil {
	//			fmt.Print(tpMatrix[i][j].value, " ")
	//		} else {
	//			fmt.Print(".", " ")
	//		}
	//	}
	//	fmt.Println()
	//}
	//fmt.Println(trailheads)
	return trailheads
}
