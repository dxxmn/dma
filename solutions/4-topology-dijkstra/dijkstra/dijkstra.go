package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	To     int
	Weight int
}

func dijkstra(graph map[int][]Edge, start int) (map[int]int, map[int]int) {
	dist := make(map[int]int)
	parent := make(map[int]int)
	unvisited := make(map[int]bool)

	for v := range graph {
		dist[v] = 1e9
		parent[v] = -1
		unvisited[v] = true
	}
	dist[start] = 0

	for len(unvisited) > 0 {
		v := -1
		minDist := int(1e9)

		for u := range unvisited {
			if dist[u] < minDist {
				minDist = dist[u]
				v = u
			}
		}

		if v == -1 {
			break
		}

		delete(unvisited, v)

		for _, edge := range graph[v] {
			u := edge.To
			if dist[u] > dist[v]+edge.Weight {
				dist[u] = dist[v] + edge.Weight
				parent[u] = v
			}
		}
	}

	return dist, parent
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(parts[0])
	a, _ := strconv.Atoi(parts[1])
	b, _ := strconv.Atoi(parts[2])

	graph := make(map[int][]Edge)
	for i := 1; i <= n; i++ {
		graph[i] = []Edge{}
	}

	for i := 1; i <= n; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		k, _ := strconv.Atoi(line[0])

		for j := 0; j < k; j++ {
			v, _ := strconv.Atoi(line[2*j+1])
			w, _ := strconv.Atoi(line[2*j+2])
			graph[i] = append(graph[i], Edge{To: v, Weight: w})
		}
	}

	distances, parents := dijkstra(graph, a)

	if distances[b] == 1e9 {
		fmt.Println(-1)
	} else {
		path := []int{}
		current := b
		for current != -1 {
			path = append([]int{current}, path...)
			current = parents[current]
		}

		fmt.Println(distances[b])
		for i, v := range path {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
	}
}
