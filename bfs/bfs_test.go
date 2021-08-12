package bfs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBFSSearchMaxValue(t *testing.T) {
	t.Parallel()

	type bfsSearchMaxValueParam struct {
		matrix      [][]int
		vertexCount int
		values      []int
		expected    int
	}

	graph := bfsSearchMaxValueParam{
		matrix: [][]int{
			{1, 1, 1, 1, 1},
			{0, 1, 0, 1, 0},
			{1, 0, 1, 1, 0},
			{0, 0, 1, 1, 1},
			{1, 1, 1, 0, 1},
		},
		vertexCount: 5,
		values:      []int{50, 42, 14, 95, 23},
		expected:    95,
	}

	value := bfsSearchMaxValue(graph.vertexCount, graph.values, graph.matrix)
	assert.Equal(t, graph.expected, value)
}

func TestBFSSearchPath(t *testing.T) {
	t.Parallel()

	type bfsSearchPathParam struct {
		graph       map[int][]int
		vertexCount int
		fromVertex  int
		toVertex    int
		expected    int
	}

	param := bfsSearchPathParam{
		graph: map[int][]int{
			0: {1},
			1: {2},
			2: {3},
			3: {4},
			4: {},
		},
		vertexCount: 5,
		fromVertex:  0,
		toVertex:    4,
		expected:    4,
	}

	value := bfsSearchPath(param.vertexCount, param.graph, param.fromVertex, param.toVertex)
	assert.Equal(t, param.expected, value)
}
