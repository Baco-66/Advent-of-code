package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Node represents a point in the grid for A*
type Node_v2 struct {
	Point  Point
	Type   rune
	Cost   int
	Dist   int
	Parent []*Node_v2
	Index  int // For heap.Interface
}

// PriorityQueue implements a min-heap for Nodes
type PriorityQueue_v2 []*Node_v2

func (pq PriorityQueue_v2) Len() int { return len(pq) }

func (pq PriorityQueue_v2) Less(i, j int) bool {
	return (pq[i].Cost*100)+pq[i].Dist < (pq[j].Cost*100)+pq[j].Dist
}

func (pq PriorityQueue_v2) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index, pq[j].Index = i, j
}

func (pq *PriorityQueue_v2) Push(x interface{}) {
	node := x.(*Node_v2)
	node.Index = len(*pq)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue_v2) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	node.Index = -1 // For safety
	*pq = old[0 : n-1]
	return node
}

func reconstructAllPaths(node *Node_v2) []*Node_v2 {
	// Resulting path as a list of *Node_v2
	path := []*Node_v2{}

	// If the current node is not nil, process it
	if current := node; current != nil {
		// Append the current node to the path if it's unique
		path = appendUniqueNodes([]*Node_v2{current}, path)

		// Recursively reconstruct paths for each parent node
		for i := range current.Parent {
			path = appendUniqueNodes(path, reconstructAllPaths(current.Parent[i]))
		}
	}
	return path
}

// Helper function to append unique nodes to the list
func appendUniqueNodes(dest, src []*Node_v2) []*Node_v2 {
	existing := make(map[Point]bool)

	// Mark all nodes already in the destination as existing
	for _, node := range dest {
		existing[node.Point] = true
	}

	// Add nodes from the source if they are not already in the destination
	for _, node := range src {
		if !existing[node.Point] {
			dest = append(dest, node)
			existing[node.Point] = true
		}
	}
	return dest
}

func weight_controll_reverse(current Node_v2, next, finish Point) *Node_v2 {
	remaining_distance := distance(next, finish)

	var directions = map[rune][2]int{
		'^': {-1, 0},
		'v': {1, 0},
		'<': {0, -1},
		'>': {0, 1},
	}

	direction := [2]int{next.X - current.Point.X, next.Y - current.Point.Y}

	newNode := &Node_v2{
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
			var oposites = map[rune]rune{
				'^': 'v',
				'v': '^',
				'<': '>',
				'>': '<',
			}
			if oposites[current.Type] == newNode.Type {
				return nil
			}
		} else {
			panic(fmt.Sprintf("Invalid direction: %v. Cannot determine type.", direction))
		}
	}

	return newNode
}

func best_spots(matrix [][]rune, fn func(Node_v2, Point, Point) *Node_v2) ([]*Node_v2, bool) {
	var start, end Point
	if x, y, err := findInMatrix(matrix, 'S'); err {
		start.X, start.Y = x, y
	} else {
		return nil, false
	}
	if x, y, err := findInMatrix(matrix, 'E'); err {
		end.X, end.Y = x, y
	} else {
		return nil, false
	}

	directions := []Point{
		{-1, 0}, // Up
		{1, 0},  // Down
		{0, -1}, // Left
		{0, 1},  // Right
	}

	openSet := make(PriorityQueue_v2, 0)
	heap.Init(&openSet)

	costMap := make(map[Point]*Node_v2)

	startNode := &Node_v2{
		Point: start,
		Type:  '>',
		Cost:  0,
		Dist:  distance(start, end),
	}

	heap.Push(&openSet, startNode)

	costMap[start] = startNode

	goal := 0
	var best_spots []*Node_v2

	for openSet.Len() > 0 {
		current := heap.Pop(&openSet).(*Node_v2)

		if goal != 0 {
			if current.Cost <= goal && current.Point == end {
				best_spots = appendUniqueNodes(best_spots, reconstructAllPaths(current))
				continue
			}
		} else {
			if current.Point == end {
				goal = current.Cost
				best_spots = appendUniqueNodes(best_spots, reconstructAllPaths(current))
				continue
			}
		}

		for _, dir := range directions {
			neighbor := Point{current.Point.X + dir.X, current.Point.Y + dir.Y}

			if !isValid(neighbor, matrix) || matrix[neighbor.X][neighbor.Y] == '#' {
				continue
			}

			newNode := fn(*current, neighbor, end)

			if newNode == nil {
				continue
			}

			if old_node, exists := costMap[neighbor]; exists {
				if newNode.Cost == old_node.Cost {
					newNode.Parent = append(newNode.Parent, current)
					old_node.Parent = append(old_node.Parent, current)
				} else if math.Abs(float64(newNode.Cost-old_node.Cost)) <= +1000 {
					costMap[neighbor] = newNode
					newNode.Parent = append(newNode.Parent, current)
					heap.Push(&openSet, newNode)
				}

			} else {
				newNode.Parent = append(newNode.Parent, current)
				heap.Push(&openSet, newNode)
				costMap[neighbor] = newNode
			}

		}
	}

	if goal != 0 {
		return best_spots, true
	}

	return nil, false
}
