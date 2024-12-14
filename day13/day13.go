package day13

import (
	"adventOfCode2024/utils"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const adjustment = 10000000000000

func Part1() (totalCost int64) {
	defer utils.TimeTrack(time.Now())
	games := getInputValues("day13/input.txt")
	for _, game := range games {
		ax, ay := game.aButton.xOffset, game.aButton.yOffset
		bx, by := game.bButton.xOffset, game.bButton.yOffset
		xGoal, yGoal := game.prizeLocation.x, game.prizeLocation.y
		for i := int64(0); i < int64(100); i++ {
			for j := int64(0); j < int64(100); j++ {
				if ax*i+bx*j == xGoal && ay*i+by*j == yGoal {
					totalCost += 3*i + j
				}
			}
		}
	}
	return totalCost
}

func Part2() (totalCost int64) {
	defer utils.TimeTrack(time.Now())
	games := getInputValues("day13/input.txt")
	for _, game := range games {
		ax, ay := game.aButton.xOffset, game.aButton.yOffset
		bx, by := game.bButton.xOffset, game.bButton.yOffset
		xGoal, yGoal := game.prizeLocation.x+adjustment, game.prizeLocation.y+adjustment
		divisor := ay*bx - by*ax
		if divisor != 0 {
			b := (xGoal*ay - yGoal*ax) / divisor
			a := (xGoal*by - yGoal*bx) / -divisor
			if ax*a+bx*b == xGoal && ay*a+by*b == yGoal {
				totalCost += 3*a + b
			}
		} else {
			if ay == by && ax == bx {
				if xGoal%bx == 0 {
					totalCost += xGoal / bx
				}
			}
		}

	}
	return totalCost
}

type ButtonPress struct {
	xOffset, yOffset int64
}
type Position struct {
	x, y int64
}
type GameInfo struct {
	aButton, bButton ButtonPress
	prizeLocation    Position
}

func getInputValues(filename string) (games []GameInfo) {
	input := strings.Split(strings.Trim(utils.GetFileText(filename), " \n"), "\n\n")
	buttonsRe := regexp.MustCompile(`X\+(?P<x>\d+), Y\+(?P<y>\d+)`)
	prizeRe := regexp.MustCompile(`X=(?P<x>\d+), Y=(?P<y>\d+)`)
	for _, game := range input {
		buttonsRes := buttonsRe.FindAllStringSubmatch(game, -1)
		var buttons []ButtonPress
		for _, match := range buttonsRes {
			x, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(match[2])
			if err != nil {
				panic(err)
			}
			buttons = append(buttons, ButtonPress{int64(x), int64(y)})
		}
		prizeRes := prizeRe.FindStringSubmatch(game)
		x, err := strconv.Atoi(prizeRes[1])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(prizeRes[2])
		if err != nil {
			panic(err)
		}
		prizeLocation := Position{int64(x), int64(y)}
		games = append(games, GameInfo{buttons[0], buttons[1], prizeLocation})
	}
	return games
}
