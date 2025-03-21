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
	X, _ := strconv.Atoi(nx[1])

	scanner.Scan()
	arrStr := strings.Fields(scanner.Text())
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i], _ = strconv.Atoi(arrStr[i])
	}

	left, right := 0, N-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] >= X {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if left == 0{
		fmt.Println(left)
	} else if left < N{
		if (arr[left] - X) >= (X - arr[left - 1]){
			fmt.Println(left-1)
		} else {
			fmt.Println(left)
		}
	} else {
		fmt.Println(-1)
	}
}