package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	firstLine := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(firstLine[0])
	a, _ := strconv.Atoi(firstLine[1])
	b, _ := strconv.Atoi(firstLine[2])

	graph := make([][]int, n+1)

	for i := 1; i <= n; i++ {
		scanner.Scan()
		arrStr := strings.Fields(scanner.Text())

		k, _ := strconv.Atoi(arrStr[0])
		graph[i] = make([]int, k)

		for j := 0; j < k; j++ {
			graph[i][j], _ = strconv.Atoi(arrStr[j+1])
		}
	}

	distance, path := shortestPath(graph, a, b, n)

	if distance == -1 {
		fmt.Println(-1)
	} else {
		fmt.Println(distance)
		for i, v := range path {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
	}
}

func shortestPath(graph [][]int, start, end, n int) (int, []int) {
	if start == end {
		return 0, []int{start}
	}

	visited := make([]bool, n+1)
	parent := make([]int, n+1)
	distance := make([]int, n+1)

	for i := 1; i <= n; i++ {
		parent[i] = -1
		distance[i] = -1
	}

	queue := []int{start}
	visited[start] = true
	distance[start] = 0

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		for _, neighbor := range graph[vertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
				parent[neighbor] = vertex
				distance[neighbor] = distance[vertex] + 1

				if neighbor == end {
					path := []int{}
					current := end
					for current != -1 {
						path = append([]int{current}, path...)
						current = parent[current]
					}
					return distance[end], path
				}
			}
		}
	}

	return -1, nil
}
