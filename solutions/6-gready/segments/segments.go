package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Segment struct {
	l int
	r int
}

func mergeSort(arr []Segment) []Segment {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	result := make([]Segment, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i].r <= right[j].r {
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

	segments := make([]Segment, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		parts := strings.Fields(scanner.Text())
		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])
		segments[i] = Segment{l: l, r: r}
	}

	segments = mergeSort(segments)

	count := 0
	lastEnd := 0

	for _, seg := range segments {
		if seg.l >= lastEnd {
			count++
			lastEnd = seg.r
		}
	}

	fmt.Println(count)
}
