package day5

import (
	"adventOfCode2024/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Part1() int {
	defer utils.TimeTrack(time.Now())
	rules, updates := getInputValues("day5/input.txt")
	middleUpdateOKSum := 0
	for _, updateArr := range updates {
		updateOk := true
		for i := 0; i < len(updateArr); i++ {
			rulesForUpdate, ok := rules[updateArr[i]]
			if !ok {
				continue
			}
			for j := i + 1; j < len(updateArr); j++ {
				if rulesForUpdate[updateArr[j]] {
					updateOk = false
					break
				}
			}
			if !updateOk {
				break
			}
		}
		if updateOk {
			middleUpdateOKSum += updateArr[(len(updateArr)-1)/2]
		}
	}
	return middleUpdateOKSum
}

func Part2() int {
	defer utils.TimeTrack(time.Now())
	rules, updates := getInputValues("day5/input.txt")
	sum := 0
	for _, updateArr := range updates {
		updateOK := true
		i := 0
		for i < len(updateArr) {
			rule, ok := rules[updateArr[i]]
			if ok {
				arrayResorted := false
				for j := i + 1; j < len(updateArr); j++ {
					if rule[updateArr[j]] {
						updateOK = false
						arrayResorted = true
						tmp := updateArr[j]
						updateArr[j] = updateArr[i]
						updateArr[i] = tmp
					}
				}
				if !arrayResorted {
					i++
				}
			} else {
				i++
			}
		}
		if !updateOK {
			sum += updateArr[(len(updateArr)-1)/2]
		}
	}
	return sum
}

func getInputValues(filename string) (map[int]map[int]bool, [][]int) {
	input := utils.GetFileText(filename)
	rulesUpdatesSplit := strings.Split(input, "\n\n")
	rulesSplit := strings.Split(rulesUpdatesSplit[0], "\n")
	updatesSplit := strings.Split(rulesUpdatesSplit[1], "\n")
	rules := make(map[int]map[int]bool)
	for _, rule := range rulesSplit {
		ruleSplit := strings.Split(strings.Trim(rule, " "), "|")
		pageNum, err := strconv.Atoi(ruleSplit[1])
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		precedingPageNum, err := strconv.Atoi(ruleSplit[0])
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		if _, ok := rules[pageNum]; !ok {
			rules[pageNum] = make(map[int]bool)
		}
		rules[pageNum][precedingPageNum] = true
	}
	var updates [][]int
	for _, update := range updatesSplit {
		trimUpdate := strings.Trim(update, " ")
		if trimUpdate == "" {
			continue
		}
		var updateArr []int
		updateSplit := strings.Split(trimUpdate, ",")
		for _, updateStr := range updateSplit {
			updateNum, err := strconv.Atoi(updateStr)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			updateArr = append(updateArr, updateNum)
		}
		updates = append(updates, updateArr)
	}
	return rules, updates
}
