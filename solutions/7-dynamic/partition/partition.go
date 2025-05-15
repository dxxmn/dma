package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}

	for i := 1; i <= n; i++ {
		f[i][i] = 1
	}

	for i := 1; i <= n; i++ {
		for k := 1; k < i; k++ {
			for j := 1; j <= k; j++ {
				f[i][k] += f[i-k][j]
			}
		}
	}

	sum := 0
	for k := 1; k <= n; k++ {
		sum += f[n][k]
	}

	fmt.Println(sum)
}
