package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func quicksortWithPartition(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quicksortWithPartition(arr, low, pivotIndex-1)
		quicksortWithPartition(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	randomIndex := rand.Intn(high-low+1) + low
	arr[randomIndex], arr[high] = arr[high], arr[randomIndex]
	pivot := arr[high]

	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	const maxCapacity = 100 * 1024 * 1024 
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	
	scanner.Scan()
	arrStr := strings.Fields(scanner.Text())
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(arrStr[i])
	}

	quicksortWithPartition(arr, 0, len(arr)-1)
	for i:= 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}

}