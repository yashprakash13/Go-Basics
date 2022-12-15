package main

import "fmt"

type Graph struct {
	Vertices int
	Edge     map[int][]int
}

func (graph Graph) createGraph(numCourses int, prerequisites [][]int) Graph {
	graph.Vertices = numCourses
	graph.Edge = make(map[int][]int)
	for _, prereq := range prerequisites {
		graph.Edge[prereq[1]] = append(graph.Edge[prereq[1]], prereq[0])
	}
	fmt.Println(graph.Vertices)
	return graph
}

func (graph Graph) ifAllCoursesComplete() bool {
	inDegree := make([]int, graph.Vertices)
	for _, node := range graph.Edge {
		for _, prereq := range node {
			inDegree[prereq] += 1
		}
	}
	queue := make([]int, 0)
	for index, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, index)
		}
	}
	countVis := 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		countVis += 1

		for _, prereq := range graph.Edge[node] {
			inDegree[prereq] -= 1
			if inDegree[prereq] == 0 {
				queue = append(queue, prereq)
			}
		}
	}
	fmt.Println(graph.Vertices, countVis)
	if countVis == graph.Vertices {
		return true
	} else {
		return false
	}
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses == 1 {
		return true
	}
	if len(prerequisites) == 0 {
		return true
	}
	var graph Graph
	graph = graph.createGraph(numCourses, prerequisites)
	for _, i := range graph.Edge {
		fmt.Println(i)
	}
	return graph.ifAllCoursesComplete()
}

// to test on your machine, use this main function, otherwise copy all the above to leetcode
// func main() {
// 	prereq := [][]int{{1, 0}, {0, 1}}
// 	fmt.Println((canFinish(2, prereq)))

// 	prereq2 := [][]int{{1, 0}}
// 	fmt.Println((canFinish(2, prereq2)))
// }
