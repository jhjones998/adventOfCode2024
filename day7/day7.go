package day7

import (
	"adventOfCode2024/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Operation struct {
	opTotal   int
	opMembers []int
}

var validMathOperations = []func(int, int) int{
	func(a, b int) int { return a + b },
	func(a, b int) int { return a * b },
	func(a, b int) int {
		stra := strconv.Itoa(a)
		strb := strconv.Itoa(b)
		strb += stra
		concat, err := strconv.Atoi(strb)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		return concat
	},
}

func Part1() int {
	defer utils.TimeTrack(time.Now())
	opSum := 0
	operations := getInputValues("day7/input.txt")
	for _, op := range operations {
		_, totals := testOp(op.opMembers, len(op.opMembers), nil, validMathOperations[:2])
		if totals[op.opTotal] {
			opSum += op.opTotal
		}
	}
	return opSum
}

func Part2() int {
	defer utils.TimeTrack(time.Now())
	opSum := 0
	operations := getInputValues("day7/input.txt")
	for _, op := range operations {
		_, totals := testOp(op.opMembers, len(op.opMembers), nil, validMathOperations)
		if totals[op.opTotal] {
			opSum += op.opTotal
		}
	}
	return opSum
}

func getInputValues(filename string) []Operation {
	input := utils.GetFileText(filename)
	rows := strings.Split(strings.Trim(input, " \n"), "\n")
	operations := make([]Operation, len(rows))
	for i, row := range rows {
		rowSplit := strings.Split(row, ":")
		opTotal, err := strconv.Atoi(strings.Trim(rowSplit[0], " "))
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		membersSplit := strings.Split(strings.Trim(rowSplit[1], " "), " ")
		opMembers := make([]int, len(membersSplit))
		for j, member := range membersSplit {
			memberInt, err := strconv.Atoi(member)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			opMembers[j] = memberInt
		}
		operations[i] = Operation{
			opTotal:   opTotal,
			opMembers: opMembers,
		}
	}
	return operations
}

func testOp(
	members []int,
	memberLen int,
	totals map[int]bool,
	validMathOperations []func(int, int) int,
) ([]int, map[int]bool) {
	if totals == nil {
		totals = make(map[int]bool)
	}
	if len(members) == 1 {
		results := []int{members[0]}
		return results, totals
	}
	var results []int
	nextRes, _ := testOp(members[:len(members)-1], memberLen, nil, validMathOperations)
	for _, op := range validMathOperations {
		for _, res := range nextRes {
			results = append(results, op(members[len(members)-1], res))
		}
		if memberLen == len(members) {
			for _, res := range results {
				totals[res] = true
			}
		}
	}
	return results, totals
}
