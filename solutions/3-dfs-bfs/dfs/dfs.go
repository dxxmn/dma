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
	
	result := dfs(graph, 1, n)
	
	for i, v := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
}

func dfs(graph [][]int, start, n int) []int {
	visited := make([]bool, n+1)
	stack := []int{start}
	result := []int{}
	
	for len(stack) > 0 {
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		if !visited[vertex] {
			visited[vertex] = true
			result = append(result, vertex)
			
			for i := len(graph[vertex]) - 1; i >= 0; i-- {
				neighbor := graph[vertex][i]
				if !visited[neighbor] {
					stack = append(stack, neighbor)
				}
			}
		}
	}
	
	return result
}
