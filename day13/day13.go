package day13

import (
	"adventOfCode2024/utils"
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Part1() int {
	defer utils.TimeTrack(time.Now())
	games := getInputValues("day13/input.txt")
	totalCost := 0
	for _, game := range games {
		cheapestSoln, cheapestCost, err := checkWinPossible(game)
		fmt.Println(game, cheapestSoln, cheapestCost, err)
		if err != nil {
			continue
		}
		if cheapestSoln == [2]int{0, 0} {
			continue
		}
		totalCost += cheapestCost
	}
	return totalCost
}

func Part2() int {
	defer utils.TimeTrack(time.Now())
	return 0
}

type ButtonPress struct {
	xOffset, yOffset int
}
type Position struct {
	x, y int
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
			buttons = append(buttons, ButtonPress{x, y})
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
		prizeLocation := Position{x, y}
		games = append(games, GameInfo{buttons[0], buttons[1], prizeLocation})
	}
	return games
}

func checkWinPossible(game GameInfo) (cheapestSoln [2]int, cheapestCost int, err error) {
	a, b, c := game.aButton.xOffset, game.bButton.xOffset, game.prizeLocation.x
	positiveXSolns, err := getSolutions(a, b, c)
	if err != nil {
		return [2]int{0, 0}, 0, err
	}
	if len(positiveXSolns) == 0 {
		return [2]int{0, 0}, 0, errors.New("no positive x solutions found")
	}
	a, b, c = game.aButton.yOffset, game.bButton.yOffset, game.prizeLocation.y
	cheapestCost = math.MaxInt32
	for _, soln := range positiveXSolns {
		solnCost := 3*soln[0] + soln[1]
		if a*soln[0]+b*soln[1] == c && solnCost < cheapestCost {
			cheapestCost = solnCost
			cheapestSoln = soln
		}
	}
	return cheapestSoln, cheapestCost, nil
}

func getSolutions(a, b, c int) (solutions [][2]int, err error) {
	gcd, x, y, s, t := utils.ExtendedGcd(a, b)
	if c%gcd != 0 {
		return solutions, errors.New("gcd does not divide c")
	}
	x *= c / gcd
	y *= c / gcd
	absS := utils.IntAbs(s)
	absT := utils.IntAbs(t)
	upperBound := x / absS
	lowerBound := -y/absT + 1
	for k := lowerBound; k <= upperBound; k++ {
		newX := x - k*absS
		newY := y + k*absT
		if newX >= 0 && newY >= 0 {
			solutions = append(solutions, [2]int{newX, newY})
		}
	}
	return solutions, nil
}
