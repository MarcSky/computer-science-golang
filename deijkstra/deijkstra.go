package deijkstra

const maxWeight = 100

func findDistance(vertexCount int, m map[int]map[int]int, from, to int) int {
	//watched := make([]bool, vertexCount)
	distances := make([]int, vertexCount)

	for i := 0; i < vertexCount; i++ {
		distances[i] = maxWeight
	}

	//parents := make([]int, vertexCount)

	return distances[to]
}
