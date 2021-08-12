package bfs

import "math/rand"

func randomMatrix(vertexCount int) [][]int {
	matrix := make([][]int, vertexCount)
	for i := 0; i < vertexCount; i++ {
		matrix[i] = make([]int, vertexCount)
	}

	for i := 0; i < vertexCount; i++ {
		for j := 0; j < vertexCount; j++ {
			if i == j {
				matrix[i][j] = 1
			} else {
				matrix[i][j] = rand.Intn(2)
			}
		}
	}

	return matrix
}
