package route

import (
	"errors"
	"transit-server/models"
)

type matrixEntry struct {
	Actual    float64
	Heuristic float64
}

type pathEntry struct {
	Nodes []int
	Cost  float64
}

func FindPath(stops []models.Stop, connections []models.Connection, srcIndex, destIndex int) (GetRouteResult, error) {
	adjMatrix := make([][]matrixEntry, len(stops))
	for i := 0; i < len(adjMatrix); i++ {
		adjMatrix[i] = make([]matrixEntry, len(stops))
	}

	idIndexMap := make(map[uint]int)
	for i := 0; i < len(stops); i++ {
		idIndexMap[stops[i].StopId] = i
	}

	visitedArray := make([]bool, len(stops))
	for i := 0; i < len(stops); i++ {
		visitedArray[i] = false
	}

	for i := 0; i < len(connections); i++ {
		firstStopIdx := idIndexMap[connections[i].FirstStopId]
		secondStopIdx := idIndexMap[connections[i].SecondStopId]

		distance := HaversineDistance(Coordinate{stops[firstStopIdx].Lat, stops[firstStopIdx].Lng}, Coordinate{stops[secondStopIdx].Lat, stops[secondStopIdx].Lng})

		adjMatrix[firstStopIdx][secondStopIdx].Actual = distance
		adjMatrix[firstStopIdx][secondStopIdx].Heuristic = distance
	}

	for i := 0; i < len(adjMatrix); i++ {
		for j := 0; j < len(adjMatrix[i]); j++ {
			if i != j && adjMatrix[i][j].Actual == 0.0 {
				adjMatrix[i][j].Actual = 6371.0
				adjMatrix[i][j].Heuristic = HaversineDistance(Coordinate{stops[i].Lat, stops[i].Lng}, Coordinate{stops[j].Lat, stops[j].Lng})
			}
		}
	}

	expandQueue := make(PrioQueue, 0)
	expandQueue.Enqueue(pathEntry{[]int{srcIndex}, 0.0})

	currPath := pathEntry{[]int{srcIndex}, 0.0}
	currIndex := srcIndex

	for currIndex != destIndex {
		visitedArray[currIndex] = true
		for i := 0; i < len(adjMatrix[currIndex]); i++ {
			if adjMatrix[currIndex][i].Actual != 0 {
				newPath := pathEntry{}
				newPath.Nodes = make([]int, len(currPath.Nodes))
				copy(newPath.Nodes, currPath.Nodes)
				newPath.Nodes = append(newPath.Nodes, i)
				newPath.Cost = currPath.Cost + adjMatrix[currIndex][i].Actual + adjMatrix[currIndex][i].Heuristic

				expandQueue.Enqueue(newPath)
			}
		}

		currPath, _ = expandQueue.Dequeue()
		currIndex = currPath.Nodes[len(currPath.Nodes)-1]

		for visitedArray[currIndex] {
			currPath, _ = expandQueue.Dequeue()
			currIndex = currPath.Nodes[len(currPath.Nodes)-1]
		}
	}

	finalPath, ok := expandQueue.Dequeue()

	if !ok {
		return GetRouteResult{}, errors.New("failed to get path")
	}

	result := GetRouteResult{[]models.Stop{}, []Coordinate{}}

	for i := 0; i < len(finalPath.Nodes); i++ {
		currStop := stops[finalPath.Nodes[i]]
		result.Stops = append(result.Stops, currStop)
		result.Path = append(result.Path, Coordinate{currStop.Lat, currStop.Lng})
	}

	return result, nil
}
