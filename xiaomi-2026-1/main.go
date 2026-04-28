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
	firstLine := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(firstLine[0])
	l, _ := strconv.Atoi(firstLine[1])

	scanner.Scan()
	secondLine := strings.Split(scanner.Text(), " ")
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(secondLine[i])
	}

	currentSum := 0

	for i := 0; i < n; i++ {
		currentSum += a[i]
	}

	maxSum := currentSum

	for i := 0; i < l; i++ {
		currentSum = currentSum - a[i-1] + a[i]
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	fmt.Println(maxSum)
}
