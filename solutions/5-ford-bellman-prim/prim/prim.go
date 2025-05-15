package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Edge struct {
	u, v, weight int
}

type Item struct {
	vertex int
	weight int
	parent int
}

// c офф доки
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	firstLine := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(firstLine[0])
	m, _ := strconv.Atoi(firstLine[1])

	graph := make([][]struct{ to, weight int }, n+1)
	for i := 1; i <= n; i++ {
		graph[i] = []struct{ to, weight int }{}
	}

	for i := 0; i < m; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		u, _ := strconv.Atoi(line[0])
		v, _ := strconv.Atoi(line[1])
		w, _ := strconv.Atoi(line[2])

		graph[u] = append(graph[u], struct{ to, weight int }{v, w})
		graph[v] = append(graph[v], struct{ to, weight int }{u, w})
	}

	mstEdges, totalWeight := prim(graph, n)

	if len(mstEdges) != n-1 {
		fmt.Println(-1)
		return
	}

	fmt.Println(totalWeight)

	sort.Slice(mstEdges, func(i, j int) bool {
		if mstEdges[i].u != mstEdges[j].u {
			return mstEdges[i].u < mstEdges[j].u
		}
		return mstEdges[i].v < mstEdges[j].v
	})

	for _, edge := range mstEdges {
		fmt.Print(edge.u)
		fmt.Print(" ")
		fmt.Print(edge.v)
		fmt.Print(" ")
	}
}

func prim(graph [][]struct{ to, weight int }, n int) ([]Edge, int) {
	visited := make([]bool, n+1)
	mstEdges := []Edge{}
	totalWeight := 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{vertex: 1, weight: 0, parent: -1})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		u := item.vertex

		if visited[u] {
			continue
		}

		visited[u] = true

		if item.parent != -1 {
			u1, v1 := item.parent, u
			if u1 > v1 {
				u1, v1 = v1, u1
			}
			mstEdges = append(mstEdges, Edge{u1, v1, item.weight})
			totalWeight += item.weight
		}

		for _, neighbor := range graph[u] {
			v := neighbor.to
			weight := neighbor.weight
			if !visited[v] {
				heap.Push(&pq, &Item{vertex: v, weight: weight, parent: u})
			}
		}
	}

	for i := 1; i <= n; i++ {
		if !visited[i] {
			return mstEdges, totalWeight
		}
	}

	return mstEdges, totalWeight
}
