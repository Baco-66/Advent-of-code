package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Node represents a point in the grid for A*
type Node struct {
	Point  Point
	Type   rune
	Cost   int
	Dist   int
	Parent *Node
	Index  int // For heap.Interface
}

// PriorityQueue implements a min-heap for Nodes
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return (pq[i].Cost*100)+pq[i].Dist < (pq[j].Cost*100)+pq[j].Dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index, pq[j].Index = i, j
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*Node)
	node.Index = len(*pq)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	node.Index = -1 // For safety
	*pq = old[0 : n-1]
	return node
}

// Utility functions
func isValid(point Point, grid [][]rune) bool {
	return point.X >= 0 && point.X < len(grid) && point.Y >= 0 && point.Y < len(grid[0])
}

func distance(start, finish Point) int {
	return int(math.Sqrt(math.Pow(float64(finish.X-start.X), 2) + math.Pow(float64(finish.Y-start.Y), 2)))
}

// Function to find a key by value
func findKeyByValue(m map[rune][2]int, value [2]int) (rune, bool) {
	for key, val := range m {
		if val == value {
			return key, true
		}
	}
	return ' ', false
}

func weight(current Node, next, finish Point) *Node {
	remaining_distance := distance(next, finish)

	var directions = map[rune][2]int{
		'^': {-1, 0},
		'v': {1, 0},
		'<': {0, -1},
		'>': {0, 1},
	}

	direction := [2]int{current.Point.X - next.X, current.Point.Y - next.Y}

	newNode := &Node{
		Point: next,
		Cost:  current.Cost,
		Dist:  remaining_distance,
	}

	if directions[current.Type] == direction {
		newNode.Cost += 1
		newNode.Type = current.Type
	} else {
		newNode.Cost += 1001
		if dir, err := findKeyByValue(directions, direction); err {
			newNode.Type = dir
		} else {
			panic(fmt.Sprintf("Invalid direction: %v. Cannot determine type.", direction))
		}
	}

	return newNode
}

func reconstructPath(node *Node) []Point {
	path := []Point{}
	for current := node; current != nil; current = current.Parent {
		path = append([]Point{current.Point}, path...)
	}
	return path
}

func a_star(matrix [][]rune, fn func(Node, Point, Point) *Node) ([]Point, int, bool) {
	var start, end Point
	if x, y, err := findInMatrix(matrix, 'S'); err {
		start.X, start.Y = x, y
	} else {
		return nil, -1, false
	}
	if x, y, err := findInMatrix(matrix, 'E'); err {
		end.X, end.Y = x, y
	} else {
		return nil, -1, false
	}

	directions := []Point{
		{-1, 0}, // Up
		{1, 0},  // Down
		{0, -1}, // Left
		{0, 1},  // Right
	}

	openSet := make(PriorityQueue, 0)
	heap.Init(&openSet)

	startNode := &Node{
		Point: start,
		Type:  '>',
		Cost:  0,
		Dist:  distance(start, end),
	}

	heap.Push(&openSet, startNode)

	visited := make(map[Point]bool)

	for openSet.Len() > 0 {
		current := heap.Pop(&openSet).(*Node)

		if current.Point == end {
			return reconstructPath(current), current.Cost, true
		}

		visited[current.Point] = true

		for _, dir := range directions {
			neighbor := Point{current.Point.X + dir.X, current.Point.Y + dir.Y}

			if !isValid(neighbor, matrix) || matrix[neighbor.X][neighbor.Y] == '#' || visited[neighbor] {
				continue
			}
			newNode := fn(*current, neighbor, end)
			newNode.Parent = current
			heap.Push(&openSet, newNode)
		}
	}

	return nil, -1, false
}
