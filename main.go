package main

import (
	"fmt"
	"strconv"
)

var dp []string = make([]string, 200)

func main() {
	dp[0] = "1"
	dp[1] = "1"
	i := 1
	for i <= 100 {
		fmt.Println(i)
		fmt.Println("new", factorial(strconv.Itoa(i)))
		fmt.Println("------------------------------------")
		i += 1
	}
}

func factorial(n string) string {
	x, err := strconv.Atoi(n)
	if err != nil {
		return ""
	}
	if dp[x] != "" {
		return dp[x]
	}
	less := strconv.Itoa(x - 1)
	return Multiply(factorial(less), n)
}

func Add(n, o string) string {
	large := n
	small := o
	if len(large) < len(small) {
		small = n
		large = o
	}
	i := 1
	tmp := 0
	c := 0
	var out string
	var value = 0
	ln := len(large)
	sn := len(small)
	for i <= ln {
		tmp += int(large[ln-i] - '0')
		if i <= sn {
			tmp += int(small[sn-i] - '0')
		}
		tmp += c
		value = tmp % 10
		c = tmp / 10
		out = strconv.Itoa(value) + out
		tmp = 0
		i += 1
	}
	if c != 0 {
		out = strconv.Itoa(c) + out
	}
	return out
}

func multiplyOne(n string, ch byte) string {
	i := 1
	tmp := 0
	c := 0
	out := ""
	v := 0
	ln := len(n)
	for i <= ln {
		x := int(ch - '0')
		y := int(n[ln-i] - '0')
		tmp = x*y + c
		c = tmp / 10
		v = tmp % 10
		out = strconv.Itoa(v) + out
		i += 1
		tmp = 0
	}

	if c != 0 {
		out = strconv.Itoa(c) + out
	}

	return out
}

func Multiply(n, o string) string {
	small := n
	large := o
	if len(small) > len(large) {
		small = o
		large = n
	}
	out := ""
	i := 1
	sn := len(small)
	tmp := ""
	for i <= len(small) {
		char := small[sn-i]
		tmp = multiplyOne(large, char)
		tmp = AppendZeros(tmp, i-1)
		out = Add(out, tmp)
		i += 1
	}
	return out
}

func AppendZeros(input string, n int) string {
	i := 0
	for i < n {
		input += "0"
		i += 1
	}
	return input
}
