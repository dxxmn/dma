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

	result := countComponents(graph, n)
	fmt.Println(result)
}

func countComponents(graph [][]int, n int) int {
	visited := make([]bool, n+1)
	count := 0

	for i := 1; i <= n; i++ {
		if !visited[i] {
			count++
			bfs(graph, i, visited)
		}
	}

	return count
}

func bfs(graph [][]int, start int, visited []bool) {
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		for _, neighbor := range graph[vertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}
