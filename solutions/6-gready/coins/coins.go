package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	result := make([]int, 0, len(arr))
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
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

	const maxCapacity = 100 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(parts[0])
	k, _ := strconv.ParseInt(parts[1], 10, 64)

	scanner.Scan()
	pricesStr := strings.Fields(scanner.Text())
	prices := make([]int, n)
	for i := 0; i < n; i++ {
		prices[i], _ = strconv.Atoi(pricesStr[i])
	}

	prices = mergeSort(prices)

	count := 0
	var sum int64 = 0
	for i := 0; i < n; i++ {
		if sum+int64(prices[i]) <= k {
			sum += int64(prices[i])
			count++
		} else {
			break
		}
	}

	fmt.Println(count)
}
