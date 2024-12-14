package day6

import (
	"adventOfCode2024/utils"
	"strings"
	"time"
)

// TODO: See if there's a faster way to do Part2
// Roughly halved by not running the complete path for each cell on the original path

type GuardPos struct {
	row int
	col int
}

type GuardDir struct {
	rowOffset int
	colOffset int
}

type GuardPandD struct {
	pos GuardPos
	dir GuardDir
}

type GuardMap [][]rune
type GuardMapRow []rune

func (gd *GuardDir) turnRight90Degrees() GuardDir {
	return GuardDir{
		rowOffset: gd.colOffset,
		colOffset: -gd.rowOffset,
	}
}

func (gp *GuardPos) move(gd GuardDir) GuardPos {
	return GuardPos{
		row: gp.row + gd.rowOffset,
		col: gp.col + gd.colOffset,
	}
}

const obstacle = '#'

func (gp *GuardPos) validateInBounds(guardMap *GuardMap) bool {
	gm := *guardMap
	return gp.row >= 0 && gp.row < len(gm) && gp.col >= 0 && gp.col < len((gm)[0])
}

func (gp *GuardPos) validateNotObstacle(guardMap *GuardMap) bool {
	gm := *guardMap
	return gm[gp.row][gp.col] != obstacle
}

func Part1() int {
	defer utils.TimeTrack(time.Now())
	guardMap, guardPos, guardDir := getInputValues("day6/input.txt")
	visitedCells := make(map[GuardPos]bool)
	visitedCells[guardPos] = true

	newGuardPos := guardPos.move(guardDir)
	for newGuardPos.validateInBounds(&guardMap) {
		if newGuardPos.validateNotObstacle(&guardMap) {
			visitedCells[newGuardPos] = true
			guardPos = newGuardPos
		} else {
			guardDir = guardDir.turnRight90Degrees()
		}
		newGuardPos = guardPos.move(guardDir)
	}
	return len(visitedCells)
}

func Part2() int {
	defer utils.TimeTrack(time.Now())
	guardMap, guardPos, guardDir := getInputValues("day6/input.txt")
	isLoop, cellsOnOriginalPath := checkIfLoop(&guardMap, guardPos, guardDir)
	loopingLayoutCount := utils.BoolToInt(isLoop)
	prevGuardPos := guardPos
	prevGuardDir := guardDir
	tested := make(map[GuardPos]bool)
	for _, guardPosOnInitialPath := range cellsOnOriginalPath {
		if guardPosOnInitialPath.pos != guardPos {
			if !tested[guardPosOnInitialPath.pos] {
				tmp := guardMap[guardPosOnInitialPath.pos.row][guardPosOnInitialPath.pos.col]
				guardMap[guardPosOnInitialPath.pos.row][guardPosOnInitialPath.pos.col] = obstacle
				isLoop, _ = checkIfLoop(&guardMap, prevGuardPos, prevGuardDir)
				loopingLayoutCount += utils.BoolToInt(isLoop)
				guardMap[guardPosOnInitialPath.pos.row][guardPosOnInitialPath.pos.col] = tmp
				tested[guardPosOnInitialPath.pos] = true
				prevGuardDir = guardPosOnInitialPath.dir
				prevGuardPos = guardPosOnInitialPath.pos
			}
		}
	}
	return loopingLayoutCount
}

func checkIfLoop(guardMap *GuardMap, guardPos GuardPos, guardDir GuardDir) (bool, []GuardPandD) {
	visitedCellsWithDir := make(map[GuardPos]map[GuardDir]bool)
	visitedCellsWithDir[guardPos] = map[GuardDir]bool{guardDir: true}
	path := []GuardPandD{{guardPos, guardDir}}

	isLoop := false
	nextGuardPos := guardPos.move(guardDir)
	for nextGuardPos.validateInBounds(guardMap) {
		if nextGuardPos.validateNotObstacle(guardMap) {
			visitedCellDirMap, ok := visitedCellsWithDir[nextGuardPos]
			if !ok {
				visitedCellsWithDir[nextGuardPos] = map[GuardDir]bool{guardDir: true}
				path = append(path, GuardPandD{
					pos: nextGuardPos,
					dir: guardDir,
				})
				guardPos = nextGuardPos
			} else if !visitedCellDirMap[guardDir] {
				visitedCellDirMap[guardDir] = true
				path = append(path, GuardPandD{
					pos: nextGuardPos,
					dir: guardDir,
				})
				guardPos = nextGuardPos
			} else {
				isLoop = true
				break
			}
		} else {
			guardDir = guardDir.turnRight90Degrees()
		}
		nextGuardPos = guardPos.move(guardDir)
	}
	return isLoop, path
}

func getInputValues(filename string) (guardMap GuardMap, guardPos GuardPos, guardDir GuardDir) {
	rows := strings.Split(strings.Trim(utils.GetFileText(filename), " \n"), "\n")
	for i, row := range rows {
		var guardMapRow GuardMapRow
		for j, char := range row {
			guardMapRow = append(guardMapRow, char)
			if char == '^' {
				guardPos = GuardPos{i, j}
				guardDir = GuardDir{-1, 0}
			} else if char == 'v' {
				guardPos = GuardPos{i, j}
				guardDir = GuardDir{1, 0}
			} else if char == '<' {
				guardPos = GuardPos{i, j}
				guardDir = GuardDir{0, -1}
			} else if char == '>' {
				guardPos = GuardPos{i, j}
				guardDir = GuardDir{0, 1}
			}
		}
		guardMap = append(guardMap, guardMapRow)
	}
	return guardMap, guardPos, guardDir
}
