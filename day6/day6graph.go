package day6

import (
	"adventOfCode2024/utils"
	"strings"
	"time"
)

/* Around half the time again by using a graph structure instead of a matrix */

type node struct {
	pos        position
	isObstacle bool
	edges      map[direction]*node
}

type position [2]int

type direction int

type GuardPosAndDir struct {
	node *node
	dir  direction
}

const (
	up direction = iota
	right
	down
	left
)

var offMapNode = node{pos: position{-1, -1}}

func getInputValuesGraph(filename string) (guardPos *node, guardDir direction) {
	rows := strings.Split(strings.Trim(utils.GetFileText(filename), " \n"), "\n")
	createdNodes := map[position]*node{}
	charMap := map[rune]direction{
		'^': up,
		'>': right,
		'v': down,
		'<': left,
	}
	for i, row := range rows {
		for j, char := range row {
			pos := position{i, j}
			thisNode, ok := createdNodes[pos]
			if !ok {
				thisNode = &node{pos: pos, edges: map[direction]*node{}}
				createdNodes[pos] = thisNode
			}
			if dir, ok := charMap[char]; ok {
				guardDir = dir
				guardPos = createdNodes[pos]
			}
			if char != obstacle {
				if i > 0 {
					aboveNode, ok := createdNodes[position{i - 1, j}]
					if !ok {
						aboveNode = &node{pos: position{i - 1, j}, edges: map[direction]*node{}}
						createdNodes[position{i - 1, j}] = aboveNode
					}
					thisNode.edges[up] = aboveNode
					if rows[i-1][j] == obstacle {
						aboveNode.isObstacle = true
					} else {
						aboveNode.edges[down] = thisNode
					}
				} else {
					thisNode.edges[up] = &offMapNode
				}
				if j > 0 {
					leftNode, ok := createdNodes[position{i, j - 1}]
					if !ok {
						leftNode = &node{pos: position{i, j - 1}, edges: map[direction]*node{}}
						createdNodes[position{i, j - 1}] = leftNode
					}
					thisNode.edges[left] = leftNode
					if row[j-1] == obstacle {
						leftNode.isObstacle = true
					} else {
						leftNode.edges[right] = thisNode
					}
				} else {
					thisNode.edges[left] = &offMapNode
				}
				if i < len(rows)-1 {
					belowNode, ok := createdNodes[position{i + 1, j}]
					if !ok {
						belowNode = &node{pos: position{i + 1, j}, edges: map[direction]*node{}}
						createdNodes[position{i + 1, j}] = belowNode
					}
					thisNode.edges[down] = belowNode
					if rows[i+1][j] == obstacle {
						belowNode.isObstacle = true
					} else {
						belowNode.edges[up] = thisNode
					}
				} else {
					thisNode.edges[down] = &offMapNode
				}
				if j < len(row)-1 {
					rightNode, ok := createdNodes[position{i, j + 1}]
					if !ok {
						rightNode = &node{pos: position{i, j + 1}, edges: map[direction]*node{}}
						createdNodes[position{i, j + 1}] = rightNode
					}
					thisNode.edges[right] = rightNode
					if row[j+1] == obstacle {
						rightNode.isObstacle = true
					} else {
						rightNode.edges[left] = thisNode
					}
				} else {
					thisNode.edges[right] = &offMapNode
				}
			} else {
				thisNode.isObstacle = true
			}
		}
	}
	return guardPos, guardDir
}

func Part2Graph() int {
	defer utils.TimeTrack(time.Now())
	guardPos, guardDir := getInputValuesGraph("day6/input.txt")
	gpd := GuardPosAndDir{guardPos, guardDir}
	isLoop, path := checkIsLoop(gpd)
	loopConfigCount := utils.BoolToInt(isLoop)
	prevGpd := gpd
	testedNodes := map[*node]bool{guardPos: true}
	for _, gpdOnInitialPath := range path {
		if !testedNodes[gpdOnInitialPath.node] {
			gpdOnInitialPath.node.isObstacle = true
			isLoop, _ = checkIsLoop(prevGpd)
			loopConfigCount += utils.BoolToInt(isLoop)
			testedNodes[gpdOnInitialPath.node] = true
			gpdOnInitialPath.node.isObstacle = false
			prevGpd = gpdOnInitialPath
		}
	}
	return loopConfigCount
}

func checkIsLoop(gpd GuardPosAndDir) (bool, []GuardPosAndDir) {
	visitedNodes := map[GuardPosAndDir]bool{gpd: true}
	path := []GuardPosAndDir{gpd}
	isLoop := false
	for {
		nextNode, _ := gpd.node.edges[gpd.dir]
		if nextNode == &offMapNode {
			break
		}
		if nextNode.isObstacle {
			gpd.dir = (gpd.dir + 1) % 4
			continue
		}
		nextGpd := GuardPosAndDir{nextNode, gpd.dir}
		if visitedNodes[nextGpd] {
			isLoop = true
			break
		}
		visitedNodes[nextGpd] = true
		path = append(path, nextGpd)
		gpd = nextGpd
	}
	return isLoop, path
}
