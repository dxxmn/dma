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
	inDegree := make([]int, n+1)

	for i := 1; i <= n; i++ {
		scanner.Scan()
		arrStr := strings.Fields(scanner.Text())

		k, _ := strconv.Atoi(arrStr[0])
		graph[i] = make([]int, k)

		for j := 0; j < k; j++ {
			v, _ := strconv.Atoi(arrStr[j+1])
			graph[i][j] = v
			inDegree[v]++
		}
	}

	result := kanTopologicalSort(graph, inDegree, n)

	if len(result) == n {
		for i, v := range result {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
	} else {
		fmt.Println(-1)
	}
}

func kanTopologicalSort(graph [][]int, inDegree []int, n int) []int {
	queue := []int{}
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := []int{}

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		result = append(result, v)

		for _, neighbor := range graph[v] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(result) == n {
		return result
	}
	return []int{}
}
