package main

import (
	"fmt"
)

func isSubseq(s1, s2 string) bool {
	n, m := len(s1), len(s2)

	if (n-m)%2 != 0 || n < m {
		return false
	}
	k := (n - m) / 2

	for i := 0; i+k <= n; i++ {
		sub := s1[i : i+k]

		s1part := s1[:i] + s1[i+k:]

		s2part := s2 + sub

		if s1part == s2part {
			return true
		}

	}
	return false
}

func main() {
	var T int
	fmt.Scan(&T)

	for ; T > 0; T-- {
		var n, m int
		var s1, s2 string
		fmt.Scan(&n, &m)
		fmt.Scan(&s1, &s2)

		if isSubseq(s1, s2) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}
