package day12

import (
	"adventOfCode2024/utils"
	"strings"
	"time"
)

func Part1() int {
	defer utils.TimeTrack(time.Now())
	cost := 0
	farmMap, rowCount, colCount := getInputValues("day12/input.txt")
	seenPositions := make(map[Position]bool)
	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			currPos := Position{i, j}
			if !seenPositions[currPos] {
				perimeter, area := getPlotValuesP1(&farmMap, currPos, rowCount, colCount, &seenPositions)
				cost += perimeter * area
			}
		}
	}
	return cost
}

func Part2() int {
	defer utils.TimeTrack(time.Now())
	cost := 0
	seenPositions := make(map[Position]bool)
	farmMap, rowCount, colCount := getInputValues("day12/input.txt")
	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			currPos := Position{i, j}
			if !seenPositions[currPos] {
				sides, area := getPlotValuesP2(&farmMap, currPos, rowCount, colCount, &seenPositions)
				cost += sides * area
			}
		}
	}
	return cost
}

type Position [2]int

func (p Position) add(other Position, rowCount, colCount int) Position {
	newPos := Position{p[0] + other[0], p[1] + other[1]}
	if newPos[0] < 0 || newPos[1] < 0 || newPos[0] >= rowCount || newPos[1] >= colCount {
		return nilPosition
	}
	return newPos
}
func (p Position) lt(other Position) bool {
	return p[0] < other[0] && p[1] < other[1]
}
func (p Position) eq(other Position) bool {
	return p[0] == other[0] && p[1] == other[1]
}
func (p Position) gt(other Position) bool {
	return !p.eq(other) && !p.lt(other)
}
func (p Position) gte(other Position) bool {
	return p.gt(other) && p.eq(other)
}
func (p Position) lte(other Position) bool {
	return p.eq(other) && p.lt(other)
}
func (p Position) neq(other Position) bool {
	return !p.eq(other)
}

var nilPosition = Position{-1, -1}

func getPlotValuesP2(farmMapPtr *map[Position]rune, startPos Position, rowCount, colCount int, seenPositionsPtr *map[Position]bool) (corners, area int) {
	farmMap := *farmMapPtr
	cropType := farmMap[startPos]
	seenPositions := *seenPositionsPtr
	seenPositions[startPos] = true
	area++
	nextPositions := getNextPositions(cropType, farmMapPtr, startPos, rowCount, colCount)
	corners = getCornerCount(cropType, farmMapPtr, startPos, rowCount, colCount)
	if len(nextPositions) == 0 {
		return corners, area
	}
	for x := 0; x < len(nextPositions); x++ {
		np := nextPositions[x]
		if !seenPositions[np] {
			seenPositions[np] = true
			npNextPositions := getNextPositions(cropType, farmMapPtr, np, rowCount, colCount)
			nextPositions = append(nextPositions, npNextPositions...)
			corners += getCornerCount(cropType, farmMapPtr, np, rowCount, colCount)
			area++
		}
	}
	return corners, area
}

func getCornerCount(cropType rune, farmMapPtr *map[Position]rune, currPos Position, rowCount, colCount int) (cornerCount int) {
	farmMap := *farmMapPtr
	// NW, N, NE, W, E, SW, S, SE
	cardinalPositions := getCardinalPositions(currPos, rowCount, colCount)

	// check nw corner
	nw := cardinalPositions[0]
	n := cardinalPositions[1]
	w := cardinalPositions[3]
	if n.neq(nilPosition) && farmMap[n] == cropType &&
		w.neq(nilPosition) && farmMap[w] == cropType &&
		(nw.eq(nilPosition) || farmMap[nw] != cropType) {
		cornerCount++
	}
	if (n.eq(nilPosition) || farmMap[n] != cropType) &&
		(w.eq(nilPosition) || farmMap[w] != cropType) {
		cornerCount++
	}

	// check ne corner
	ne := cardinalPositions[2]
	e := cardinalPositions[4]
	if n.neq(nilPosition) && farmMap[n] == cropType &&
		e.neq(nilPosition) && farmMap[e] == cropType &&
		(ne.eq(nilPosition) || farmMap[ne] != cropType) {
		cornerCount++
	}
	if (n.eq(nilPosition) || farmMap[n] != cropType) &&
		(e.eq(nilPosition) || farmMap[e] != cropType) {
		cornerCount++
	}

	// check sw corner
	sw := cardinalPositions[5]
	s := cardinalPositions[6]
	if s.neq(nilPosition) && farmMap[s] == cropType &&
		w.neq(nilPosition) && farmMap[w] == cropType &&
		(sw.eq(nilPosition) || farmMap[sw] != cropType) {
		cornerCount++
	}
	if (s.eq(nilPosition) || farmMap[s] != cropType) &&
		(w.eq(nilPosition) || farmMap[w] != cropType) {
		cornerCount++
	}

	// check se corner
	se := cardinalPositions[7]
	if s.neq(nilPosition) && farmMap[s] == cropType &&
		e.neq(nilPosition) && farmMap[e] == cropType &&
		(se.eq(nilPosition) || farmMap[se] != cropType) {
		cornerCount++
	}
	if (s.eq(nilPosition) || farmMap[s] != cropType) &&
		(e.eq(nilPosition) || farmMap[e] != cropType) {
		cornerCount++
	}

	return cornerCount
}

func getNextPositions(cropType rune, farmMapPtr *map[Position]rune, pos Position, rowCount, colCount int) (nextPositions [][2]int) {
	farmMap := *farmMapPtr
	nextPos := pos.add(Position{-1, 0}, rowCount, colCount)
	if nextPos.neq(nilPosition) && farmMap[nextPos] == cropType {
		nextPositions = append(nextPositions, nextPos)
	}
	nextPos = pos.add(Position{0, 1}, rowCount, colCount)
	if nextPos.neq(nilPosition) && farmMap[nextPos] == cropType {
		nextPositions = append(nextPositions, nextPos)
	}
	nextPos = pos.add(Position{1, 0}, rowCount, colCount)
	if nextPos.neq(nilPosition) && farmMap[nextPos] == cropType {
		nextPositions = append(nextPositions, nextPos)
	}
	nextPos = pos.add(Position{0, -1}, rowCount, colCount)
	if nextPos.neq(nilPosition) && farmMap[nextPos] == cropType {
		nextPositions = append(nextPositions, nextPos)
	}
	return nextPositions
}

func getPlotValuesP1(farmMapPtr *map[Position]rune, startPos Position, rowCount, colCount int, seenPositionsPtr *map[Position]bool) (perimeter, area int) {
	farmMap := *farmMapPtr
	cropType := farmMap[startPos]
	seenPositions := *seenPositionsPtr
	seenPositions[startPos] = true
	area++
	nextPositions := getNextPositions(cropType, farmMapPtr, startPos, rowCount, colCount)
	if len(nextPositions) == 0 {
		perimeter = 4
		return perimeter, area
	}
	perimeter += 4 - len(nextPositions)
	for x := 0; x < len(nextPositions); x++ {
		np := nextPositions[x]
		if !seenPositions[np] {
			seenPositions[np] = true
			npNextPositions := getNextPositions(cropType, farmMapPtr, np, rowCount, colCount)
			nextPositions = append(nextPositions, npNextPositions...)
			area++
			perimeter += 4 - len(npNextPositions)
		}
	}
	return perimeter, area
}

func getInputValues(filename string) (farmMap map[Position]rune, rowCount, colCount int) {
	farmMap = make(map[Position]rune)
	input := strings.Split(strings.Trim(utils.GetFileText(filename), " \n"), "\n")
	rowCount = len(input)
	colCount = len(input[0])
	for i, row := range input {
		for j, col := range []rune(row) {
			farmMap[Position{i, j}] = col
		}
	}
	return farmMap, rowCount, colCount
}

func getCardinalPositions(pos Position, rowCount, colCount int) (cardinalPositions []Position) {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			offset := Position{x, y}
			if offset.neq(Position{0, 0}) {
				newPos := pos.add(offset, rowCount, colCount)
				cardinalPositions = append(cardinalPositions, newPos)
			}
		}
	}
	return cardinalPositions
}
