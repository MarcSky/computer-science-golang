package bfs

import (
	"container/list"
	"fmt"
)

func bfsSearchMaxValue(vertexCount int, values []int, m [][]int) int {
	type node struct {
		Value int
	}

	nodes := make(map[int]*node, vertexCount)

	for i := 0; i < vertexCount; i++ {
		nodes[i] = &node{Value: values[i]}
	}

	var watched = make([]bool, vertexCount)
	var maxValue = -1
	for i := 0; i < vertexCount; i++ {
		for j := 0; j < vertexCount; j++ {
			if m[i][j] == 1 && watched[j] == false {
				if nodes[j].Value > maxValue {
					maxValue = nodes[j].Value
				}
				watched[j] = true
			}
		}
	}

	return maxValue
}

func bfsSearchPath(vertexCount int, m map[int][]int, n1, n2 int) int {
	if n1 > vertexCount || n2 > vertexCount {
		return -1
	}

	var watched = make([]bool, vertexCount)
	var distances = make([]int, vertexCount)
	var parents = make([]int, vertexCount)
	var queue = list.New()

	watched[n1] = true
	queue.PushBack(n1)

	for queue.Len() > 0 {
		v := queue.Remove(queue.Front())
		for _, vertex := range m[v.(int)] {
			if watched[vertex] == false {
				watched[vertex] = true
				queue.PushBack(vertex)
				distances[vertex] = distances[v.(int)] + 1
				parents[vertex] = v.(int)
			}
		}
	}

	for _, parent := range parents {
		fmt.Print(parent, " ")
	}

	return distances[n2]
}
