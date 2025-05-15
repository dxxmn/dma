package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	d int
	w int
}

func mergeSort(arr []Task) []Task {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	result := make([]Task, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i].w >= right[j].w {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 1e6
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	tasks := make([]Task, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		parts := strings.Fields(scanner.Text())
		d, _ := strconv.Atoi(parts[0])
		w, _ := strconv.Atoi(parts[1])
		if d > n {
			d = n
		}
		tasks[i] = Task{d: d, w: w}
	}

	tasks = mergeSort(tasks)

	parent := make([]int, n+2)
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	sum := 0
	for _, task := range tasks {
		available := find(task.d)
		if available > 0 {
			sum += task.w
			parent[available] = available - 1
		}
	}

	fmt.Println(sum)
}
