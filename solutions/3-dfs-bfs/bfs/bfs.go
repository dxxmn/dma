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
	n, _ := strconv.Atoi(scanner.Text())

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

	result := bfs(graph, 1, n)

	for i, v := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
}

func bfs(graph [][]int, start, n int) []int {
	visited := make([]bool, n+1)
	queue := []int{start}
	visited[start] = true

	result := []int{start}

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		for _, neighbor := range graph[vertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
				result = append(result, neighbor)
			}
		}
	}

	return result
}
