package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inf = int(1e9)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	firstLine := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(firstLine[0])
	m, _ := strconv.Atoi(firstLine[1])
	s, _ := strconv.Atoi(firstLine[2])
	t, _ := strconv.Atoi(firstLine[3])

	edges := make([]struct{ from, to, w int }, m)

	for i := 0; i < m; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		from, _ := strconv.Atoi(line[0])
		to, _ := strconv.Atoi(line[1])
		w, _ := strconv.Atoi(line[2])
		edges[i] = struct{ from, to, w int }{from, to, w}
	}

	dist, path := bellmanFord(n, edges, s, t)

	if dist == inf || dist == -1 {
		fmt.Println(-1)
	} else {
		fmt.Println(dist)
		for i, v := range path {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
	}
}

func bellmanFord(n int, edges []struct{ from, to, w int }, s, t int) (int, []int) {
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = inf
	}
	dist[s] = 0

	parent := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = -1
	}

	for i := 0; i < n-1; i++ {
		for _, e := range edges {
			if dist[e.from] != inf && dist[e.from]+e.w < dist[e.to] {
				dist[e.to] = dist[e.from] + e.w
				parent[e.to] = e.from
			}
		}
	}

	for _, e := range edges {
		if dist[e.from] != inf && dist[e.from]+e.w < dist[e.to] {
			return -1, nil
		}
	}

	if dist[t] == inf {
		return inf, nil
	}

	path := []int{}
	curr := t
	for curr != -1 {
		path = append([]int{curr}, path...)
		curr = parent[curr]
	}

	return dist[t], path
}
