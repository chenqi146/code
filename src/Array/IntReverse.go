package main

import "fmt"

func reverse(x int) int {
	if x == 0 {
		return 0
	}

	isAbs := true
	if x < 0 {
		isAbs = false
	}

	temp := 0
	for i := 1; x > 10; i++ {
		temp = temp + (i-1)*10 + (x % (i * 10))
		x = x / 10
	}

	if isAbs {
		return temp
	} else {
		return -temp
	}
}

func main() {
	fmt.Print(reverse(123))
}
