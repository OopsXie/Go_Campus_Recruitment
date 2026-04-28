package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func parsePart(part string, r rune) (float64, float64) {
	coeff := 0.0
	constVal := 0.0

	if part[0] != '+' && part[0] != '-' {
		part = "+" + part
	}

	i := 0
	n := len(part)

	for i < n {
		sign := 1.0
		if part[i] == '-' {
			sign = -1.0
		}
		i++

		numStr := ""
		for i < n && unicode.IsDigit(rune(part[i])) {
			numStr += string(part[i])
			i++
		}

		num := 1.0
		if numStr != "" {
			num, _ = strconv.ParseFloat(numStr, 64)
		}

		if i < n && rune(part[i]) == r {
			coeff += sign * num
			i++

		} else {
			constVal += sign * num
		}
	}
	return coeff, constVal
}

func main() {
	var equation string
	fmt.Scanln(&equation)

	parts := strings.Split(equation, "=")
	left, right := parts[0], parts[1]

	var r rune

	for _, c := range left {
		if unicode.IsLower(c) {
			r = c
			break
		}
	}
	if r == 0 {
		for _, c := range right {
			if unicode.IsLower(c) {
				r = c
				break
			}
		}
	}
	coeffLeft, constLeft := parsePart(left, r)
	coeffRight, constRight := parsePart(right, r)

	coeff := coeffLeft - coeffRight
	constVal := constRight - constLeft

	var result float64
	if coeff != 0 {
		result = constVal / coeff
	} else {
		result = 0
	}

	if result == 0 {
		fmt.Printf("%c=0.000\n", r)
	} else {
		fmt.Printf("%c=%.3f\n", r, result)
	}
}
