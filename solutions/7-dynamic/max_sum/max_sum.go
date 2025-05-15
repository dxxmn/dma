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
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i], _ = strconv.Atoi(parts[i])
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	if n > 1 {
		dp[1] = nums[0] + nums[1]
	}
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]) + nums[i]
	}
	fmt.Println(dp[n-1])
}
