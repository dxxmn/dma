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
	nx := strings.Fields(scanner.Text())
	N, _ := strconv.Atoi(nx[0])

	scanner.Scan()
	arrStr := strings.Fields(scanner.Text())
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i], _ = strconv.Atoi(arrStr[i])
	}

	left, right := 0, N-1
	for left < right {
		mid := (left + right) / 2
		if (mid + 1) <= (N - 1) && arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	fmt.Println(left)
}