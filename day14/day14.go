package day14

import (
	"adventOfCode2024/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Part1() int {
	defer utils.TimeTrack(time.Now())
	// grid := getInputValues("day14/test.txt", 11, 7)
	grid := getInputValues("day14/input.txt", 101, 103)
	grid.moveRobots(100)
	quadrantMap := grid.getRobotsInQuadrants()
	safetyScore := 1
	for i := range quadrantMap {
		safetyScore *= len(quadrantMap[i])
	}

	return safetyScore
}

/* Low safety score works and feels better for Part2, but the condition for breaking out of the loop hurts my brain
   right now */

func Part2() int {
	defer utils.TimeTrack(time.Now())
	grid := getInputValues("day14/input.txt", 101, 103)
	robotsInARow := 0
	seconds := 0
	for {
		robotPositions := make(map[[2]int]bool)
		for _, r := range grid.robots {
			robotPositions[[2]int{r.xPos, r.yPos}] = true
		}
		for i := 0; i < grid.gridHeight; i++ {
			for j := 0; j < grid.gridWidth; j++ {
				if robotPositions[[2]int{j, i}] {
					robotsInARow++
				} else {
					robotsInARow = 0
				}
				if robotsInARow >= 10 {
					return seconds
				}
			}
		}
		seconds++
		grid.moveRobots(1)
	}
}

type Robot struct {
	xPos, yPos int
	xVel, yVel int
}

func (r *Robot) move(times, gridWidth, gridHeight int) {
	r.xPos = utils.PositiveMod(r.xPos+r.xVel*times, gridWidth)
	r.yPos = utils.PositiveMod(r.yPos+r.yVel*times, gridHeight)
}

type Grid struct {
	robots                                                       []*Robot
	gridWidth, gridHeight, gridWidthMidpoint, gridHeightMidpoint int
}

func (g *Grid) moveRobots(times int) {
	for _, r := range g.robots {
		r.move(times, g.gridWidth, g.gridHeight)
	}
}

func (g *Grid) getRobotsInQuadrants() (quadrantMap map[int][]*Robot) {
	quadrantMap = make(map[int][]*Robot)
	for i := 1; i <= 4; i++ {
		quadrantMap[i] = []*Robot{}
	}
	for _, r := range g.robots {
		xPos, yPos := r.xPos, r.yPos
		if xPos < g.gridWidthMidpoint && yPos < g.gridHeightMidpoint {
			quadrantMap[1] = append(quadrantMap[1], r)
		} else if xPos > g.gridWidthMidpoint && yPos < g.gridHeightMidpoint {
			quadrantMap[2] = append(quadrantMap[2], r)
		} else if xPos < g.gridWidthMidpoint && yPos > g.gridHeightMidpoint {
			quadrantMap[3] = append(quadrantMap[3], r)
		} else if xPos > g.gridWidthMidpoint && yPos > g.gridHeightMidpoint {
			quadrantMap[4] = append(quadrantMap[4], r)
		}
	}
	return quadrantMap
}

func (g *Grid) printGrid() {
	for y := 0; y < g.gridHeight; y++ {
		for x := 0; x < g.gridWidth; x++ {
			isRobot := false
			for _, r := range g.robots {
				if r.xPos == x && r.yPos == y {
					fmt.Print("R")
					isRobot = true
					break
				}
			}
			if !isRobot {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func getInputValues(filename string, gridX, gridY int) Grid {
	re := regexp.MustCompile(`p=(?P<posX>\d+),(?P<posY>\d+) v=(?P<velX>-?\d+),(?P<velY>-?\d+)`)
	input := strings.Trim(utils.GetFileText(filename), " \n")
	var robots []*Robot
	match := re.FindAllStringSubmatch(input, -1)
	for _, m := range match {
		posX, err := strconv.Atoi(m[1])
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		posY, err := strconv.Atoi(m[2])
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		velX, err := strconv.Atoi(m[3])
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		velY, err := strconv.Atoi(m[4])
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		robots = append(robots, &Robot{posX, posY, velX, velY})
	}
	return Grid{
		robots,
		gridX,
		gridY,
		(gridX - 1) / 2,
		(gridY - 1) / 2,
	}
}
