package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func lowerBound(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := (left + right) / 2
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 1e6
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i], _ = strconv.Atoi(parts[i])
	}

	lis := make([]int, 0, n)
	for _, num := range nums {
		pos := lowerBound(lis, num)
		if pos == len(lis) {
			lis = append(lis, num)
		} else {
			lis[pos] = num
		}
	}

	fmt.Println(len(lis))
}
